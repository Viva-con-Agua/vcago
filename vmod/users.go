package vmod

import "github.com/Viva-con-Agua/vcago/civi"

type (
	//User represents the root struct for user handling in viva-con-agua api.
	User struct {
		ID            string     `json:"user_id" bson:"_id" validation:"required,uuid"`
		Email         string     `json:"email" bson:"email" validation:"required,email"`
		PrivacyPolicy string     `json:"policies" bson:"-"`
		Permission    Permission `json:"permission" bson:"permission" validation:"required"`
		Modified      Modified   `json:"modified" bson:"modified" validation:"required"`
		Country       string     `json:"country" bson:"country" validation:"required"`
		Profile       Profile    `json:"profile" bson:"-"`
	}
)

func (u *User) CrmDataBody(civiID int, activity string) *civi.CrmDataBody() {
	return &civi.CrmDataBody{
		CrmData: civi.CrmData{
			CampaignID: civiID,
			DropsID: u.ID,
			Activity: "EVENT_JOIN",
			Country: u.Country
		}
	}
	
}
