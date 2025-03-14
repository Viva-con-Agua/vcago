package vcago

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CookieConfig represents the cookie parameters
type CookieConfig struct {
	SameSite http.SameSite
	Secure   bool
	HttpOnly bool
}

// NewCookieConfig loads the cookie parameters from the .env file and return a new CookieConfig.
func NewCookieConfig() *CookieConfig {
	cookieSameSite := Settings.String("COOKIE_SAME_SITE", "w", "strict")
	sameSite := http.SameSiteStrictMode
	if cookieSameSite == "lax" {
		sameSite = http.SameSiteLaxMode
	}
	if cookieSameSite == "strict" {
		sameSite = http.SameSiteStrictMode
	}
	if cookieSameSite == "none" {
		sameSite = http.SameSiteNoneMode
	}
	return &CookieConfig{
		SameSite: sameSite,
		Secure:   Settings.Bool("COOKIE_SECURE", "w", true),
		HttpOnly: Settings.Bool("COOKIE_HTTP_ONLY", "w", true),
	}
}

// Cookie returns an http.Cookie using the CookieConfig parameters.
// @param name  cookie name
// @param value cookie value
func (i *CookieConfig) Cookie(name string, value string) *http.Cookie {
	return &http.Cookie{
		SameSite: i.SameSite,
		Secure:   i.Secure,
		HttpOnly: i.HttpOnly,
		Path:     "/",
		Name:     name,
		Value:    value,
	}
}

// AccessCookieConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
// The token is reachable with c.Get("token")
func AccessCookieMiddleware(i jwt.Claims) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      i,
			ContextKey:  "token",
			TokenLookup: "cookie:access_token",
			SigningKey:  []byte(jwtSecret),
		})
}

// RefreshCookieConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
// The token is reachable with c.Get("token")
func RefreshCookieMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      &RefreshToken{},
			ContextKey:  "token",
			TokenLookup: "cookie:refresh_token",
			SigningKey:  []byte(jwtSecret),
		})
}
