package vmdb

import (
	"context"
	"log"
	"net/http"

	"github.com/Viva-con-Agua/vcago"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Collection represents an mongo db database collection
type Collection struct {
	Name         string
	DatabaseName string
	Collection   *mongo.Collection
}

func (i Collection) Response(err error) error {
	if mongo.IsDuplicateKeyError(err) {
		return vcago.NewResp(
			http.StatusConflict,
			"error",
			"duplicate key error",
			i.Name,
			getKeyFromDupKey(err),
		)
	}
	return vcago.NewInternalServerError(i.Name)

	/*switch err.Error() {
	case mongo.ErrNoDocuments:
		return vcago.NewResp(http.StatusNotFound, "error", "document not found", nil, nil).Response()
	case ErrMongoUpdate:
		return NewResp(http.StatusNotFound, "error", "document not updated", nil, nil).Response()
	case ErrMongoDelete:
		return NewResp(http.StatusNotFound, "error", "document not deleted", nil, nil).Response()
	default:
		return NewInternalServerError(i.Collection).Response()

	}*/
}

//CreateIndex creates an index for a given collection.
func (i *Collection) CreateIndex(field string, unique bool) *Collection {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}
	_, err := i.Collection.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Print("database failed to create index")
	}
	log.Print("database index created for: " + i.Name)
	return i
}

//InsertOne inserts a value and return an MongoError as error.
func (i *Collection) InsertOne(ctx context.Context, value interface{}) (err error) {
	_, err = i.Collection.InsertOne(ctx, value)
	if err != nil {
		return
	}
	return
}
