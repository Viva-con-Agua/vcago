package vcago

import (
	"net/http"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/labstack/echo/v4"
)

func JsonErrorHandling(c echo.Context, i interface{}) (r_err *verr.ResponseError) {
	if err := c.Bind(i); err != nil {
		return &verr.ResponseError{Code: http.StatusBadRequest, Response: err}
	}
	// validate body
	if err := c.Validate(i); err != nil {
		return &verr.ResponseError{Code: http.StatusBadRequest, Response: verr.JsonErrorResponse(err)}
	}
	return nil

}

func ResponseErrorHandling(c echo.Context, api_err *verr.ApiError, i interface{}) (r_err *verr.ResponseError) {
	if api_err.Error != nil {
		if api_err.ResponseError.Code == http.StatusInternalServerError {
			LogApiError(api_err, c, i)
			return api_err.ResponseError
		}
		return api_err.ResponseError
	}
	return nil
}
