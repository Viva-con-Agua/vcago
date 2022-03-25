package vcago

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

//HTTPErrorHandler handles echo.HTTPError and return the correct response.
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if resp, ok := err.(*Response); ok {
		c.JSON(resp.Response())
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Code == http.StatusInternalServerError {
			c.JSON(NewInternalServerError("-").Response())
		} else {
			c.JSON(code, err)
		}
	} else if resp, ok := err.(*MongoError); ok {
		c.JSON(MongoErrorResponseHandler(resp))
	} else if resp, ok := err.(*ValidationError); ok {
		c.JSON(ValidationErrorResponseHandler(resp))
	} else {
		c.JSON(NewInternalServerError("-").Response())
	}
}

//ValidationErrorResponseHandler handles the response for the ValidationError type.
func ValidationErrorResponseHandler(i *ValidationError) (int, interface{}) {
	return NewBadRequest("-", "validation error", i).Response()
}

//MongoErrorResponseHandler handles the response for the MongoError type.
func MongoErrorResponseHandler(i *MongoError) (int, interface{}) {
	if strings.Contains(i.Message, "duplicate key error") {
		temp := strings.Split(i.Message, "key: {")
		temp = strings.Split(temp[1], "}")
		return NewResp(http.StatusConflict, "error", "duplicate key error", i.Collection, temp[0]).Response()
	}
	switch i.Err {
	case mongo.ErrNoDocuments:
		return NewResp(http.StatusNotFound, "error", "document not found", i.Collection, i.Filter).Response()
	case ErrMongoUpdate:
		return NewResp(http.StatusNotFound, "error", "document not updated", i.Collection, i.Filter).Response()
	case ErrMongoDelete:
		return NewResp(http.StatusNotFound, "error", "document not deleted", i.Collection, i.Filter).Response()
	default:
		return NewInternalServerError(i.Collection).Response()
	}
}
