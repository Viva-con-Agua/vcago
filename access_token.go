package vcago

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AccessToken struct {
	User User `json:"user"`
	jwt.StandardClaims
}

func NewAccessToken(user *User) *AccessToken {
	return &AccessToken{
		*user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
}

func (i *AccessToken) SignedString(secret string) (string, error) {
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, i)
	return temp.SignedString([]byte(secret))
}

//AccessCookieConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
//The token is reachable with c.Get("token")
func AccessCookieConfig() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      &AccessToken{},
			ContextKey:  "token",
			TokenLookup: "cookie:access_token",
			SigningKey:  []byte(jwtSecret),
		})
}

func AccessCookieUser(c echo.Context) (*User, error) {
	token := c.Get("token").(*jwt.Token)
	if token == nil {
		return nil, errors.New("No user in Conext")
	}
	user := &token.Claims.(*AccessToken).User
	return user, nil
}
