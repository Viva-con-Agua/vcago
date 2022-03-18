package vcago

import (
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoMatch bson.D

//NewMongoMatch creates an new MongoMatch struct.
func NewMongoMatch() *MongoMatch {
	return &MongoMatch{}
}

//AddEqualString
//match if value is equal to the value of the key in a database entry.
func (i *MongoMatch) AddEqualString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: value})
	}
}

//AddEqualBool the value is a string representation of an bool.
//match if value is equal to the value of the key in a database entry.
func (i *MongoMatch) AddEqualBool(key string, value string) {
	if value != "" {
		if value == "false" {
			*i = append(*i, bson.E{Key: key, Value: false})
		}
		if value == "true" {
			*i = append(*i, bson.E{Key: key, Value: true})
		}
	}
}

//AddEqualInt the value is an string representation of an int64.
//match if value is equal to the value of the key in a database entry.
func (i *MongoMatch) AddEqualInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//AddEqualInt the value is an string representation of an int.
//match if the value is equal to the value of the given key in an database entry.
func (i *MongoMatch) AddEqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//AddElemMatch TODO
func (i *MongoMatch) AddElemMatch(list string, key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: list, Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: key, Value: value}}}}})
	}
}

//AddElemMatchList TODO
func (i *MongoMatch) AddElemMatchList(list string, key string, value []string) {
	if value != nil {
		filter := bson.A{}
		for n, _ := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: list, Value: bson.D{
			{Key: "$elemMatch", Value: bson.D{
				{Key: "$or", Value: filter},
			}},
		}})
	}
}

func (i *MongoMatch) AddLikeString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}},
		}})
	}
}

func (i *MongoMatch) AddGteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) AddLteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) AddGteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) AddLteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) AddStringList(key string, value []string) {
	if value != nil {
		filter := bson.A{}
		for n, _ := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: "$or", Value: filter})
	}
}

//AddExpIn TODO
func (i *MongoMatch) AddExpIn(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$in", Value: bson.A{value}}}})
	}
}
