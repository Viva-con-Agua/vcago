package vcago

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = Settings.String("JWT_SECRET", "w", "secret")

var authCookie = NewCookieConfig()

//AuthToken represents the authentication tokens for handling access via jwt.
type AuthToken struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
}

//NewAuthToken creates an new access and refresh token for the given user.
func NewAuthToken(accessToken jwt.Claims, refreshToken jwt.Claims) (r *AuthToken, err error) {
	r = new(AuthToken)
	if r.AccessToken, err = SignedString(accessToken); err != nil {
		return
	}
	r.RefreshToken, err = SignedString(refreshToken)
	return
}

func SignedString(token jwt.Claims) (string, error) {
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return temp.SignedString([]byte(jwtSecret))
}

//AccessCookie return an cookie conains the access_token.
func (i *AuthToken) AccessCookie() (r *http.Cookie) {
	return authCookie.Cookie("access_token", i.AccessToken)
}

//ResetAccessCookie returns an cookie for reset the access_token.
func ResetAccessCookie() *http.Cookie {
	return authCookie.Cookie("access_token", "")
}

//RefreshCookie returns an cookie conains the refresh_token.
func (i *AuthToken) RefreshCookie() *http.Cookie {
	return authCookie.Cookie("refresh_token", i.RefreshToken)
}

//ResetRefreshCookie returns an cookie for reset the refresh_token.
func ResetRefreshCookie() *http.Cookie {
	return authCookie.Cookie("refresh_token", "")
}
