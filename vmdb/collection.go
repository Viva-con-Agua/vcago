package vmdb

import (
	"context"
	"log"
	"runtime"
	"strings"
	"time"

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
		return i.log(err)
	}
	return
}

//InsertMany inserts a list of value and return an MongoError as error.
func (i *Collection) InsertMany(ctx context.Context, value []interface{}) (err error) {
	_, err = i.Collection.InsertMany(ctx, value)
	if err != nil {
		return i.log(err)
	}
	return
}

func (i *Collection) FindOne(ctx context.Context, filter bson.D, value interface{}) (err error) {
	if err = i.Collection.FindOne(ctx, filter).Decode(value); err != nil {
		return i.log(err)
	}
	return
}

func (i *Collection) AggregateOne(ctx context.Context, pipeline mongo.Pipeline, value interface{}) (err error) {
	var cursor *mongo.Cursor
	cursor, err = i.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return i.log(err)
	}
	if cursor.TryNext(ctx) {
		if err = cursor.Decode(value); err != nil {
			return i.log(err)
		}
	} else {
		return i.log(mongo.ErrNoDocuments)
	}
	return
}

func (i *Collection) Find(ctx context.Context, filter bson.D, value interface{}) (err error) {
	var cursor *mongo.Cursor
	cursor, err = i.Collection.Find(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	if err = cursor.All(ctx, value); err != nil {
		return i.log(err)
	}
	return
}

//Aggregate provide using aggregations.
func (i *Collection) Aggregate(ctx context.Context, filter mongo.Pipeline, value interface{}) (err error) {
	cursor, err := i.Collection.Aggregate(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	if err = cursor.All(ctx, value); err != nil {
		return i.log(err)
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
		return i.log(err)
	}
	if updateResult.MatchedCount == 0 {
		return i.log(mongo.ErrNoDocuments)
	}
	if result != nil {
		if err = i.Collection.FindOne(ctx, filter).Decode(result); err != nil {
			return i.log(err)
		}
	}
	return
}

func (i *Collection) TryUpdateOne(ctx context.Context, filter bson.D, value interface{}) (err error) {
	_, err = i.Collection.UpdateOne(
		ctx,
		filter,
		value,
	)
	if err != nil {
		return i.log(err)
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
		return i.log(err)
	}
	if result.MatchedCount == 0 {
		return i.log(mongo.ErrNoDocuments)
	}
	return

}

func (i *Collection) TryUpdateMany(ctx context.Context, filter bson.A, value bson.M) (err error) {
	_, err = i.Collection.UpdateMany(
		ctx,
		filter,
		value,
	)
	if err != nil {
		return i.log(err)
	}
	return
}

//DeleteOne deletes an element from given collection by the bson.M filter.
func (i *Collection) DeleteOne(ctx context.Context, filter bson.D) (err error) {
	result, err := i.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	if result.DeletedCount == 0 {
		return i.log(mongo.ErrNoDocuments)
	}
	return
}

func (i *Collection) DeleteMany(ctx context.Context, filter bson.D) (err error) {
	result, err := i.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	if result.DeletedCount == 0 {
		return i.log(mongo.ErrNoDocuments)
	}
	return
}

//DeleteOne deletes an element from given collection by the bson.M filter.
func (i *Collection) TryDeleteOne(ctx context.Context, filter bson.D) (err error) {
	_, err = i.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	return
}

func (i *Collection) TryDeleteMany(ctx context.Context, filter bson.D) (err error) {
	_, err = i.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	return
}

func (i *Collection) log(err error) error {
	var lvl = "ERROR"
	if strings.Contains(err.Error(), "duplicate key error") {
		lvl = "DEBUG"
	} else if err == mongo.ErrNoDocuments {
		lvl = "DEBUG"
	}
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()
	return &vcago.Error{
		Time:    time.Now().String(),
		Level:   lvl,
		File:    file,
		Line:    line,
		Message: err.Error(),
		Model:   i.Name,
		Err:     err,
		Type:    "mongo",
	}
}

func ErrNoDocuments(err error) bool {
	e := err.(*vcago.Error)
	if e.Err == mongo.ErrNoDocuments {
		return true
	}
	return false
}
