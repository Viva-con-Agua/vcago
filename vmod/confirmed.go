package vmod

type (
	Policies struct {
		Status       bool         `json:"status" bson:"status"`
		PoliciesData PoliciesData `json:"data" bson:"data"`
	}
	Policy struct {
		Status   bool  `json:"status" bson:"status"`
		Modified int64 `json:"modified" bson:"modified,omitempty"`
	}
	PoliciesData map[string]Policy
)

func InitPolicies(p string, s bool, m int64) *Policies {
	policies := make(PoliciesData)
	policy := Policy{Status: s, Modified: m}
	policies[p] = policy
	return &Policies{Status: s, PoliciesData: policies}
}

func (pol *Policies) Add(p string, s bool, m int64) *Policies {
	policy := Policy{Status: s, Modified: m}
	pol.PoliciesData[p] = policy
	return pol
}

func (pol *Policies) Update(p string, m int64) *Policies {
	poll := pol.PoliciesData[p]
	poll.Status = !poll.Status
	return pol
}
