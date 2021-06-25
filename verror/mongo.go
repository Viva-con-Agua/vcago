package verror

import (
	"errors"

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
