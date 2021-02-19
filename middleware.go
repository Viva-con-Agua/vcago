package vcago

import (
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//NewCORSConfig create a echo middleware for cors handling.
func NewCORSConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins:     strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowCredentials: true,
	}
}

//SessionOptions used for user session
var SessionOptions = new(sessions.Options)
//SessionLoadEnv loads session
func SessionLoadEnv() {
	var l LoadEnv
sameSite, l := l.GetEnvString("COOKIE_SAME_SITE", "w", "strict")
	CookieSecure, l := l.GetEnvBool("COOKIE_SECURE", "w", true)
	HTTPOnly, l := l.GetEnvBool("COOKIE_HTTP_ONLY", "w", true)
	//MaxAge, l := l.GetEnvInt("COOKIE_MAX_AGE", "w", 86400*7)
	var SameSite http.SameSite
	if sameSite == "lax" {
		SameSite = http.SameSiteLaxMode
	}
	if sameSite == "strict" {
		SameSite = http.SameSiteStrictMode
	}
	if sameSite == "none" {
		SameSite = http.SameSiteNoneMode
	}
	SessionOptions = &sessions.Options{
		Path:     "/",
	//	MaxAge:   MaxAge,
		HttpOnly: HTTPOnly,
		SameSite: SameSite,
		Secure:   CookieSecure,
	}
}
//SessionWithID return session wir id
func SessionWithID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if val, ok := sess.Values["sessionID"]; ok {
			c.Set("session_id", val)
			return next(c)
		}
		sess.Options = SessionOptions
		sessionID := uuid.New().String()
		sess.Values["sessionID"] = sessionID
		sess.Save(c.Request(), c.Response())
		c.Set("session_id", sessionID)
		return next(c)
}
}
