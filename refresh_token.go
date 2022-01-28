package vcago

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	temp := jwt.NewWithClaims(jwt.SigningMethodES256, i)
	return temp.SignedString([]byte(secret))
}
