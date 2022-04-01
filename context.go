package vcago

import (
	"context"
	"errors"
	"log"

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
	log.Print(t)
	token(t.(*jwt.Token).Claims.())
	log.Print(token)
	return
}

func (i *Context) Created(payload interface{}) (err error) {
	return NewCreated(i.Model, payload)
}
