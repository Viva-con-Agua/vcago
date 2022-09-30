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

// Context extends an echo.Context type. Inititial an Context for an echo.Group via
// echo.Group.Use(Handler.Context()). The Context model is set by the Model param of the Handler model.
//
// Example:
//
//	func (i *ExampleHandler) Routes(group echo.Group) {
//		group.Use(i.Context)
//	}
//
// Echo need and function that provide an echo.Context as parameter.
// So you need to convert the echo.Context to an Context in a function block.
//
// Example:
//
//	func (i *ExampleHandler) Create(cc echo.Context) (err error) {
//		c := cc.(vcago.Context)
//		...
//	}
type Context struct {
	Model string
	echo.Context
}

// Ctx return echo.Context.Context().
func (i *Context) Ctx() context.Context {
	return i.Request().Context()
}

// BindAndValidate binds the request data in an the body interface.
// Define param:"" for bind the params in a struct.
// Define json:"" for bind the request body as json in a struct.
// Define query:"" for bind the query parameters in a struct.
func (i *Context) BindAndValidate(body interface{}) error {
	if err := i.Bind(body); err != nil {
		return NewError(err, "DEBUG", "bind").AddModel(i.Model)
	}
	if err := i.Validate(body); err != nil {
		return NewError(err, "DEBUG", "validation").AddModel(i.Model)
	}
	return nil
}

// AccessToken binds the accessToken form an cookie into the token interface.
// The token needs to extends jwt.StandardClaims.
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

// RefreshToken returns the user id of an refresh token.
func (i *Context) RefreshTokenID() (string, error) {
	token := i.Get("token").(*jwt.Token)
	if token == nil {
		return "", errors.New("No user in Conext")
	}
	return token.Claims.(*RefreshToken).UserID, nil
}

// ErrorResonse match the error and returns the correct error response.
//
// Error: mongo.IsDuplicateKeyError
//
// Status: 409 Conflict
//
// JSON:
//
//	{
//		"type": "error",
//		"message": "duplicate_key_error",
//		"model": model,
//		"payload": key from error
//	}
//
// Error: mongo.ErrNoDocument
//
// Status: 404 Not Found
//
// JSON:
//
//	{
//		"type": "error",
//		"message": "not_found",
//		"model": model
//	}
//
// Error: default
//
// Status: 500 Internal Server Error
//
// JSON:
//
//	{
//		"type": "error",
//		"message": "internal_server_error",
//		"model": model
//	}
func (i *Context) ErrorResponse(err error) error {
	if e, ok := err.(*Error); ok {
		return i.JSON(e.Response())
	}
	return i.JSON(NewInternalServerError(i.Model).Response())
}

// Log return a function call for handling Debug and Error logs.
// Usage: c.Log(err)(err)
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

// Created returns an Created response.
//
// Status: 201 Created
//
// JSON:  model == "example"
//
//	{
//		"type": "success",
//		"message": "successfully_created",
//		"model": "example",
//		"payload": payload
//	}
func (i *Context) Created(payload interface{}) (err error) {
	return i.JSON(NewCreated(i.Model, payload).Response())
}

// Selected returns an Selected response.
//
// Status: 200 OK
//
// JSON:  model == "example"
//
//	{
//		"type": "success",
//		"message": "successfully_selected",
//		"model": "example",
//		"payload": payload
//	}
func (i *Context) Selected(payload interface{}) (err error) {
	return i.JSON(NewSelected(i.Model, payload).Response())
}

// Listed return an List response.
//
// Status: 200 OK
//
// JSON:  model == "example"
//
//	{
//		"type": "success",
//		"message": "successfully_selected",
//		"model": "example_list",
//		"payload": payload
//	}
func (i *Context) Listed(payload interface{}) (err error) {
	return i.JSON(NewSelected(i.Model+"_list", payload).Response())
}

// Updated returns an Updated response
//
// Status: 200 OK
//
// JSON:  model == "example"
//
//	{
//		"type": "success",
//		"message": "successfully_updated",
//		"model": "example",
//		"payload": payload
//	}
func (i *Context) Updated(payload interface{}) (err error) {
	return i.JSON(NewUpdated(i.Model, payload).Response())
}

// Deleted returns an Deleted response.
//
// Status: 200 OK
//
// JSON:  model == "example"
//
//	{
//		"type": "success",
//		"message": "successfully_deleted",
//		"model": "example",
//		"payload": payload
//	}
func (i *Context) Deleted(payload interface{}) (err error) {
	return i.JSON(NewDeleted(i.Model, payload).Response())
}

// SuccessResponse returns an new success 200 OK response with an custom message string.
//
// Status: 200 OK
//
// JSON:  model == "example"
//
//	{
//		"type": "success",
//		"message": message,
//		"model": "example",
//		"payload": payload
//	}
func (i *Context) SuccessResponse(status int, message string, model string, payload interface{}) (err error) {
	return i.JSON(NewResp(status, "success", message, model, payload).Response())
}

func (i *Context) BadRequest(message string, payload interface{}) (err error) {
	return i.JSON(NewResp(http.StatusBadRequest, "bad_request", message, i.Model, payload).Response())
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
