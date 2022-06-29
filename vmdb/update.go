package vmdb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//NewUpdateSet set all values in the database.
func NewUpdateSet(value interface{}) bson.D {
	return bson.D{
		{Key: "$set", Value: value},
		{Key: "$set", Value: bson.D{{Key: "modified.updated", Value: time.Now().Unix()}}},
	}
}
