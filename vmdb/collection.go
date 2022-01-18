package vmdb

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Collection represents an mongo db database collection
type Collection struct {
	Name     string
	Database *Database
}

//NewCollection return a new Collection.
func NewCollection(name string) *Collection {
	return &Collection{
		Name:     name,
		Database: DB,
	}
}

//CreateIndex creates an index for a given collection.
func (i *Collection) CreateIndex(field string, unique bool) {

	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}
	_, err := i.Database.DB.Collection(i.Name).Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Print("database failed to create index")
	}
	log.Print("database index created for: " + i.Name)
}

//FindOne select an element form database by using a given filter.
//The element will bind to the value interface{}.
//In error case it return an MongoError as error.
func (i *Collection) FindOne(ctx context.Context, filter bson.M, value interface{}) (err error) {
	err = i.Database.DB.Collection(i.Name).FindOne(
		ctx,
		filter,
	).Decode(value)
	if err != nil {
		return NewMongoError(err, value, filter, i.Database.Name, i.Name)
	}
	return
}

//Find return a list of objects by given filter
func (i *Collection) Find(ctx context.Context, filter bson.M, value interface{}) (err error) {
	cursor, err := i.Database.DB.Collection(i.Name).Find(ctx, filter)
	if err != nil {
		return NewMongoError(err, value, filter, i.Database.Name, i.Name)
	}
	if err = cursor.All(ctx, value); err != nil {
		return NewMongoError(err, value, filter, i.Database.Name, i.Name)
	}
	return
}

//Aggregate provide using aggregations.
func (i *Collection) Aggregate(ctx context.Context, filter bson.D, value interface{}) (err error) {
	cursor, err := i.Database.DB.Collection(i.Name).Aggregate(ctx, filter)
	if err != nil {
		return
	}
	if err = cursor.All(ctx, value); err != nil {
		return
	}
	return
}

//ErrMongoUpdate represents an update error in mongo case
var ErrMongoUpdate = errors.New("no updated document")

//UpdateOne updates a value via "$set" and the given bson.M filter. Return an MongoError in case that no element has updated.
func (i *Collection) UpdateOne(ctx context.Context, filter bson.M, value interface{}) (err error) {
	result, err := i.Database.DB.Collection(i.Name).UpdateOne(
		ctx,
		filter,
		bson.M{"$set": value},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrMongoUpdate
	}
	return
}

//ErrMongoDelete represents an delete error in mongo case
var ErrMongoDelete = errors.New("no delete document")

//DeleteOne deletes an element from given collection by the bson.M filter.
func (i *Collection) DeleteOne(ctx context.Context, filter bson.M) (err error) {
	result, err := i.Database.DB.Collection(i.Name).DeleteOne(
		ctx,
		bson.M{"_id": filter},
	)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrMongoDelete
	}
	return
}
