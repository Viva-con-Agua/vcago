package vcago

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//HTTPErrorHandler handles echo.HTTPError and return the correct response.
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Code == http.StatusInternalServerError {
			c.JSON(code, ErrorResponse{Message: "internal_server_error"})
		} else {
			c.JSON(code, err)
		}
	} else if resp, ok := err.(*MongoError); ok {
		c.JSON(resp.Response())
	} else if resp, ok := err.(*ValidationError); ok {
		c.JSON(resp.Response())
	} else if resp, ok := err.(*Status); ok {
		c.JSON(resp.Response())
	} else if resp, ok := err.(*ErrorResponse); ok {
		c.JSON(resp.Response())
	} else {
		c.JSON(code, ErrorResponse{Message: "internal_server_error"})
	}
}
