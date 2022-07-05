package vmdb

import "go.mongodb.org/mongo-driver/bson"

type Pipeline struct {
	Pipe []bson.D
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Pipe: []bson.D{},
	}
}

func (i *Pipeline) Match(m *Match) *Pipeline {
	if *m != nil {
		match := bson.D{{
			Key:   "$match",
			Value: *m,
		}}
		i.Pipe = append(i.Pipe, match)
	}
	return i
}

func (i *Pipeline) LookupUnwind(from string, root string, child string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: root},
			{Key: "foreignField", Value: child},
			{Key: "as", Value: as},
		}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$" + as}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	i.Pipe = append(i.Pipe, lookup)
	i.Pipe = append(i.Pipe, unwind)
}

func (i *Pipeline) Lookup(from string, root string, child string, as string) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: root},
			{Key: "foreignField", Value: child},
			{Key: "as", Value: as},
		}}}
	i.Pipe = append(i.Pipe, lookup)
}

func (i *Pipeline) LookupUnwindMatch(from string, root string, child string, as string, match bson.D) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: root},
			{Key: "foreignField", Value: child},
			{Key: "pipeline", Value: []bson.D{match}},
			{Key: "as", Value: as},
		}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$" + as}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
	i.Pipe = append(i.Pipe, lookup)
	i.Pipe = append(i.Pipe, unwind)
}

func (i *Pipeline) LookupMatch(from string, root string, child string, as string, match bson.D) {
	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: from},
			{Key: "localField", Value: root},
			{Key: "foreignField", Value: child},
			{Key: "pipeline", Value: []bson.D{match}},
			{Key: "as", Value: as},
		}}}
	i.Pipe = append(i.Pipe, lookup)
}
