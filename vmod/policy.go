package vmod

import "github.com/google/uuid"

type (
	//Policies represents the root struct for policy handling.
	Policies struct {
		ID           string       `bson:"_id" json:"id" validate:"required"`
		UserID       string       `bson:"user_id" json:"user_id" validate:"required"`
		Status       bool         `json:"status" bson:"status"`
		PoliciesData PoliciesData `json:"data" bson:"data"`
	}
	//Policy represents whether the user accepted a policy and when the status was modified
	Policy struct {
		Status   bool  `json:"status" bson:"status"`
		Modified int64 `json:"modified" bson:"modified,omitempty"`
	}
	//PoliciesData map policy name to Policy struct
	PoliciesData map[string]Policy
)

//InitPolicies initial a Policies struct
func InitPolicies(userID string, p string, m int64) *Policies {
	policiesData := make(PoliciesData)
	policy := Policy{Status: false, Modified: m}
	policiesData[p] = policy
	return &Policies{
		ID:           uuid.New().String(),
		UserID:       userID,
		Status:       false,
		PoliciesData: policiesData,
	}
}

//Add Policy to Policies
func (pol *Policies) Add(p string, s bool, m int64) *Policies {
	policy := Policy{Status: s, Modified: m}
	pol.PoliciesData[p] = policy
	return pol
}

//Update the state of Policy in Policies
func (pol *Policies) Update(p string, m int64) *Policies {
	poll := pol.PoliciesData[p]
	poll.Status = !poll.Status
	return pol
}
