package vmdb

import (
	"context"
	"log"

	"github.com/Viva-con-Agua/vcago"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database represents the initial struct for an Mongo Database connection named Name. Database is an pointer to an mongo.Database.
type Database struct {
	Name     string //Database Name
	URI      string //Database URI
	Database *mongo.Database
}

func mongoURI() (uri string) {
	uri = "mongodb://"
	user := vcago.Settings.String("MONGO_DB_USER", "w", "")
	password := vcago.Settings.String("MONGO_DB_PASSWORD", "w", "")
	if password != "" && user != "" {
		uri = uri + user + ":" + password + "@"
	}
	host := vcago.Settings.String("MONGO_DB_HOST", "w", "localhost")
	port := vcago.Settings.String("MONGO_DB_PORT", "w", "27017")
	uri = uri + host + ":" + port
	return
}

//NewDatabase creates a new Database object and connect it to mongoDB.
//@PARAM name string  # database name.

func NewDatabase(name string) *Database {
	return &Database{
		Name: name,
		URI:  mongoURI(),
	}
}

//Connect creates an mongo db client and initial an connection.
//
//You can use the following parameters as environment variable or as flag to define the connection parameters.
//	MONGO_DB_HOST=<host>,
// 	MONGO_DB_PORT=<port>,
// 	MONGO_DB_USER=<user>,
// 	MONGO_DB_PASSWORD=<password>
//
//if the username or password is not defined, the Client try to connect without an user.
func (i *Database) Connect() *Database {
	log.Print("MongoDB connection to " + i.URI)
	opts := options.Client()
	opts.ApplyURI(i.URI)
	opts.SetMaxPoolSize(5)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	i.Database = client.Database(i.Name)
	log.Print("MongoDB successfully connected!")
	return i
}

func (i *Database) ConnectTEST() *Database {
	log.Print("Create Memory MongoDB server instance.")
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()
	log.Print("Connect to Memory MongoDB server")
	opts := options.Client()
	opts.ApplyURI(mongoServer.URI())
	opts.SetMaxPoolSize(5)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	i.Database = client.Database(i.Name)
	log.Print("MongoDB successfully connected!")
	return i
}

//Collection initial an new mongodb collection named by the name parameter.
//Use [NewDatabase] for initial an database connection.
func (i *Database) Collection(name string) *Collection {
	return &Collection{
		Name:         name,
		DatabaseName: i.Name,
		Collection:   i.Database.Collection(name),
	}
}
