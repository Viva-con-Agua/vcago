package vcago

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mongo represents the initial struct for an Mongo connection.
type MongoDB struct {
	Host     string
	Port     string
	Name     string
	Database *mongo.Database
}

//NewMongoDB creates a new MongoDB connects to mongoDB and return an Mongo struct.
func NewMongoDB(name string) (r *MongoDB) {
	r = new(MongoDB)
	r.Host = Config.GetEnvString("DB_HOST", "w", "localhost")
	r.Port = Config.GetEnvString("DB_PORT", "w", "27017")
	r.Name = name
	uri := "mongodb://" + r.Host + ":" + r.Port
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

func (i *MongoDB) Collection(name string) *MongoColl {
	return &MongoColl{
		Name:         name,
		DatabaseName: i.Name,
		Collection:   i.Database.Collection(name),
	}
}
