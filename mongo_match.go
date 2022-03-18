package vcago

import "go.mongodb.org/mongo-driver/bson"

type MongoMatch bson.D

func (i *MongoMatch) Init() {
	i = &MongoMatch{}
}

func (i *MongoMatch) AddString(key string, value string) {
	if value != "" {
		*i = append(*i, bson.E{Key: key, Value: value})
	}
}

func (i *MongoMatch) AddStringList(key string, value []string) {
	if value != nil {
		filter := bson.A{}
		for n, _ := range value {
			filter = append(filter, bson.D{{Key: key, Value: value[n]}})
		}
		*i = append(*i, bson.E{Key: "$or", Value: filter})
	}
}
