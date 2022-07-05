package vcago

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEchoServer(service string) (r *echo.Echo) {
	Settings.Bool("DEBUG", "w", true)
	r = echo.New()
	r.Debug = true
	r.Use(CORS.Init())
	r.HTTPErrorHandler = HTTPErrorHandler
	r.Use(Logger.Init(service))
	r.Validator = JSONValidator
	r.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewString()
		},
		TargetHeader: echo.HeaderXRequestID,
	}))
	return
}
