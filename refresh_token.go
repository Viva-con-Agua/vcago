package vcago

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RefreshToken struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func NewRefreshToken(userID string) *RefreshToken {
	return &RefreshToken{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

}

func (i *RefreshToken) SignedString(secret string) (string, error) {
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, i)
	return temp.SignedString([]byte(secret))
}

//RefreshCookieConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
//The token is reachable with c.Get("token")
func RefreshCookieConfig() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      &RefreshToken{},
			ContextKey:  "token",
			TokenLookup: "cookie:refresh_token",
			SigningKey:  []byte(jwtSecret),
		})
}

func RefreshCookieUserID(c echo.Context) (string, error) {
	token := c.Get("token").(*jwt.Token)
	if token == nil {
		return "", errors.New("No user in Conext")
	}
	return token.Claims.(*RefreshToken).UserID, nil
}
