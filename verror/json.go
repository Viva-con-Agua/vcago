package verror

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	//JSONValidator represents the json validator struct.
	JSONValidator struct {
		Validator *validator.Validate
	}
	//JSONError represents a json validation error. Used om return of JSONErrorResponse.
	JSONError struct {
		Errors []string `json:"errors"`
	}
)

type ErrJSONResponse struct {
	JSONError JSONError
}

func (e *ErrJSONResponse) Error() string {
	return fmt.Sprintf("%v", e.JSONError)
}

//Validate extend JSONValidator with Validate function.
func (cv *JSONValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

//JSONErrorResponse creates a response for json validation error.
func NewJSONError(err error) (jsonErr *JSONError) {
	jsonErr = new(JSONError)
	errorList := strings.Split(err.Error(), "\n")
	for _, val := range errorList {
		eList := strings.Split(val, "Key: ")
		eList = strings.Split(eList[1], " Error:")
		snake := matchFirstCap.ReplaceAllString(eList[1], "${1}${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
		jsonErr.Errors = append(jsonErr.Errors, strings.ToLower(snake))
	}
	return
}
func (i *JSONError) Error() string {
	return fmt.Sprintf("%v", i.Errors)
}

//JSONValidate validates a json bind in echo.Context.
//The interface i is used for validation.
//If the c.Bind(i) or the validation returns errors the function return an APIError.
func JSONValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return BadRequest("json bind error", err)
	}
	// validate body
	if err := c.Validate(i); err != nil {
		return BadRequest("json validation error", NewJSONError(err))
	}
	return nil
}
