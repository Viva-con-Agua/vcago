package vmdb

import (
	"context"
	"log"

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

type DeletedResult struct {
	Model string `json:"model"`
	ID    string `json:"id"`
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

func (i *Collection) CreateMultiIndex(filter bson.D, unique bool) *Collection {
	mod := mongo.IndexModel{
		Keys:    filter,
		Options: options.Index().SetUnique(unique),
	}
	_, err := i.Collection.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Print("database failed to create index " + i.Name)
	} else {
		log.Print("database index created for: " + i.Name)
	}
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

//InsertMany inserts a list of value and return an MongoError as error.
func (i *Collection) InsertMany(ctx context.Context, value []interface{}) (err error) {
	_, err = i.Collection.InsertMany(ctx, value)
	if err != nil {
		return
	}
	return
}

func (i *Collection) FindOne(ctx context.Context, pipeline mongo.Pipeline, value interface{}) (err error) {
	var cursor *mongo.Cursor
	cursor, err = i.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return
	}
	if cursor.TryNext(ctx) {
		if err = cursor.Decode(value); err != nil {
			return
		}
	} else {
		return mongo.ErrNoDocuments
	}
	return
}

func (i *Collection) Find(ctx context.Context, pipeline mongo.Pipeline, value interface{}) (err error) {
	var cursor *mongo.Cursor
	cursor, err = i.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return
	}
	if err = cursor.All(ctx, value); err != nil {
		return
	}
	return
}

//Aggregate provide using aggregations.
func (i *Collection) Aggregate(ctx context.Context, filter []bson.D, value interface{}) (err error) {
	cursor, err := i.Collection.Aggregate(ctx, filter)
	if err != nil {
		return
	}
	if err = cursor.All(ctx, value); err != nil {
		return
	}
	return
}

func (i *Collection) UpdateOne(ctx context.Context, filter bson.D, value interface{}, result interface{}) (err error) {
	updateResult, err := i.Collection.UpdateOne(
		ctx,
		filter,
		value,
	)
	if err != nil {
		return
	}
	if updateResult.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	if result != nil {
		if err = i.Collection.FindOne(ctx, filter).Decode(result); err != nil {
			return
		}
	}
	return
}

func (i *Collection) UpdateMany(ctx context.Context, filter bson.A, value bson.M) (err error) {
	result, err := i.Collection.UpdateMany(
		ctx,
		filter,
		value,
	)
	if err != nil {
		return
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return

}

//DeleteOne deletes an element from given collection by the bson.M filter.
func (i *Collection) DeleteOne(ctx context.Context, filter bson.D) (err error) {
	result, err := i.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return
}

func (i *Collection) DeleteMany(ctx context.Context, filter bson.M) (err error) {
	result, err := i.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return
}
