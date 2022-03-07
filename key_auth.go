package vcago

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ApiKey = Config.GetEnvString("API_KEY", "w", "secret")

func KeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == ApiKey, nil
	})
}
