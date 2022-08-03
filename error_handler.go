package vcago

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func HTTPErrorHandler(err error, c echo.Context) {
	//set default error code
	code := http.StatusInternalServerError
	//check if the err is an normal response
	if resp, ok := err.(*Response); ok {
		c.JSON(resp.Response())
		//check if error is normal error
	} else if resp, ok := err.(*Error); ok {
		id := c.Request().Header.Get(echo.HeaderXRequestID)
		if id == "" {
			id = c.Response().Header().Get(echo.HeaderXRequestID)
		}
		resp.Print(id)
		c.JSON(resp.Response())
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Code == http.StatusInternalServerError {
			c.JSON(NewInternalServerError("-").Response())
		} else {
			c.JSON(code, err)
		}
	} else if resp, ok := err.(*ValidationError); ok {
		c.JSON(NewBadRequest("-", "validation error", resp).Response())
	} else if err == mongo.ErrNoDocuments {
		c.JSON(NewResp(http.StatusNotFound, "error", "document not found", "", nil).Response())
	} else {
		c.JSON(NewInternalServerError("-").Response())
	}
}
