package vmod

type (
	Request struct {
		CampaignID string      `json:"campaign_id"`
		Body       interface{} `json:"body"`
	}
)
