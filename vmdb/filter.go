package vmdb

import (
	"regexp"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Filter represents an mongo filter object.
type Filter bson.D

// NewMatch return an empty Match type.
func NewFilter() *Filter {
	return &Filter{}
}

// Bson return the bson.D object.
func (i *Filter) Bson() bson.D {
	return bson.D(*i)
}

// EqualString match if the value is equal to the value of the key in a database collection.
//
// MongoDB:
//
//	{
//		key: value
//	}
func (i *Filter) EqualString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: value})
	}
}

// EqualStringList match if the the value of the key param is matching an element in the value param slice.
//
// MongoDB:
//
//	{
//		key: {"$or": value}
//	}
func (i *Filter) EqualStringList(key string, value []string) {
	if value != nil && len(value) != 0 {
		filter := bson.A{}
		for n := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: "$or", Value: filter})
	}
}

// EqualBool the value is a string representation of an bool.
// match if value is equal to the value of the key in a database entry.
//
// MongoDB:
//
//	{
//		key: value as boolean
//	}
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

// EqualInt the value is an string representation of an int64.
// match if value is equal to the value of the key in a database entry.
//
// MongoDB:
//
//	{
//		key: value as int64
//	}
func (i *Filter) EqualInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

// EqualInt the value is an string representation of an int.
// match if the value is equal to the value of the given key in an database entry.
//
// MongoDB:
//
//	{
//		key: value as int
//	}
func (i *Filter) EqualInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: valueInt})
		}
	}
}

// ElemMatch TODO
func (i *Filter) ElemMatch(list string, key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: list, Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: key, Value: value}}}}})
	}
}

// ElemMatchList TODO
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

// LikeString use regex for handling a substring matching at the start of a string.
//
// MongoDB:
//
//	{
//		key: {"$regex": ^value}
//	}
func (i *Filter) LikeString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: "^" + regexp.QuoteMeta(value)}},
			{Key: "$options", Value: "i"},
		}})
	}
}

// ContainsString use regex for handling a substring matching.
//
// MongoDB:
//
//	{
//		key: {"$regex": .*value.*}
//	}
func (i *Filter) ContainsString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{
			{Key: "$regex", Value: primitive.Regex{Pattern: ".*" + regexp.QuoteMeta(value) + ".*"}},
			{Key: "$options", Value: "i"},
		}})
	}
}

// GteInt64 provides $gte for key they have an int64 datatype.
// If the value element is "" or not an int64 formated string no element will be added to the filter object.
//
// MongoDB:
//
//	{
//		key: {"$gte": value as int64}
//	}
func (i *Filter) GteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

// LteInt64 provides $lte for key they have an int64 datatype.
// If the value element is "" or not an int64 formated string no element will be added to the filter object.
//
// MongoDB:
//
//	{
//		key: {"$lte": value as int64}
//	}
func (i *Filter) LteInt64(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.ParseInt(value, 10, 64); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

// GteInt provides $lte for key they have an int6 datatype.
// If the value element is "" or not an int formated string no element will be added to the filter object.
//
// MongoDB:
//
//	{
//		key: {"$gte": value as int}
//	}
func (i *Filter) GteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$gte", Value: valueInt}}})
		}
	}
}

// LteInt provides $lte for key they have an int datatype.
// If the value element is "" or not an int formated string no element will be added to the filter object.
//
// MongoDB:
//
//	{
//		key: {"$lte": value as int}
//	}
func (i *Filter) LteInt(key string, value string) {
	if value != "" {
		if valueInt, err := strconv.Atoi(value); err == nil {
			*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$lte", Value: valueInt}}})
		}
	}
}

// SearchString searchs for a given string in all the given fields
// If the value search string is "" no search string will be added to the filter object.
//
// MongoDB:
//
//	{
//		key: {"$or": key: {"$regex": ".*value.*"}}
//	}
func (i *Filter) SearchString(fields []string, value string) bson.D {
	if value != "" {
		filter := bson.A{}
		for _, field := range fields {
			filter = append(filter, bson.D{{Key: field, Value: bson.D{
				{Key: "$regex", Value: primitive.Regex{Pattern: ".*" + regexp.QuoteMeta(value) + ".*"}},
				{Key: "$options", Value: "i"},
			}}})
		}
		*i = append(*i, bson.E{Key: "$or", Value: filter})
	}
	return bson.D{}
}

// ExpIn TODO
func (i *Filter) ExpIn(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: bson.D{{Key: "$in", Value: bson.A{value}}}})
	}
}

func (i *Filter) Append(value bson.E) {
	*i = append(*i, value)
}
