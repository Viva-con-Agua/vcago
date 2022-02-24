package vcago

import "strings"

type Role struct {
	Name  string `json:"name" bson:"name"`
	Label string `json:"label" bson:"label"`
	Root  string `json:"root" bson:"root"`
}

type RoleList []Role

func (i *RoleList) In(roles string) bool {
	for n, _ := range *i {
		if (*i)[n].Name == roles {
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

func (i *RoleList) Append(role *Role) {
	if !i.In(role.Name) {
		*i = append(*i, *role)
	}
}

func RoleMember() *Role {
	return &Role{
		Name:  "member",
		Label: "Member",
		Root:  "system",
	}
}

func RoleAdmin() *Role {
	return &Role{
		Name:  "admin",
		Label: "Admin",
		Root:  "system;admin",
	}
}

func RoleEmployee() *Role {
	return &Role{
		Name:  "employee",
		Label: "Employee",
		Root:  "system;admin;employee",
	}
}
