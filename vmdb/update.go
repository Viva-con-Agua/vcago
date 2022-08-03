package vmdb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//UpdateSet set all values in the database.
//
//MongoDB:
//	{
//		"$set": value,
//		"$set": {
// 			"modified.updated": time.Now().Unix(),
//		}
//	}
func UpdateSet(value interface{}) bson.D {
	return bson.D{
		{Key: "$set", Value: value},
		{Key: "$set", Value: bson.D{{Key: "modified.updated", Value: time.Now().Unix()}}},
	}
}
