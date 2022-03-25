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

//EqualString
//match if value is equal to the value of the key in a database entry.
func (i *MongoMatch) EqualString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: value})
	}
}

//EqualBool the value is a string representation of an bool.
//match if value is equal to the value of the key in a database entry.
func (i *MongoMatch) EqualBool(key string, value string) {
	if value != "" {
		if value == "false" {
			*i = append(*i, bson.E{Key: key, Value: false})
		}
		if value == "true" {
			*i = append(*i, bson.E{Key: key, Value: true})
		}
	}
}

//EqualInt the value is an string representation of an int64.
//match if value is equal to the value of the key in a database entry.
func (i *MongoMatch) EqualInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//EqualInt the value is an string representation of an int.
//match if the value is equal to the value of the given key in an database entry.
func (i *MongoMatch) EqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//ElemMatch TODO
func (i *MongoMatch) ElemMatch(list string, key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: list, Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: key, Value: value}}}}})
	}
}

//ElemMatchList TODO
func (i *MongoMatch) ElemMatchList(list string, key string, value []string) {
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

func (i *MongoMatch) LikeString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}},
		}})
	}
}

func (i *MongoMatch) GteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) LteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) GteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) LteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

func (i *MongoMatch) StringList(key string, value []string) {
	if value != nil {
		filter := bson.A{}
		for n, _ := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: "$or", Value: filter})
	}
}

//ExpIn TODO
func (i *MongoMatch) ExpIn(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$in", Value: bson.A{value}}}})
	}
}
