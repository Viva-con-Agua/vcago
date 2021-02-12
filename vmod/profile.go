package vmod

type (
	//Profile represents user Profile
	Profile struct {
		ID          string   `bson:"_id" json:"profile_id" validate:"required"`
		UserID      string   `bson:"user_id" json:"user_id" validate:"required"`
		FirstName   string   `bson:"first_name" json:"first_name" validate:"required"`
		LastName    string   `bson:"last_name" json:"last_name" validate:"required"`
		FullName    string   `bson:"full_name" json:"full_name"`
		DisplayName string   `bson:"display_name" json:"display_name"`
		Gender      string   `bson:"gender" json:"gender"`
		Country     string   `bson:"country" json:"country"`
		Avatar      Avatar   `bson:"avatar" json:"avatar"`
		Modified    Modified `json:"modified" bson:"modified" validation:"required"`
	}
	//Avatar represents the avatar for an User
	Avatar struct {
		URL     string `bson:"url" json:"url"`
		Type    string `bson:"type" json:"type"`
		Updated int64  `bson:"updated" json:"updated" validate:"required"`
		Created int64  `bson:"created" json:"created" validate:"required"`
	}
)
