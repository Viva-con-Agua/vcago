package vmod

type (
	User struct {
		ID         string     `json:"user_id" bson:"_id" validation:"required,uuid"`
		Email      string     `json:"email" bson:"email" validation:"required,email"`
		Policies   Policies   `json:"policies" bson:"policies"`
		Permission Permission `json:"permission" bson:"permission" validation:"required"`
		Modified   Modified   `json:"modified" bson:"modified" validation:"required"`
		Profile    Profile    `json:"profile" bson:"-"`
	}
)
