package vmod

import (
	"strings"

	"github.com/google/uuid"
)

// Role represents an struct for handling access roles.
type Role struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Label  string `json:"label" bson:"label"`
	Root   string `json:"root" bson:"root"`
	UserID string `json:"user_id" bson:"user_id"`
}
type WebAppRole struct {
	Name  string    `json:"name" validate:"required"`
	Roles []AppRole `json:"roles" validate:"required"`
}
type AppRole struct {
	Label string `json:"label" validate:"required"`
	Role  string `json:"role" validate:"required"`
}

// RoleListCookie used for the access_token.
type RoleListCookie []string

// RoleList represents an slice of Role
type RoleList []Role

// Cookie return an slice of strings they contains all role names.
func (i *RoleList) Cookie() (r *RoleListCookie) {
	r = new(RoleListCookie)
	for n := range *i {
		*r = append(*r, (*i)[n].Name)
	}
	return
}

// CheckRoot check if the RoleListCookie object contains an role in the Root param.
// So you can check if an user is allow to give an other user an role.
func (i *RoleListCookie) CheckRoot(role *Role) bool {
	for n := range *i {
		if strings.Contains(role.Root, (*i)[n]) {
			return true
		}
	}
	return false
}

// Validate check if an RoleListCookie contains an role in the roles string. We seperate the role by ;.
//
// Example:
//
//	"admin;employee"
func (i *RoleListCookie) Validate(roles string) bool {
	for n := range *i {
		validations := strings.Split(roles, ";")
		for _, validation := range validations {
			if validation == (*i)[n] {
				return true
			}
		}
	}
	return false
}

// In check if a the RoleList contains an role.
func (i *RoleList) In(role string) bool {
	for n := range *i {
		if (*i)[n].Name == role {
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

// Append append a Role role to a RoleList i. If i contains the role, nothing happend.
func (i *RoleList) Append(role *Role) {
	if !i.In(role.Name) {
		*i = append(*i, *role)
	}
}

// RoleMember represents the member role.
func RoleMember(userID string) *Role {
	return &Role{
		ID:     uuid.NewString(),
		Name:   "member",
		Label:  "Member",
		Root:   "system",
		UserID: userID,
	}
}

// RoleAdmin represents the admin role.
func RoleAdmin(userID string) *Role {
	return &Role{
		ID:     uuid.NewString(),
		Name:   "admin",
		Label:  "Admin",
		Root:   "system;admin",
		UserID: userID,
	}
}

// RoleEmployee represents the employee role.
func RoleEmployee(userID string) *Role {
	return &Role{
		ID:     uuid.NewString(),
		Name:   "employee",
		Label:  "Employee",
		Root:   "system;admin;employee",
		UserID: userID,
	}
}
