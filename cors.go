package vcago

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSConfig struct {
	allowOrigins []string
}

var CORS = new(CORSConfig)

//NewCORSConfig create a echo middleware for cors handling.
func (i *CORSConfig) Init() echo.MiddlewareFunc {
	allowOrigins := Settings.StringList("ALLOW_ORIGINS", "w", []string{"localhost:8080"})
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowCredentials: true,
	})
}
