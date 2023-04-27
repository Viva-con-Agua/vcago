package vmdb

import "go.mongodb.org/mongo-driver/bson"

// Sort represents an type for manage sorting in mongo db.
type Sort struct {
	Value bson.D
}

// NewSort creates a new Sort object.
func NewSort() *Sort {
	return &Sort{Value: bson.D{}}
}

// Add adds an sort option to the sort object. The function convert ASC to 1 and DESC to -1.
func (i *Sort) Add(key string, value string) {
	if value != "" {
		if value == "ASC" {
			i.Value = append(i.Value, bson.E{Key: key, Value: 1})
		}
		if value == "DESC" {
			i.Value = append(i.Value, bson.E{Key: key, Value: -1})
		}
	}
}

// Bson returns the bson object.
func (i *Sort) Bson() bson.D {
	return i.Value
}
