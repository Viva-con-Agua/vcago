package vcago

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

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

func GetAccessToken(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

var CORSConfig = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins:     strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	AllowCredentials: true,
})

/*
func CheckPermission(permission *PermissionList) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			val := sess.Values["user"]
			var user = &User{}
			user, _ = val.(*User)
			for _, v := range user.Access {
				if permission.Contains(v) {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusUnauthorized, resp.Unauthorized())
		}
	}
}*/
