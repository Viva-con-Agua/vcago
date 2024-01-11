package vmdb

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type CollectionUpdate struct {
	Collection *Collection
}

type CollectionUpdateModel struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func NewCollectionUpdate(collection *Collection) *CollectionUpdate {
	return &CollectionUpdate{Collection: collection}
}

func (i CollectionUpdate) Insert(ctx context.Context, name string) {
	insert := &CollectionUpdateModel{ID: uuid.NewString(), Name: name}
	if err := i.Collection.InsertOne(ctx, insert); err != nil {
		log.Print(err)
	}
}

func (i CollectionUpdate) Check(ctx context.Context, name string) bool {
	check := new(CollectionUpdateModel)
	if err := i.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}, &check); err != nil {
		if ErrNoDocuments(err) {
			return false
		}
		log.Print(err)
	}
	return true
}
