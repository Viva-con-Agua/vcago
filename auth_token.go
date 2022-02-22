package vcago

import (
	"log"
	"net/http"
)

var jwtSecret = Config.GetEnvString("JWT_SECRET", "w", "secret")

//AuthCookieDefault returns an http.Cookie with the default set parameters.
func AuthCookieDefault() (r *http.Cookie) {
	r = new(http.Cookie)
	var sameSite string
	sameSite = Config.GetEnvString("COOKIE_SAME_SITE", "w", "strict")
	if sameSite == "lax" {
		r.SameSite = http.SameSiteLaxMode
	}
	if sameSite == "strict" {
		r.SameSite = http.SameSiteStrictMode
	}
	if sameSite == "none" {
		r.SameSite = http.SameSiteNoneMode
	}
	r.Secure = Config.GetEnvBool("COOKIE_SECURE", "w", true)
	r.HttpOnly = Config.GetEnvBool("COOKIE_HTTP_ONLY", "w", true)
	r.Path = "/"
	//r.Domain = vcago.Config.GetEnvString("COOKIE_DOMAIN", "w", "localhost")
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
		log.Print(err)
		return
	}
	r.RefreshToken, err = NewRefreshToken(user.ID).SignedString(jwtSecret)
	return
}

func (i *AuthToken) AccessCookie() (r *http.Cookie) {
	r = AuthCookieDefault()
	r.Name = "access_token"
	r.Value = i.AccessToken
	return
}
func ResetAccessCookie() (r *http.Cookie) {
	r = AuthCookieDefault()
	r.Name = "access_token"
	r.Value = ""
	return
}
func (i *AuthToken) RefreshCookie() (r *http.Cookie) {
	r = AuthCookieDefault()
	r.Name = "refresh_token"
	r.Value = i.RefreshToken
	return
}

func ResetRefreshCookie() (r *http.Cookie) {
	r = AuthCookieDefault()
	r.Name = "refresh_token"
	r.Value = ""
	return
}
