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

// Collection represents an mongo db database collection
type Collection struct {
	Name         string
	DatabaseName string
	Collection   *mongo.Collection
}

type DeletedResult struct {
	Model string `json:"model"`
	ID    string `json:"id"`
}

var (
	FSChunkCollection *Collection
	FSFileCollection  *Collection
)

// CreateIndex creates an index for a given collection.
func (i *Collection) CreateIndex(field string, unique bool) *Collection {
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}
	_, err := i.Collection.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// CreateMultiIndex creates an index for an value combination.
func (i *Collection) CreateMultiIndex(filter bson.D, unique bool) *Collection {
	mod := mongo.IndexModel{
		Keys:    filter,
		Options: options.Index().SetUnique(unique),
	}
	_, err := i.Collection.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Fatal(err)
	}
	return i

}

// InsertOne inserts a value and return an MongoError as error.
func (i *Collection) InsertOne(ctx context.Context, value interface{}) (err error) {
	_, err = i.Collection.InsertOne(ctx, value)
	if err != nil {
		return i.log(err)
	}
	return
}

// InsertMany inserts a list of value and return an MongoError as error.
func (i *Collection) InsertMany(ctx context.Context, value []interface{}) (err error) {
	if len(value) != 0 {
		_, err = i.Collection.InsertMany(ctx, value)
		if err != nil {
			return i.log(err)
		}
	}
	return
}

// FindOne use the mongo.Collection.FindOne function for select one element from collection.
func (i *Collection) FindOne(ctx context.Context, filter bson.D, value interface{}) (err error) {
	if err = i.Collection.FindOne(ctx, filter).Decode(value); err != nil {
		return i.log(err)
	}
	return
}

// AggregateOne use an aggregation pipeline for creating an struct that contains the information from more than one collection.
// If the result cursor contains objects, the first one will be decoded in the value param.
func (i *Collection) AggregateOne(ctx context.Context, pipeline mongo.Pipeline, value interface{}, opts ...*options.AggregateOptions) (err error) {
	var cursor *mongo.Cursor
	cursor, err = i.Collection.Aggregate(ctx, pipeline, opts...)
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

// Find use the mongo.Collection.Find function for select a list of elements from a collection.
// The result will decode in the value param. So the value need to be a slice struct.
func (i *Collection) Find(ctx context.Context, filter bson.D, value interface{}, opts ...*options.FindOptions) (err error) {
	var cursor *mongo.Cursor
	cursor, err = i.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return i.log(err)
	}
	if err = cursor.All(ctx, value); err != nil {
		return i.log(err)
	}
	return
}

// FindAndCount use the mongo.Collection.Find function for select a list of elements from a collection.
// The result will decode in the value param. So the value need to be a slice struct.
// The return value listSize counts all elements in the collection that match the given filter.
func (i *Collection) FindAndCount(ctx context.Context, filter bson.D, value interface{}, opts ...*options.FindOptions) (listSize int64, err error) {
	var cursor *mongo.Cursor
	cursor, err = i.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return 0, i.log(err)
	}
	if err = cursor.All(ctx, value); err != nil {
		return 0, i.log(err)
	}
	opts_count := options.Count().SetHint("_id_")
	if cursor, cErr := i.Collection.CountDocuments(ctx, filter, opts_count); cErr != nil {
		listSize = 0
	} else {
		listSize = cursor
	}
	return
}

// Aggregate use the mongo.Collection.Aggregate function for select a list of elements using an aggregation pipeline.
func (i *Collection) Aggregate(ctx context.Context, filter mongo.Pipeline, value interface{}, opts ...*options.AggregateOptions) (err error) {
	cursor, err := i.Collection.Aggregate(ctx, filter, opts...)
	if err != nil {
		return i.log(err)
	}
	if err = cursor.All(ctx, value); err != nil {
		return i.log(err)
	}
	return
}

// UpdateOne use the mongo.Collection.UpdateOne function for update one element in an collection.
// If the result.MatchedCount == 0, the function returns an mongo.ErrNoDocuments error.
// If the result.MatchedCount != 0, the i.Collection.FindOne function is used to select the updated element and decode it into the result interface.
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

// UpdateOneAggregate works the same way than UpdateOne but you can define the pipeline param for decode an aggregated model into the result interface.
func (i *Collection) UpdateOneAggregate(ctx context.Context, filter bson.D, value interface{}, result interface{}, pipeline mongo.Pipeline) (err error) {
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
		if err = i.AggregateOne(ctx, pipeline, result); err != nil {
			return
		}
	}
	return
}

// TryUpdateOne returns no error if the model is not updated.
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

// UpdateMany updates an slice of interfaces.
// @TODO: create an result.
func (i *Collection) UpdateMany(ctx context.Context, filter bson.D, value bson.D) (err error) {
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

// TryUpdateMany returns no error if no models was updated.
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

// DeleteOne deletes an element from given collection by the bson.M filter.
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

// DeleteMany deletes all elements from collection they match the filter object.
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

// DeleteOne deletes an element from given collection by the bson.M filter.
func (i *Collection) TryDeleteOne(ctx context.Context, filter bson.D) (err error) {
	_, err = i.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	return
}

// TryDeleteMany returns no error if no element was deleted.
func (i *Collection) TryDeleteMany(ctx context.Context, filter bson.D) (err error) {
	_, err = i.Collection.DeleteMany(ctx, filter)
	if err != nil {
		return i.log(err)
	}
	return
}

// CountDocuments count all documents filtered by the filter object.
func (i *Collection) CountDocuments(ctx context.Context, filter bson.D) (count int64, err error) {
	count, err = i.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, i.log(err)
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

// ErrNoDocuments return true if the error is an mongo.ErrNoDocuments error. Else the function returns false.
func ErrNoDocuments(err error) bool {
	if err != nil {
		e := err.(*vcago.Error)
		return e.Err == mongo.ErrNoDocuments
	}
	return false
}

func ErrDuplicateKey(err error) bool {
	if err != nil {
		e := err.(*vcago.Error)
		if strings.Contains(e.Err.Error(), "duplicate key error") {
			return true
		}
	}
	return false
}
