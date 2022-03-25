package vcago

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

var JSONValidator = NewValidator()

func NewValidator() (r *Validator) {
	r = new(Validator)
	v := validator.New()
	r.validator = v
	return
}

func (i *Validator) Validate(valid interface{}) error {
	return i.validator.Struct(valid)
}

type ValidationError struct {
	ErrorType string   `json:"error_type"`
	Errors    []string `json:"errors"`
}

func (i *ValidationError) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}

func NewValidationError(err string) *ValidationError {
	return &ValidationError{
		Errors: []string{err},
	}
}

func (i *ValidationError) Valid(err error) {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	list := strings.Split(err.Error(), "\n")
	for _, val := range list {
		temp := strings.Split(val, "Key: ")
		temp = strings.Split(temp[1], " Error:")
		snake := matchFirstCap.ReplaceAllString(temp[1], "${1}${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
		i.Errors = append(i.Errors, strings.ToLower(snake))
	}
}

func (i *ValidationError) Bind(err error) {
	dummy := err.Error()
	expected := ""
	if strings.Contains(dummy, "expected=") {
		temp := strings.Split(dummy, "expected=")
		temp = strings.Split(temp[1], ",")
		expected = temp[0]
	}
	got := ""
	if strings.Contains(dummy, "got=") {
		temp := strings.Split(dummy, "got=")
		temp = strings.Split(temp[1], ",")
		got = temp[0]
	}
	field := ""
	if strings.Contains(dummy, "field=") {
		temp := strings.Split(dummy, "field=")
		temp = strings.Split(temp[1], ",")
		field = temp[0]
	}
	if expected != "" && got != "" && field != "" {
		i.Errors = append(i.Errors, "field "+field+" expected "+expected+" but got "+got+".")
	} else {
		i.Errors = append(i.Errors, dummy)
	}
}

func BindAndValidate(c echo.Context, i interface{}) error {
	vErr := new(ValidationError)
	if err := c.Bind(i); err != nil {
		vErr.Bind(err)
		return vErr
	}
	if err := c.Validate(i); err != nil {
		vErr.Valid(err)
		return vErr
	}
	return nil
}
