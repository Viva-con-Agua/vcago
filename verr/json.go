package verr

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	//JSONValidator validation struct using for JSON
	JSONValidator struct {
		Validator *validator.Validate
	}

	//JSONError represents a json validation error. Used om return of JSONErrorResponse:w
	JSONError struct {
		Error []string `json:"error"`
	}
)

var ErrJSONBind = errors.New("json can't bind to the interface")
var ErrJSONValidate = errors.New("json is not valid")

//Validate extend JSONValidator with Validate function.
func (cv *JSONValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

//JSONErrorResponse creates a response for json validation error.
func JSONErrorResponse(err error) (jList []JSONError) {
	if err != ErrJSONBind {
		jsonErr := new(JSONError)
		errorList := strings.Split(err.Error(), "\n")
		for _, val := range errorList {
			eList := strings.Split(val, "Key: ")
			eList = strings.Split(eList[1], " Error:")
			jsonErr.Error = append(jsonErr.Error, ToSnakeCase(eList[1]))
		}
	}
	return jList
}

//JSONValidate validates a json bind in echo.Context.
//The interface i is used for validation.
//If the c.Bind(i) or the validation returns errors the function return an APIError.
func JSONValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		LogError(c.Request().Context(), err, "debug")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// validate body
	if err := c.Validate(i); err != nil {
		LogError(c.Request().Context(), err, "debug")
		return echo.NewHTTPError(http.StatusBadRequest, JSONErrorResponse(err))
	}
	return nil
}
