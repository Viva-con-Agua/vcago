package vmdb

import (
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Match represents an mongo $match aggregation. Can be used with Pipeline.Match(Match).
type Match bson.D

//NewMatch return an empty Match type.
func NewMatch() *Match {
	return &Match{}
}

//EqualString match if the value is equal to the value of the key in a database collection.
func (i *Match) EqualString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: value})
	}
}

func (i *Match) EqualStringList(key string, value []string) {
	if value != nil {
		filter := bson.A{}
		for n, _ := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: "$or", Value: filter})
	}
}

//EqualBool the value is a string representation of an bool.
//match if value is equal to the value of the key in a database entry.
func (i *Match) EqualBool(key string, value string) {
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
func (i *Match) EqualInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//EqualInt the value is an string representation of an int.
//match if the value is equal to the value of the given key in an database entry.
func (i *Match) EqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

//ElemMatch TODO
func (i *Match) ElemMatch(list string, key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: list, Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: key, Value: value}}}}})
	}
}

//ElemMatchList TODO
func (i *Match) ElemMatchList(list string, key string, value []string) {
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

func (i *Match) LikeString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}},
		}})
	}
}

func (i *Match) GteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *Match) LteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

func (i *Match) GteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *Match) LteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

//ExpIn TODO
func (i *Match) ExpIn(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$in", Value: bson.A{value}}}})
	}
}
