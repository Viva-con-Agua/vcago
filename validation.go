package vcago

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func JsonValid(c echo.Context, i interface{}) (err error) {
	if err = c.Bind(i); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(i); err != nil {
		return c.JSON(http.StatusBadRequest, JsonErrorResponse(err))
	}
	return nil

}
