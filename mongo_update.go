package vcago

import "go.mongodb.org/mongo-driver/bson"

type MongoUpdate struct {
	Update bson.M
}

//func (i *MongoUpdate) Set(value interface)
