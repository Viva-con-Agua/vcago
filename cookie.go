package vcago

import "net/http"

//CookieConfig represents the cookie parameters
type CookieConfig struct {
	SameSite http.SameSite
	Secure   bool
	HttpOnly bool
}

//NewCookieConfig loads the cookie parameters from the .env file and return a new CookieConfig.
func NewCookieConfig() *CookieConfig {
	cookieSameSite := Config.GetEnvString("COOKIE_SAME_SITE", "w", "strict")
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
		Secure:   Config.GetEnvBool("COOKIE_SECURE", "w", true),
		HttpOnly: Config.GetEnvBool("COOKIE_HTTP_ONLY", "w", true),
	}
}

//Cookie returns an http.Cookie using the CookieConfig parameters.
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
