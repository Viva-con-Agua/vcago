package vcago

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AccessToken struct {
	ID            string         `json:"id,omitempty" bson:"_id"`
	Email         string         `json:"email" bson:"email" validate:"required,email"`
	FirstName     string         `json:"first_name" validate:"required"`
	LastName      string         `json:"last_name" validate:"required"`
	FullName      string         `json:"full_name"`
	DisplayName   string         `json:"display_name"`
	Roles         RoleListCookie `json:"system_roles"`
	Country       string         `json:"country"`
	PrivacyPolicy bool           `json:"privacy_policy"`
	Confirmed     bool           `json:"confirmed"`
	LastUpdate    string         `json:"last_update"`
	jwt.StandardClaims
}

func NewAccessToken(user *User) *AccessToken {
	return &AccessToken{
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.FullName,
		user.DisplayName,
		*user.Roles.Cookie(),
		user.Country,
		user.PrivacyPolicy,
		user.Confirmd,
		user.LastUpdate,
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
	//Deprecated: AccessCookieConfig not longer supported, use AccessCookieMiddleware
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      &AccessToken{},
			ContextKey:  "token",
			TokenLookup: "cookie:access_token",
			SigningKey:  []byte(jwtSecret),
		})
}

func AccessCookieUser(c echo.Context) (r *AccessToken, err error) {
	token := c.Get("token").(*jwt.Token)
	if token == nil {
		return nil, errors.New("No user in Conext")
	}
	r = token.Claims.(*AccessToken)
	return
}
