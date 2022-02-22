package vcago

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoColl represents an mongo db database collection
type MongoColl struct {
	Name         string
	DatabaseName string
	Collection   *mongo.Collection
}

//CreateIndex creates an index for a given collection.
func (i *MongoColl) CreateIndex(field string, unique bool) *MongoColl {
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
func (i *MongoColl) InsertOne(ctx context.Context, value interface{}) (err error) {
	_, err = i.Collection.InsertOne(ctx, value)
	if err != nil {
		return NewMongoError(err, value, bson.M{}, i.DatabaseName, i.Name)
	}
	return
}

//FindOne select an element form database by using a given filter.
//The element will bind to the value interface{}.
//In error case it return an MongoError as error.
func (i *MongoColl) FindOne(ctx context.Context, filter bson.M, value interface{}) (err error) {
	err = i.Collection.FindOne(
		ctx,
		filter,
	).Decode(value)
	if err != nil {
		return NewMongoError(err, value, filter, i.DatabaseName, i.Name)
	}
	return
}

//Find return a list of objects by given filter
func (i *MongoColl) Find(ctx context.Context, filter bson.M, value interface{}) (err error) {
	cursor, err := i.Collection.Find(ctx, filter)
	if err != nil {
		return NewMongoError(err, value, filter, i.DatabaseName, i.Name)
	}
	if err = cursor.All(ctx, value); err != nil {
		return NewMongoError(err, value, filter, i.DatabaseName, i.Name)
	}
	return
}

//Aggregate provide using aggregations.
func (i *MongoColl) Aggregate(ctx context.Context, filter []bson.D, value interface{}) (err error) {
	cursor, err := i.Collection.Aggregate(ctx, filter)
	if err != nil {
		return NewMongoError(err, value, bson.M{"filter": filter}, i.DatabaseName, i.Name)
	}
	if err = cursor.All(ctx, value); err != nil {
		return NewMongoError(err, value, bson.M{"filter": filter}, i.DatabaseName, i.Name)
	}
	return
}

//ErrMongoUpdate represents an update error in mongo case
var ErrMongoUpdate = errors.New("no updated document")

//UpdateOne updates a value via "$set" and the given bson.M filter. Return an MongoError in case that no element has updated.
func (i *MongoColl) UpdateOne(ctx context.Context, filter bson.M, value interface{}) (err error) {
	result, err := i.Collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": value},
	)
	if err != nil {
		return NewMongoError(err, value, filter, i.DatabaseName, i.Name)
	}
	if result.MatchedCount == 0 {
		return NewMongoError(ErrMongoUpdate, value, filter, i.DatabaseName, i.Name)
	}
	return
}

//ErrMongoDelete represents an delete error in mongo case
var ErrMongoDelete = errors.New("no delete document")

//DeleteOne deletes an element from given collection by the bson.M filter.
func (i *MongoColl) DeleteOne(ctx context.Context, filter bson.M) (err error) {
	result, err := i.Collection.DeleteOne(
		ctx,
		filter,
	)
	if err != nil {
		return NewMongoError(err, nil, filter, i.DatabaseName, i.Name)
	}
	if result.DeletedCount == 0 {
		return NewMongoError(ErrMongoDelete, nil, filter, i.DatabaseName, i.Name)
	}
	return
}

//InsertOrUpdate updates a value via "$set" and the given bson.M filter. Return an MongoError in case that no element has updated.
func (i *MongoColl) InsertOrUpdate(ctx context.Context, filter bson.M, value interface{}) (err error) {
	result, err := i.Collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": value},
	)
	if err != nil {
		return NewMongoError(err, value, filter, i.DatabaseName, i.Name)
	}
	if result.MatchedCount == 0 {
		_, err = i.Collection.InsertOne(ctx, value)
		if err != nil {
			return NewMongoError(err, value, filter, i.DatabaseName, i.Name)
		}
		return
	}
	return
}
