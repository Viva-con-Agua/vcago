package vcago

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//SessionAuth go to next if the request has a session else return 401.
func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["valid"] != nil {
			return next(c)
		}
		log.Print(sess)
		return echo.NewHTTPError(http.StatusUnauthorized, "")
	}
}

// CORSConfig for api services. Can be configured via .env.
var CORSConfig = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins:     strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	AllowCredentials: true,
})
