package vmdb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongo represents the initial struct for an Mongo connection.
type Database struct {
	Host     string
	Port     string
	Name     string
	Database *mongo.Database
}

//NewMongoDB creates a new MongoDB connects to mongoDB and return an Mongo struct.
func NewDatabase(name string, host string, port string) (r *Database) {
	r = new(Database)
	r.Name = name
	uri := "mongodb://" + host + ":" + port
	log.Print("MongoDB connection to " + uri)
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
	r.Database = client.Database(r.Name)
	log.Print("MongoDB successfully connected!")
	return
}

func (i *Database) Collection(name string) *Collection {
	return &Collection{
		Name:         name,
		DatabaseName: i.Name,
		Collection:   i.Database.Collection(name),
	}
}
