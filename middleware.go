package vcago

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSConfig for api services. Can be configured via .env.
var CORSConfig = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins:     strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	AllowCredentials: true,
})
