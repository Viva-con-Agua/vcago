package vcago

import "time"

type Modified struct {
	Updated int64 `json:"updated" bson:"updated"`
	Created int64 `json:"created" bson:"created"`
}

type ModifiedQuery struct {
	UpdateFrom string `query:"update_from"`
	UpdateTo   string `query:"update_to"`
	CreateFrom string `query:"create_from"`
	CreateTo   string `query:"create_to"`
}

func (i *ModifiedQuery) Filter(filter *MongoFilterM) {
	filter.LteInt64("modified.created", i.CreateTo)
	filter.GteInt64("modified.created", i.CreateFrom)
	filter.LteInt64("modified.updated", i.UpdateTo)
	filter.GteInt64("modified.updated", i.UpdateFrom)

}

//NewModified initial Modified model, cTime is the current time Unix format.
func NewModified() Modified {
	t := time.Now().Unix()
	m := Modified{Updated: t, Created: t}
	return m
}

func (i *Modified) Update() {
	i.Updated = time.Now().Unix()
}
