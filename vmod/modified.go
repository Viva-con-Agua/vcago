package vmod

type (
	Modified struct {
		Updated int64 `json:"updated" bson:"updated,omitempty"`
		Created int64 `json:"created" bson:"created,omitempty"`
	}
)

func InitModified(c_time int64) *Modified {
	m := Modified{Updated: c_time, Created: c_time}
	return &m
}
