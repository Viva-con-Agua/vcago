package vcago

import (
	"net/http"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

//Validator represents a Json validator
type (
	Validator struct {
		Validator *validator.Validate
	}
)

//Validate interface i
func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

//JSONErrorHandler formats JsonError to ResponseError
func JSONErrorHandler(c echo.Context, i interface{}) (rErr *verr.ResponseError) {
	if err := c.Bind(i); err != nil {
		return &verr.ResponseError{Code: http.StatusBadRequest, Response: err}
	}
	// validate body
	if err := c.Validate(i); err != nil {
		return &verr.ResponseError{Code: http.StatusBadRequest, Response: verr.JsonErrorResponse(err)}
	}
	return nil

}

//ResponseErrorHandler handles ApiError
func ResponseErrorHandler(c echo.Context, apiErr *verr.ApiError, i interface{}) (rErr *verr.ResponseError) {
	if apiErr.Error != nil {
		if apiErr.ResponseError.Code == http.StatusInternalServerError {
			LogAPIError(apiErr, c, i)
			return apiErr.ResponseError
		}
		return apiErr.ResponseError
	}
	return nil
}
