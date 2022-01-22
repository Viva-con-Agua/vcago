package vauth

import (
	"github.com/Viva-con-Agua/vcago/vutils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSConfig struct {
	allowOrigins []string
}

var CORS = new(CORSConfig)

func (i *CORSConfig) LoadEnv() *CORSConfig {
	i.allowOrigins = vutils.Config.GetEnvStringList("ALLOW_ORIGINS", "w", []string{"localhost:8080"})
	return i
}

//NewCORSConfig create a echo middleware for cors handling.
func (i *CORSConfig) New() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins:     i.allowOrigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowCredentials: true,
	}
}
