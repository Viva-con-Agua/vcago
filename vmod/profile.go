package vmod

type (
	//Profile represents user Profile
	Profile struct {
		ID          string   `bson:"profile_id" json:"profile_id" validate:"required"`
		UserID      string   `bson:"user_id" json:"user_id" validate:"required"`
		FirstName   string   `json:"first_name" validate:"required"`
		LastName    string   `json:"last_name" validate:"required"`
		FullName    string   `json:"full_name" validate:"required"`
		DisplayName string   `json:"display_name" validate:"required"`
		Gender      string   `json:"gender"`
		Country     string   `bson:"country" json:"country"`
		Avatar      Avatar   `json:"avatar"`
		Modified    Modified `json:"modified" bson:"modified" validation:"required"`
	}
	//Avatar represents the avatar for an User
	Avatar struct {
		URL     string `json:"url"`
		Type    string `json:"type"`
		Updated int64  `json:"updated" validate:"required"`
		Created int64  `json:"created" validate:"required"`
	}
)
