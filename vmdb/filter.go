package vmdb

import (
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Filter represents an mongo $match aggregation. Can be used with Pipeline.Filter(Filter).
type Filter bson.D

//NewMatch return an empty Match type.
func NewFilter() *Filter {
	return &Filter{}
}

func (i *Filter) Bson() bson.D {
	return bson.D(*i)
}

//EqualString match if the value is equal to the value of the key in a database collection.
func (i *Filter) EqualString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: value})
	}
}

func (i *Filter) EqualStringList(key string, value []string) {
	if value != nil {
		filter := bson.A{}
		for n := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: "$or", Value: filter})
	}
}

//EqualBool the value is a string representation of an bool.
//match if value is equal to the value of the key in a database entry.
func (i *Filter) EqualBool(key string, value string) {
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
func (i *Filter) EqualInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//EqualInt the value is an string representation of an int.
//match if the value is equal to the value of the given key in an database entry.
func (i *Filter) EqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//ElemMatch TODO
func (i *Filter) ElemMatch(list string, key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: list, Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: key, Value: value}}}}})
	}
}

//ElemMatchList TODO
func (i *Filter) ElemMatchList(list string, key string, value []string) {
	if value != nil {
		filter := bson.A{}
		for n := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: list, Value: bson.D{
			{Key: "$elemMatch", Value: bson.D{
				{Key: "$or", Value: filter},
			}},
		}})
	}
}

func (i *Filter) LikeString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}},
		}})
	}
}

func (i *Filter) GteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *Filter) LteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

func (i *Filter) GteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *Filter) LteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

//ExpIn TODO
func (i *Filter) ExpIn(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$in", Value: bson.A{value}}}})
	}
}
