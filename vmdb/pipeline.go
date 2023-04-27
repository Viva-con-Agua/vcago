// Package vmdb
package vmdb

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Pipeline represents an helper for handling mongodb pipeline. The Pipe param contains an []bson.D that represents an mongo pipeline.
type Pipeline struct {
	Pipe []bson.D
}

// NewPipeline creates an new Pipeline struct.
func NewPipeline() *Pipeline {
	return &Pipeline{
		Pipe: []bson.D{},
	}
}

// Match adds the filter param as $match to the end of the Pipeline struct.
//
// MongoDB:
//
//	{"$match": filter}
func (i *Pipeline) Match(filter bson.D) *Pipeline {
	if filter != nil {
		match := bson.D{{
			Key:   "$match",
			Value: filter,
		}}
		i.Pipe = append(i.Pipe, match)
	}
	return i
}

// LockupUnwind represents the lookup and unwind combination to join an element from a second collection to the result.
//
// MongoDB:
//
//	{
//		"$lookup":{
//			"from": from,
//			"localField": localField,
//			"foreignField" foreignField,
//			"as": as
//		},
//		"$unwind": {
//			"path": "$as",
//			"preserveNullAndEmptyArrays": true
//		}
//	}
func (i *Pipeline) LookupUnwind(from string, localField string, foreignField string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: localField},
			{Key: "foreignField", Value: foreignField},
			{Key: "as", Value: as},
		}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$" + as}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	i.Pipe = append(i.Pipe, lookup)
	i.Pipe = append(i.Pipe, unwind)
}

// Skip used for the mongo function $skip
func (i *Pipeline) Skip(value int64, defaultValue int64) *Pipeline {
	var skip = defaultValue
	if value > 0 {
		skip = value
	}
	skipPipe := bson.D{{Key: "$skip", Value: skip}}
	i.Pipe = append(i.Pipe, skipPipe)
	return i

}

// Sort is used for the mongo function $sort. Use Sort object for input.
func (i *Pipeline) Sort(value bson.D) *Pipeline {
	if len(value) > 0 {
		sort := bson.D{{Key: "$sort", Value: value}}
		i.Pipe = append(i.Pipe, sort)
	}
	return i
}

// Limit is used for the mongo function $limit
func (i *Pipeline) Limit(value int64, defaultValue int64) *Pipeline {
	var limit = defaultValue
	if value > 0 {
		limit = value
	}
	limitPipe := bson.D{{Key: "$limit", Value: limit}}
	i.Pipe = append(i.Pipe, limitPipe)
	return i

}

// Lookup represents an lookup to join an list of elements from a second collection to the result.
//
// MongoDB:
//
//	{
//		"$lookup":{
//			"from": from,
//			"localField": localField,
//			"foreignField" foreignField,
//			"as": as
//		}
//	}
func (i *Pipeline) Lookup(from string, localField string, foreignField string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: localField},
			{Key: "foreignField", Value: foreignField},
			{Key: "as", Value: as},
		}}}
	i.Pipe = append(i.Pipe, lookup)
}

// LookupUnwindMatch represents the lookup and unwind combination to join an element from a second collection to the result.
// The joined element can be filtered by the match param.
//
// MongoDB:
//
//	{
//		"$lookup":{
//			"from": from,
//			"localField": localField,
//			"foreignField" foreignField,
//			"pipeline": [{"$match": match}]
//			"as": as
//		},
//		"$unwind": {
//			"path": "$as",
//			"preserveNullAndEmptyArrays": true
//		}
//	}
func (i *Pipeline) LookupUnwindMatch(from string, localField string, foreignField string, as string, match bson.D) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: localField},
			{Key: "foreignField", Value: foreignField},
			{Key: "pipeline", Value: append([]bson.D{}, bson.D{{Key: "$match", Value: match}})},
			{Key: "as", Value: as},
		}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$" + as}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	i.Pipe = append(i.Pipe, lookup)
	i.Pipe = append(i.Pipe, unwind)
}

// LookupMatch represents an lookup to join an list of elements from a second collection to the result.
// The joined elements can be filtered by the match param.
//
// MongoDB:
//
//	{
//		"$lookup":{
//			"from": from,
//			"localField": localField,
//			"foreignField" foreignField,
//			"pipeline": [{"$match": match}]
//			"as": as
//		}
//	}
func (i *Pipeline) LookupMatch(from string, localField string, foreignField string, as string, match bson.D) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: localField},
			{Key: "foreignField", Value: foreignField},
			{Key: "pipeline", Value: append([]bson.D{}, bson.D{{Key: "$match", Value: match}})},
			{Key: "as", Value: as},
		}}}
	i.Pipe = append(i.Pipe, lookup)
}

// LookupList represents an lookup to join an list of elements from a second collection to the result.
// The value of the localField need to be a list of references. If the foreignField value is in the list, the element will joined to the as value.
//
// MongoDB:
//
//	{
//		"$lookup":{
//			"from": from,
//			"let": { localField: $localField },
//			"foreignField" foreignField,
//			"pipeline": [{
//				"$match":{
//					"$expr":{"$in: ["$foreinField", "$$localField"]}
//				}
//			}],
//			"as": as
//		}
//	}
func (i *Pipeline) LookupList(from string, localField string, foreignField string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "let", Value: bson.D{{Key: localField, Value: "$" + localField}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{
					{Key: "$match", Value: bson.D{{
						Key: "$expr", Value: bson.D{{Key: "$in", Value: bson.A{"$" + foreignField, "$$" + localField}}},
					}}},
				},
			}},
			{Key: "as", Value: as},
		},
	}}
	i.Pipe = append(i.Pipe, lookup)
}

// Append appends the elements in pipe to the the Pipeline object.
func (i *Pipeline) AppendSlice(pipe []bson.D) {
	for n := range pipe {
		i.Pipe = append(i.Pipe, pipe[n])
	}
}

func (i *Pipeline) Append(entry bson.D) {
	i.Pipe = append(i.Pipe, entry)
}
