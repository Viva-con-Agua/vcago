package verror

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	JSONValidator struct {
		Validator *validator.Validate
	}
	//JSONError represents a json validation error. Used om return of JSONErrorResponse:w
	JSONError struct {
		Key   string
		Error string
	}
)

//Validate extend JSONValidator with Validate function.
func (cv *JSONValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

//JSONErrorResponse creates a response for json validation error.
func JSONErrorResponse(e error) (jList []JSONError) {
	jsonErr := new(JSONError)
	errorList := strings.Split(e.Error(), "\n")
	for _, val := range errorList {
		eList := strings.Split(val, "Key: ")
		eList = strings.Split(eList[1], " Error:")
		jsonErr.Key = eList[0]
		jsonErr.Error = eList[1]
		jList = append(jList, *jsonErr)
	}
	return jList
}

//JSONValidate validates a json bind in echo.Context.
//The interface i is used for validation.
//If the c.Bind(i) or the validation returns errors the function return an APIError.
func JSONValidate(c echo.Context, i interface{}) *Error {
	if err := c.Bind(i); err != nil {
		return New(err).Status(http.StatusBadRequest)
		//return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// validate body
	if err := c.Validate(i); err != nil {
		return New(err).Body(JSONErrorResponse(err)).Status(http.StatusBadRequest)
	}
	return nil
}
