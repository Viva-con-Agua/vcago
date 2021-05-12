package vmod

type SimpleID struct {
	ID    string `bson:"id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Model string `bson:"model" json:"model"`
}
