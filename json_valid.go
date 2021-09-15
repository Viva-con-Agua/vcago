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

//Validate extend JSONValidator with Validate function.

func (i *Validator) New(v *validator.Validate) *Validator {
	i.validator = v
	return i
}

func (i *Validator) Validate(valid interface{}) error {
	return i.validator.Struct(valid)
}

type ValidationError struct {
	Errors []string `json:"errors"`
}

func (i *ValidationError) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
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
	i.Errors = append(i.Errors, err.Error())
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