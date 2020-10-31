package api

import "log"

type (
	User struct {
		ID            string       `json:"uuid" bson:"_id,omitempty" validation:"required,uuid"`
		Email         string       `json:"email" bson:"email" validation:"required,email"`
		Confirmed     Confirmed    `json:"confirmed" bson:"confirmed" validation:"required"`
		PrivacyPolicy Confirmed    `json:"privacy_policy" bson:"privacy_policy" validation:"required"`
		Modified      Modified     `json:"modified" bson:"modified" validation:"required"`
		PasswordInfo  PasswordInfo `json:"-" bson:"password"`
		Profile       Profile      `json:"profile" bson:"-"`
	}
	Profile struct {
		ID          string   `bson:"_id" json:"id" validate:"required"`
		UserID      string   `bson:"user_id" json:"user_id" validate:"required"`
		Avatar      Avatar   `json:"avatar"`
		FirstName   string   `json:"first_name" validate:"required"`
		LastName    string   `json:"last_name" validate:"required"`
		FullName    string   `json:"full_name" validate:"required"`
		DisplayName string   `json:"display_name" validate:"required"`
		Gender      string   `json:"gender"`
		Country     string   `bson:"country" json:"country"`
		Modified    Modified `json:"modified" bson:"modified" validation:"required"`
	}
	Avatar struct {
		Url     string `json:"url"`
		Type    string `json:"type"`
		Updated int64  `json:"updated" validate:"required"`
		Created int64  `json:"created" validate:"required"`
	}

	Confirmed struct {
		Status   bool  `json:"status" bson:"status"`
		Modified int64 `json:"modified" bson:"modified,omitempty"`
	}
	Modified struct {
		Updated int64 `json:"updated" bson:"updated,omitempty"`
		Created int64 `json:"created" bson:"created,omitempty"`
	}
	PasswordInfo struct {
		Password []byte   `bson:"password" json:"-"`
		Hasher   string   `bson:"hasher" json:"-"`
		Modified Modified `bson:"modified" json:"-"`
	}
	Access struct {
		As []string `bson:"as" json:"as"`
		On string   `bson:"on" json:"on"`
	}
)

func Dummy() {
	log.Print("dummy")
}
