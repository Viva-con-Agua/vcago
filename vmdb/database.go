//Database connection for mongo db. The database is accessable by using the DB variable.
// requires DB_HOST, DB_PORT, DB_NAME
package vmdb

import (
	"context"
	"log"

	"github.com/Viva-con-Agua/vcago/vutils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database represents the mongo database connection object.
type Database struct {
	Host string
	Port string
	Name string
	DB   *mongo.Database
}

//DB store the current database connection for use it global in your project.a
var DB = new(Database)

//Connect connects the DB named database controller with an mongodb database
func Connect() {
	DB.Host = vutils.Config.GetEnvString("DB_HOST", "w", "localhost")
	DB.Port = vutils.Config.GetEnvString("DB_PORT", "w", "27017")
	DB.Name = vutils.Config.GetEnvString("DB_NAME", "w", "default")
	uri := "mongodb://" + DB.Host + ":" + DB.Port
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
	DB.DB = client.Database(DB.Name)
	log.Print("MongoDB successfully connected!")
}
