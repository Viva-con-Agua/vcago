package verr

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type (
    //MongoConflict represents the reponse type for mongo database Conficts.
    MongoConflict struct {
        Message string `json:"message"`
        Coll string `json:"collection"`
        Value string `json:"value"`
    }
    //MongoBase represents response type basic mongo error with collection involved.
    MongoBase struct {
        Message string `json:"message"`
        Coll string `json:"collection,omitempty"`
    }
)

//MongoInsertOneError handles error response and logging for collection.insertOne() function.
func MongoInsertOneError(ctx context.Context, err error, coll string) error {
    if strings.Contains(err.Error(), "duplicate key error"){
        LogError(ctx, err, "debug")
        return echo.NewHTTPError(http.StatusConflict, handleDuplicateKey(err))
    }
    LogError(ctx, err, "prod")
    return InternalServerErrorMsg
} 
//MongoFindOneError handles error response and logging for collection.findOne() function.
func MongoFindOneError(ctx context.Context, err error, coll string) error {
    if err == mongo.ErrNoDocuments {
        LogError(ctx, err, "debug")
        return echo.NewHTTPError(http.StatusNotFound, MongoBase{Message: "not_found", Coll: coll})
    }
    LogError(ctx, err, "prod")
    return InternalServerErrorMsg
}
//MongoUpdateOneError handles error response and logging for collection.deleteOne() function.
func MongoUpdateOneError(ctx context.Context, err error, coll string, result *mongo.UpdateResult) error {
    if err != nil {
        LogError(ctx, err, "prod")
        return InternalServerErrorMsg
    }
    if result.MatchedCount == 0 {
        LogError(ctx, errors.New("no updated document"), "debug")
        return echo.NewHTTPError(http.StatusNotFound, MongoBase{Message: "not_found", Coll: coll})
    }
    return nil
}

//MongoDeleteOneError handles error response and logging for collection.deleteOne() function.
func MongoDeleteOneError(ctx context.Context, err error, coll string, result *mongo.DeleteResult) error {
    if err != nil {
        LogError(ctx, err, "prod")
        return InternalServerErrorMsg
    }
    if result.DeletedCount == 0 {
        LogError(ctx, errors.New("no deleted document"), "debug")
        return echo.NewHTTPError(http.StatusNotFound, MongoBase{Message: "not_found", Coll: coll})
    }
    return nil
}
//handleDublicateKey parse an mongo duplicate key error and return MongoConflict
func handleDuplicateKey(err error) *MongoConflict {
		cut1 := strings.Split(err.Error(), "collection: ")
		cut2 := strings.Split(cut1[1], " index: ") 
		cut3 := strings.Split(cut2[1], "dup")
		value := strings.Split(cut3[0], "_")
		coll := strings.Split(cut2[0], ".")
        return &MongoConflict{
            Message: "duplicate key error",
            Coll: coll[1],
            Value: value[0],
        }
}
