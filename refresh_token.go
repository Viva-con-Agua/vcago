package vcago

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
func RefreshCookieUserID(c echo.Context) (string, error) {
	token := c.Get("token").(*jwt.Token)
	if token == nil {
		return "", errors.New("no user in context")
	}
	return token.Claims.(*RefreshToken).UserID, nil
}
