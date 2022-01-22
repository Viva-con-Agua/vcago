package vmod

type SimpleID struct {
	ID    string `bson:"id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Model string `bson:"model" json:"model"`
}

type MID struct {
	ID   string `bson:"id" json:"id"`
	Type string `bson:"type" json:"type"`
}

func NewMID(id string, t string) *MID {
	return &MID{
		ID:   id,
		Type: t,
	}
}
