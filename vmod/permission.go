package vmod

type (
	//Permission represents role to an access
	Permission map[string][]Access

	//Access represents access to an model.
	Access struct {
		ModelID string `json:"model_id" bson:"model_id"`
		Created int64  `json:"created" bson:"created"`
	}
)

//NewPermission map role to []Access(modelID, created) and initial Permission.
func NewPermission(role string, modelID string, created int64) *Permission {
	access := new(Access)
	permission := make(Permission)
	access.ModelID = modelID
	access.Created = created
	permission[role] = append(permission[role], *access)
	return &permission
}

//Add Access(modelID created) to role.
func (p *Permission) Add(role string, modelID string, created int64) *Permission {
	s := (*p)[role]
	for _, v := range s {
		if v.ModelID == modelID {
			return p
		}
	}
	access := new(Access)
	access.ModelID = modelID
	access.Created = created
	(*p)[role] = append((*p)[role], *access)
	return p
}

//Delete Access from role. If []Access is nil the role will remove form Permission.
func (p *Permission) Delete(role string, modelID string) *Permission {
	s := (*p)[role]
	for i, v := range s {
		if v.ModelID == modelID {
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
