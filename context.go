package vcago

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Context struct {
	Model string
	echo.Context
}

func (i *Context) Ctx() context.Context {
	return i.Request().Context()
}

func (i *Context) BindAndValidate(body interface{}) error {
	vErr := new(ValidationError)
	if err := i.Bind(body); err != nil {
		vErr.Bind(err)
		return vErr
	}
	if err := i.Validate(body); err != nil {
		vErr.Valid(err)
		return vErr
	}
	return nil
}

func (i *Context) AccessToken(token interface{}) (err error) {
	t := i.Get("token")
	if t == nil {
		return errors.New("no token in context")
	}
	temp, ok := t.(*jwt.Token)
	if !ok {
		return errors.New("no jwt.Token type")
	}
	bytes, _ := json.Marshal(temp.Claims)
	_ = json.Unmarshal(bytes, &token)
	return
}

func (i *Context) Created(payload interface{}) (err error) {
	return NewCreated(i.Model, payload)
}

func (i *Context) Selected(payload interface{}) (err error) {
	return NewSelected(i.Model, payload)
}
func (i *Context) Listed(payload interface{}) (err error) {
	return NewSelected(i.Model+"_list", payload)
}
func (i *Context) Updated(payload interface{}) (err error) {
	return NewUpdated(i.Model, payload)
}
func (i *Context) Deleted(payload interface{}) (err error) {
	return NewDeleted(i.Model, payload)
}
