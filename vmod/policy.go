package vmod

type (
	//Policies represents the root struct for policy handling.
	Policies struct {
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
func InitPolicies(p string, s bool, m int64) *Policies {
	policies := make(PoliciesData)
	policy := Policy{Status: s, Modified: m}
	policies[p] = policy
	return &Policies{Status: s, PoliciesData: policies}
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
