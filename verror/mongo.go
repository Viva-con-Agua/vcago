package verror

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

//ErrMongoUpdate represents an update error in mongo case
var ErrMongoUpdate = errors.New("no updated document")

//ErrMongoDelete represents an delete error in mongo case
var ErrMongoDelete = errors.New("no delete document")

//MongoUpdateErr helper for handling update errors
func MongoUpdateErr(err error, result *mongo.UpdateResult) error {
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrMongoUpdate
	}
	return nil
}

//MongoDeleteErr helper for handling delete errors
func MongoDeleteErr(err error, result *mongo.DeleteResult) error {
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrMongoUpdate
	}
	return nil
}
func Mongo(err error, model string) error {
	if strings.Contains(err.Error(), "duplicate key error") {
		return echo.NewHTTPError(http.StatusConflict, &ErrorResponse{Message: "duplicate key error", Model: model})
	}
	if err == mongo.ErrNoDocuments {
		return echo.NewHTTPError(http.StatusNotFound, &ErrorResponse{Message: "document_not_found", Model: model})

	}
	if err == ErrMongoUpdate {
		return echo.NewHTTPError(http.StatusNotFound, &ErrorResponse{Message: "document_not_updated", Model: model})
	}
	if err == ErrMongoDelete {
		return echo.NewHTTPError(http.StatusNotFound, &ErrorResponse{Message: "document_not_deleted", Model: model})

	}
	return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error", Model: model})
}

func MongoCreate(err error, model string) error {
	if strings.Contains(err.Error(), "duplicate key error") {
		return echo.NewHTTPError(http.StatusConflict, &ErrorResponse{Message: "duplicate key error", Model: model})
	}
	return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error", Model: model})
}
func MongoGet(err error, model string) error {
	if err == mongo.ErrNoDocuments {
		return echo.NewHTTPError(http.StatusNotFound, &ErrorResponse{Message: "document_not_found", Model: model})
	}
	return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error", Model: model})
}

func MongoList(err error, model string) error {
	return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error", Model: model})
}

func MongoUpdate(err error, model string) error {
	if err == ErrMongoUpdate {
		return echo.NewHTTPError(http.StatusNotFound, &ErrorResponse{Message: "document_not_updated", Model: model})
	}
	return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error", Model: model})
}

func MongoDelete(err error, model string) error {
	if err == ErrMongoDelete {
		return echo.NewHTTPError(http.StatusNotFound, &ErrorResponse{Message: "document_not_deleted", Model: model})
	}
	return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error", Model: model})
}
