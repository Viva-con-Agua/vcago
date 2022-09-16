package vcago

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server represents an echo.Echo interface.
type Server struct {
	echo.Echo
}

// corsMiddleware is used for load the allow origins and return an CORSMiddleware.
func corsMiddleware() echo.MiddlewareFunc {
	allowOrigins := Settings.StringList("ALLOW_ORIGINS", "w", []string{"localhost:8080"})
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowCredentials: true,
	})
}

// Run starts the echo.Echo server on default port 1323. Use the APP_PORT param for set an custom port.
func (i *Server) Run() {
	port := Settings.String("APP_PORT", "n", "1323")
	i.Logger.Fatal(i.Start(":" + port))
}

// NewServer
func NewServer() *Server {
	Settings.Bool("DEBUG", "w", true)
	Settings.Load()
	r := echo.New()
	r.Debug = true
	r.Use(corsMiddleware())
	r.HTTPErrorHandler = HTTPErrorHandler
	r.Use(Logger.Init())
	r.Validator = JSONValidator
	r.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewString()
		},
		TargetHeader: echo.HeaderXRequestID,
	}))
	r.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthz")
	})
	return &Server{*r}
}
