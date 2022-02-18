package vcago

import "go.mongodb.org/mongo-driver/bson"

type MongoPipe struct {
	Pipe []bson.D
}

func NewMongoPipe() *MongoPipe {
	return &MongoPipe{
		Pipe: []bson.D{},
	}
}

func (i *MongoPipe) AddModelAt(from string, root string, child string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: root},
			{Key: "foreignField", Value: child},
			{Key: "as", Value: as},
		}}}
	unwind := bson.D{{Key: "$unwind", Value: "$" + as}}
	i.Pipe = append(i.Pipe, lookup)
	i.Pipe = append(i.Pipe, unwind)
}
