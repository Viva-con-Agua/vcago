package vcago

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//HTTPErrorHandler handles echo.HTTPError and return the correct response.
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if resp, ok := err.(*Response); ok {
		c.JSON(resp.Response())
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Code == http.StatusInternalServerError {
			c.JSON(NewInternalServerError("-").Response())
		} else {
			c.JSON(code, err)
		}
	} else if resp, ok := err.(*MongoError); ok {
		c.JSON(resp.Response())
	} else if resp, ok := err.(*ValidationError); ok {
		c.JSON(resp.Response())
	} else if resp, ok := err.(*Status); ok {
		c.JSON(resp.Response())
	} else {
		c.JSON(NewInternalServerError("-").Response())
	}
}
