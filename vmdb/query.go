package vmdb

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Query represents the default query set for database requests.
type Query struct {
	Search        string `query:"search"`
	SortField     string `query:"sort"`
	SortDirection string `query:"sort_dir"`
	Limit         int64  `query:"limit"`
	Skip          int64  `query:"skip"`
}

// FindOptions return mongodb find options they can be used with Collection.Find() or Collection.FindAndCount()
func (i Query) FindOptions() *options.FindOptions {
	sort := NewSort()
	sort.Add(i.SortField, i.SortDirection)
	return options.Find().SetSort(sort.Bson()).SetLimit(i.Limit).SetSkip(i.Skip)
}
