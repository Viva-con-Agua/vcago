package vcago

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var jwtSecret = Config.GetEnvString("JWT_SECRET", "w", "secret")

//AuthCookie represents the default authentication cookie
var AuthCookie = AuthCookieDefault()

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
	return &token.Claims.(*AccessToken).User, nil
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

//AuthCookieDefault returns an http.Cookie with the default set parameters.
func AuthCookieDefault() (r *http.Cookie) {
	var sameSite string
	sameSite = Config.GetEnvString("COOKIE_SAME_SITE", "w", "strict")
	if sameSite == "lax" {
		HTTPBaseCookie.SameSite = http.SameSiteLaxMode
	}
	if sameSite == "strict" {
		HTTPBaseCookie.SameSite = http.SameSiteStrictMode
	}
	if sameSite == "none" {
		HTTPBaseCookie.SameSite = http.SameSiteNoneMode
	}
	r.Secure = Config.GetEnvBool("COOKIE_SECURE", "w", true)
	r.HttpOnly = Config.GetEnvBool("COOKIE_HTTP_ONLY", "w", true)
	r.Path = "/"
	r.Domain = Config.GetEnvString("COOKIE_DOMAIN", "w", "localhost")
	return
}

type AuthToken struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
}

func NewAuthToken(user *User) (r *AuthToken, err error) {
	r = new(AuthToken)
	if r.AccessToken, err = NewAccessToken(user).SignedString(jwtSecret); err != nil {
		return
	}
	r.RefreshToken, err = NewRefreshToken(user.ID).SignedString(jwtSecret)
	return
}

func (i *AuthToken) AccessCookie() (r *http.Cookie) {
	r = AuthCookie
	r.Name = "access_token"
	r.Value = i.AccessToken
	return
}

func (i *AuthToken) RefreshCookie() (r *http.Cookie) {
	r = AuthCookie
	r.Name = "refresh_token"
	r.Value = i.RefreshToken
	return
}
