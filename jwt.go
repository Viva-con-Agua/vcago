package vcago

import (
	"net/http"
	"os"

	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/labstack/echo/v4/middleware"
)

//JWTMiddleware handles authentication by jwt
type JWTMiddleware struct {
    RequestCount uint64 `json:"request_count"`
    Scope string `json:"scope"`
}

//AccessCookieConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
//The token is reachable with c.Get("token") 
var AccessCookieConfig = middleware.JWTConfig{
		Claims:     &vmod.AccessToken{},
		ContextKey: "token",
        TokenLookup: "cookie:access_token",
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
//RefreshCookieConfig can with echo for middleware.JWTWithConfig(vmod.AccessConfig) to handling access controll
//The token is reachable with c.Get("token") 
var RefreshCookieConfig = middleware.JWTConfig{
		Claims:     &vmod.RefreshToken{},
		ContextKey: "token",
        TokenLookup: "cookie:refresh_token",
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}


//JWTNewAccessCookie create a new http.Cookie contains the access_token.
func JWTNewAccessCookie(token *vmod.JWTToken) (*http.Cookie){
    var cookie = HTTPBaseCookie
    cookie.Name = "access_token"
    cookie.Value = token.AccessToken
    return &cookie
}

//JWTNewRefreshCookie create a new http.Cookie contains the refresh_token.
func JWTNewRefreshCookie(token *vmod.JWTToken) (*http.Cookie){
    var cookie = HTTPBaseCookie
    cookie.Name = "refresh_token"
    cookie.Value = token.RefreshToken
    return &cookie
}
