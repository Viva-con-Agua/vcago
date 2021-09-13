package vutils

import (
	"context"
	"encoding/json"
	"log"

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
	var l LoadEnv
	i.Host, l = l.GetEnvString("DB_HOST", "w", "localhost")
	i.Port, l = l.GetEnvString("DB_PORT", "w", "27017")
	return i
}

//Connect connects to mongoDB and return an mongo.Database.
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

func (i Mongo) InsertOne(ctx context.Context, collection string, value interface{}) (err error) {
	_, err = i.DB.Collection(collection).InsertOne(ctx, value)
	if err != nil {
		return NewMongoError(err, value, i.DBName, collection)
	}
	return
}

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

type MongoError struct {
	Err        error       `json:"error" bson:"error"`
	Value      interface{} `json:"value" bson:"value"`
	Database   string      `json:"database" bson:"database"`
	Collection string      `json:"collection" bson:"collection"`
}

func NewMongoError(err error, value interface{}, database string, collection string) *MongoError {
	return &MongoError{
		Err:        err,
		Value:      value,
		Database:   database,
		Collection: collection,
	}
}

func (i *MongoError) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}
