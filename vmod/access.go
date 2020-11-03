package vmod

type (
	Permission map[string][]Access
	Access     struct {
		ModelId string `json:"model_id" bson:"model_id"`
		Created int64  `json:"created" bson:"created"`
	}
)

func InitPermission(role string, model_id string, created int64) *Permission {
	access := new(Access)
	permission := make(Permission)
	access.ModelId = model_id
	access.Created = created
	permission[role] = append(permission[role], *access)
	return &permission
}

func (p *Permission) Add(role string, model_id string, created int64) *Permission {
	s := (*p)[role]
	for _, v := range s {
		if v.ModelId == model_id {
			return p
		}
	}
	access := new(Access)
	access.ModelId = model_id
	access.Created = created
	(*p)[role] = append((*p)[role], *access)
	return p
}

func (p *Permission) Delete(role string, model_id string) *Permission {
	s := (*p)[role]
	for i, v := range s {
		if v.ModelId == model_id {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	(*p)[role] = s
	if s == nil {
		delete((*p), role)
	}
	return p
}
