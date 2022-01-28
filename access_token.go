package vcago

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthToken struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
}

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
	temp := jwt.NewWithClaims(jwt.SigningMethodES256, i)
	return temp.SignedString([]byte(secret))
}

func NewAuthToken(u *User) (r *AuthToken, err error) {
	access := jwt.NewWithClaims(jwt.SigningMethodES256, NewAccessToken(u))
	refresh := jwt.NewWithClaims(jwt.SigningMethodES256)
	return
}
