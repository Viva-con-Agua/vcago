package vauth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

type (
	JWToken struct {
		AccessToken  string `json:"access_token" bson:"access_token"`
		RefreshToken string `json:"refresh_token" bson:"refresh_token"`
		Type         string `json:"type" bson:"type"`
		ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
		Scope        string `json:"scope" bson:"scope"`
	}
	AccessToken struct {
		User User `json:"user"`
		jwt.StandardClaims
	}
	RefreshToken struct {
		UserID string `json:"user_id"`
		jwt.StandardClaims
	}
)

//AccessConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
//The token is reachable with c.Get("token")
var AccessConfig = middleware.JWTConfig{
	Claims:     &AccessToken{},
	ContextKey: "token",
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
}

//RefreshConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
//The token is reachable with c.Get("token")
var RefreshConfig = middleware.JWTConfig{
	Claims:     &RefreshToken{},
	ContextKey: "token",
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
}

//NewJWTToken returns a new JWTToken model contains an access and an refresh token
func NewJWTToken(u *User, scope string) (*JWToken, error) {
	var exAT = time.Now().Add(time.Minute * 15).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &AccessToken{
		*u,
		jwt.StandardClaims{
			ExpiresAt: exAT,
		},
	})
	at, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &RefreshToken{
		u.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})
	rf, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &JWToken{AccessToken: at, RefreshToken: rf, TokenType: "Bearer", ExpiresAt: exAT, Scope: scope}, nil
}
