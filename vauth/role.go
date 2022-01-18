package vauth

type Roles struct {
	Roles  string `bson:"roles" json:"roles"`
	UserID string `bson:"user_id" json:"user_id"`
}

func NewRoles(id string) *Roles {
	return &Roles{
		Roles:  "member;",
		UserID: id,
	}
}
