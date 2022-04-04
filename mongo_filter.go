package vcago

import (
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filter struct {
	Filter bson.M
}

func NewFilter() *Filter {
	return &Filter{
		Filter: bson.M{},
	}
}

func (i *Filter) In(key string, list []string) *Filter {
	if list == nil {
		list = []string{""}
	}
	(*i).Filter[key] = bson.M{"$in": list}
	return i
}

func (i *Filter) Equal(key string, value string) *Filter {
	(*i).Filter[key] = value
	return i
}

func (i *Filter) Bson() bson.M {
	return i.Filter
}

type MongoFilter struct {
	Filter bson.M
}

func CreateUpdateManyFilter(key string, values []string) (r *bson.A) {
	r = new(bson.A)
	for n, _ := range values {
		*r = append(*r, bson.D{{Key: key, Value: values[n]}})
	}
	return
}

func NewMongoFilter() *MongoFilter {
	return &MongoFilter{
		Filter: bson.M{},
	}
}

func (i *MongoFilter) Equal(key string, value string) *MongoFilter {
	if value != "" {
		i.Filter[key] = value
	}
	return i
}

func (i *MongoFilter) EqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			i.Filter[key] = valueInt
		}
	}
}
func (i *MongoFilter) EqualIn64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			i.Filter[key] = valueInt
		}
	}
}

func (i *MongoFilter) Like(key string, value string) {
	if value != "" {
		i.Filter[key] = bson.M{"$regex": primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}}
	}
}

func (i *MongoFilter) GteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			i.Filter[key] = bson.M{"$gte": valueInt}
		}
	}
}

func (i *MongoFilter) GteInt64(key string, value string) {
	if value != "" {
		if valueInt64, err := strconv.ParseInt(value, 10, 64); err == nil {
			i.Filter[key] = bson.M{"$gte": valueInt64}
		}
	}
}
func (i *MongoFilter) LteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			i.Filter[key] = bson.M{"$lte": valueInt}
		}
	}
}

func (i *MongoFilter) LteInt64(key string, value string) {
	if value != "" {
		if valueInt64, err := strconv.ParseInt(value, 10, 64); err == nil {
			i.Filter[key] = bson.M{"$lte": valueInt64}
		}
	}
}
