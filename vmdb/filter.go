package vmdb

import (
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filter bson.D

func NewFilter() *Filter {
	return &Filter{}
}

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

func (i *Filter) Append(key string, value interface{}) {
	*i = append(*i, bson.E{Key: key, Value: value})
}

func (i *Filter) EqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}
func (i *Filter) EqualIn64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

func (i *Filter) Like(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{
			Key: key, Value: bson.D{
				{Key: "$regex", Value: primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}},
			},
		})
	}
}

func (i *Filter) GteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

func (i *Filter) GteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
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

func (i *Filter) LteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}
