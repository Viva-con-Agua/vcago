package vcago

import (
	"strings"

	"github.com/google/uuid"
)

type Role struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Label  string `json:"label" bson:"label"`
	Root   string `json:"root" bson:"root"`
	UserID string `json:"user_id" bson:"user_id"`
}

type RoleListCookie []string

type RoleList []Role

func (i *RoleList) Cookie() (r *RoleListCookie) {
	r = new(RoleListCookie)
	for n, _ := range *i {
		*r = append(*r, (*i)[n].Name)
	}
	return
}

func (i *RoleListCookie) CheckRoot(role *Role) bool {
	for n, _ := range *i {
		if strings.Contains(role.Root, (*i)[n]) {
			return true
		}
	}
	return false
}

func (i *RoleListCookie) Validate(roles string) bool {
	for n, _ := range *i {
		if strings.Contains(roles, (*i)[n]) {
			return true
		}
	}
	return false
}

func (i *RoleList) In(roles string) bool {
	for n, _ := range *i {
		if (*i)[n].Name == roles {
			return true
		}
	}
	return false
}

/*
func (i *RoleList) Validate(roles string) bool {
	for n, _ := range *i {
		if strings.Contains(roles, (*i)[n].Name) {
			return true
		}
	}
	return false
}

func (i *RoleList) CheckRoot(role *Role) bool {
	for n, _ := range *i {
		if strings.Contains(role.Root, (*i)[n].Name) {
			return true
		}
	}
	return false
}
*/
func (i *RoleList) Append(role *Role) {
	if !i.In(role.Name) {
		*i = append(*i, *role)
	}
}

func RoleMember(userID string) *Role {
	return &Role{
		ID:     uuid.NewString(),
		Name:   "member",
		Label:  "Member",
		Root:   "system",
		UserID: userID,
	}
}

func RoleAdmin(userID string) *Role {
	return &Role{
		ID:     uuid.NewString(),
		Name:   "admin",
		Label:  "Admin",
		Root:   "system;admin",
		UserID: userID,
	}
}

func RoleEmployee(userID string) *Role {
	return &Role{
		ID:     uuid.NewString(),
		Name:   "employee",
		Label:  "Employee",
		Root:   "system;admin;employee",
		UserID: userID,
	}
}
