package vcago

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//APIKey represents the api Bearer token
var APIKey = Settings.String("API_KEY", "w", "secret")

//KeyAuthMiddleware middleware function for handling authentication via key.
func KeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == APIKey, nil
	})
}
