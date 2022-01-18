package vauth

import (
	"strings"

	"github.com/Viva-con-Agua/vcago/vmdb"
)

type (
	User struct {
		ID      string `bson:"_id" json:"id"`
		Profile string `bson:"profile" json:"profile"`
		Roles   string `bson:"roles" json:"roles"`
	}
	Profile struct {
		ID          string        `bson:"_id" json:"id"`
		Email       string        `bson:"email" json:"email"`
		FirstName   string        `bson:"first_name" json:"first_name" validate:"required"`
		LastName    string        `bson:"last_name" json:"last_name" validate:"required"`
		FullName    string        `bson:"full_name" json:"full_name"`
		DisplayName string        `bson:"display_name" json:"display_name"`
		Gender      string        `bson:"gender" json:"gender"`
		Country     string        `bson:"country" json:"country"`
		Avatar      Avatar        `bson:"avatar" json:"avatar"`
		UserID      string        `bson:"UserID" json:"UserID"`
		Modified    vmdb.Modified `bson:"modified" json:"modified"`
	}
	//Avatar represents the avatar for an User
	Avatar struct {
		URL  string `bson:"url" json:"url"`
		Type string `bson:"type" json:"type"`
	}
)

func (i *User) SetRole(role string) {
	r := strings.Split(i.Roles, ";")
	r = r[:len(r)-1]
	if !strings.Contains(i.Roles, role) {
		i.Roles = i.Roles + role + ";"
	}
}

func (i *User) DeleteRole(role string) {
	if strings.Contains(i.Roles, role) {
		r := strings.Split(i.Roles, ";")
		r = r[:len(r)-1]
		newRoles := ""
		for index := range r {
			if r[index] != role {
				newRoles = newRoles + r[index] + ";"
			}
		}
		i.Roles = newRoles
	}
}
