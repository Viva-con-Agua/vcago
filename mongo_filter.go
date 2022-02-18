package vcago

import (
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoFilterM struct {
	Filter bson.M
}

func NewMongoFilterM() *MongoFilterM {
	return &MongoFilterM{
		Filter: bson.M{},
	}
}

func (i *MongoFilterM) Equal(key string, value string) {
	if value != "" {
		i.Filter[key] = value
	}
}

func (i *MongoFilterM) EqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			i.Filter[key] = valueInt
		}
	}
}
func (i *MongoFilterM) EqualIn64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			i.Filter[key] = valueInt
		}
	}
}

func (i *MongoFilterM) Like(key string, value string) {
	if value != "" {
		i.Filter[key] = bson.M{"$regex": primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}}
	}
}

func (i *MongoFilterM) GteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			i.Filter[key] = bson.M{"$gte": valueInt}
		}
	}
}

func (i *MongoFilterM) GteInt64(key string, value string) {
	if value != "" {
		if valueInt64, err := strconv.ParseInt(value, 10, 64); err == nil {
			i.Filter[key] = bson.M{"$gte": valueInt64}
		}
	}
}
func (i *MongoFilterM) LteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			i.Filter[key] = bson.M{"$lte": valueInt}
		}
	}
}

func (i *MongoFilterM) LteInt64(key string, value string) {
	if value != "" {
		if valueInt64, err := strconv.ParseInt(value, 10, 64); err == nil {
			i.Filter[key] = bson.M{"$lte": valueInt64}
		}
	}
}
