package vcago

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"runtime"
	"strings"

	"github.com/Viva-con-Agua/vcago/vutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongo represents the initial struct for an Mongo connection.
type Mongo struct {
	Host   string
	Port   string
	DBName string
	DB     *mongo.Database
}

//LoadEnv loads the Host and Port From .env file.
//Host can be set via NATS_HOST
//Port can be set via NATS_PORT
func (i *Mongo) LoadEnv() *Mongo {
	var l vutils.LoadEnv
	i.Host, l = l.GetEnvString("DB_HOST", "w", "localhost")
	i.Port, l = l.GetEnvString("DB_PORT", "w", "27017")
	return i
}

var DB = new(Mongo)

//Connect connects to mongoDB and return an Mongo struct.
func (i *Mongo) Connect(dbName string) (r *Mongo) {
	log.Print("database connecting ...")
	uri := "mongodb://" + i.Host + ":" + i.Port
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	i.DBName = dbName
	i.DB = client.Database(dbName)
	log.Print("database successfully connected!")
	return i
}

//CreateIndex creates an index for a given collection.
func (i *Mongo) CreateIndex(collection string, field string, unique bool) {

	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}
	_, err := i.DB.Collection(collection).Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Print("database failed to create index")
	}
	log.Print("database index created for: " + collection)
}

//InsertOne inserts a value and return an MongoError as error.
func (i Mongo) InsertOne(ctx context.Context, collection string, value interface{}) (err error) {
	_, err = i.DB.Collection(collection).InsertOne(ctx, value)
	if err != nil {
		return NewMongoError(err, value, bson.M{}, i.DBName, collection)
	}
	return
}

//FindOne select an element form database by using a given filter.
//The element will bind to the value interface{}.
//In error case it return an MongoError as error.
func (i *Mongo) FindOne(ctx context.Context, collection string, filter bson.M, value interface{}) (err error) {
	err = i.DB.Collection(collection).FindOne(
		ctx,
		filter,
	).Decode(value)
	if err != nil {
		return NewMongoError(err, value, filter, i.DBName, collection)
	}
	return
}

func (i *Mongo) Find(ctx context.Context, collection string, filter bson.M, value interface{}) (err error) {
	cursor, err := i.DB.Collection(collection).Find(ctx, filter)
	if err != nil {
		return NewMongoError(err, value, filter, i.DBName, collection)
	}
	if err = cursor.All(ctx, value); err != nil {
		return NewMongoError(err, value, filter, i.DBName, collection)
	}
	return
}
func (i *Mongo) Aggregate(ctx context.Context, collection string, filter bson.D, value interface{}) (err error) {
	cursor, err := i.DB.Collection(collection).Aggregate(ctx, filter)
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
func (i *Mongo) UpdateOne(ctx context.Context, collection string, filter bson.M, value interface{}) (err error) {
	result, err := i.DB.Collection(collection).UpdateOne(
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
func (i *Mongo) DeleteOne(ctx context.Context, collection string, filter bson.M) (err error) {
	result, err := i.DB.Collection(collection).DeleteOne(
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

//MongoError represents the struct of an mongo Error type.
type MongoError struct {
	Err        error       `json:"-" bson:"-"`
	Message    string      `json:"message" bson:"message"`
	Filter     interface{} `json:"filter" bson:"filter"`
	Value      interface{} `json:"value" bson:"value"`
	Database   string      `json:"database" bson:"database"`
	Collection string      `json:"collection" bson:"collection"`
	Line       int         `json:"line" bson:"line"`
	File       string      `json:"file" bson:"file"`
}

//NewMongoError creates an mongo Error for an given parameter set
func NewMongoError(err error, value interface{}, filter bson.M, database string, collection string) *MongoError {
	fOutput, _ := json.Marshal(filter)
	temp := new(interface{})
	json.Unmarshal(fOutput, temp)
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()
	return &MongoError{
		Err:        err,
		Message:    err.Error(),
		Filter:     temp,
		Value:      value,
		Database:   database,
		Collection: collection,
		Line:       line,
		File:       file,
	}
}

//Error return string of the error
func (i *MongoError) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}

//Response return the ErrorResponse for handling in httpErrorHandler
func (i *MongoError) Response() (int, interface{}) {
	if strings.Contains(i.Message, "duplicate key error") {
		temp := strings.Split(i.Message, "key: {")
		temp = strings.Split(temp[1], "}")
		return Conflict("duplicate key error", "key: {"+temp[0]+"}")
	}
	switch i.Err {
	case mongo.ErrNoDocuments:
		return NotFound("document not found", i.Filter)
	case ErrMongoUpdate:
		return NotFound("document not updated", i.Filter)
	case ErrMongoDelete:
		return NotFound("document not deleted", i.Filter)
	default:
		return InternalServerError()
	}
}
