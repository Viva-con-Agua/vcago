package verror

import (
	"errors"
	"strings"

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
		return Conflict("duplicate key error", model)
	}
	switch err {
	case mongo.ErrNoDocuments:
		return NotFound("document_not_found", model)
	case ErrMongoUpdate:
		return NotFound("document_not_updated", model)
	case ErrMongoDelete:
		return NotFound("document_not_deleted", model)
	default:
		return InternalServerError(err)
	}
}

func MongoCreate(err error, model string) error {
	if strings.Contains(err.Error(), "duplicate key error") {
		return Conflict("duplicate key error", model)
	}
	return InternalServerError(err)

}
func MongoGet(err error, model string) error {
	if err == mongo.ErrNoDocuments {
		return NotFound("document_not_found", model)
	}
	return InternalServerError(err)
}

func MongoList(err error, model string) error {
	return InternalServerError(err)
}

func MongoUpdate(err error, model string) error {
	if err == ErrMongoUpdate {
		return NotFound("document_not_updated", model)
	}
	return InternalServerError(err)
}

func MongoDelete(err error, model string) error {
	if err == ErrMongoDelete {
		return NotFound("document_not_deleted", model)
	}
	return InternalServerError(err)
}
