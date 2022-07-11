package vcago

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

//Context represents an extended echo.Context models used for basic functions in a Handler.
type Context struct {
	Model string
	echo.Context
}

//Ctx return the context.Context model of the request.
func (i *Context) Ctx() context.Context {
	return i.Request().Context()
}

//BindAndValidate binds the request data in an the body interface.
//Define param:"" for bind the params in a struct.
//Define json:"" for bind the request body as json in a struct.
//Define query:"" for bind the query parameters in a struct.
func (i *Context) BindAndValidate(body interface{}) error {
	if err := i.Bind(body); err != nil {
		return NewError(err, "DEBUG", "bind").AddModel(i.Model)
	}
	if err := i.Validate(body); err != nil {
		return NewError(err, "DEBUG", "validate").AddModel(i.Model)
	}
	return nil
}

//AccessToken binds the accessToken form an cookie into an struct.
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
	err = json.Unmarshal(bytes, &token)
	return
}

//ErrorResonse return an http error for an given error.
func (i *Context) ErrorResponse(err error) error {
	if mongo.IsDuplicateKeyError(err) {
		return NewResp(
			http.StatusConflict,
			"error",
			"duplicate key error",
			i.Model,
			getKeyFromDupKey(err),
		)
	} else if err == mongo.ErrNoDocuments {
		return NewNotFound(i.Model, nil)
	}
	return NewInternalServerError(i.Model)
}

//Log return a function call for handling Debug and Error logs.
//Usage: c.Log(err)(err)
func (i *Context) Log(err error) {
	id := i.Request().Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = i.Response().Header().Get(echo.HeaderXRequestID)
	}
	var output *Error
	if err == mongo.ErrNoDocuments || strings.Contains(err.Error(), "duplicate key error") {
		output = NewError(err, "DEBUG", "")
	} else {
		output = NewError(err, "ERROR", "")
	}
	output.Print(id)
}

//Created returns an Created response.
func (i *Context) Created(payload interface{}) (err error) {
	return i.JSON(NewCreated(i.Model, payload).Response())
}

//Selected returns an Selected response.
func (i *Context) Selected(payload interface{}) (err error) {
	return i.JSON(NewSelected(i.Model, payload).Response())
}

//Listed return an List response.
func (i *Context) Listed(payload interface{}) (err error) {
	return i.JSON(NewSelected(i.Model+"_list", payload).Response())
}

//Updated returns an Updated response
func (i *Context) Updated(payload interface{}) (err error) {
	return i.JSON(NewUpdated(i.Model, payload).Response())
}

//Deleted returns an Deleted response.
func (i *Context) Deleted(payload interface{}) (err error) {
	return i.JSON(NewDeleted(i.Model, payload).Response())
}

func (i *Context) SuccessResponse(status int, message string, model string, payload interface{}) (err error) {
	return i.JSON(NewResp(status, "success", message, model, payload).Response())
}

func getCollectionFromDupKey(err error) string {
	temp := strings.Split(err.Error(), "collection: ")
	temp = strings.Split(temp[1], " key:")
	return temp[0]
}

func getKeyFromDupKey(err error) string {
	temp := strings.Split(err.Error(), "key: {")
	temp = strings.Split(temp[1], "}")
	return temp[0]
}
