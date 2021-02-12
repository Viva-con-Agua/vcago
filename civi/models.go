package civi


type (
	//CrmUser represents civi crm user
	CrmUser struct {
		Email         string `json:"email" validate:"required"`
		FirstName     string `json:"first_name" validate:"required"`
		LastName      string `json:"last_name" validate:"required"`
		PrivacyPolicy bool   `json:"privacy_policy"`
		Country       string `json:"country"`
	}
	//CrmDataBody represents civi crm data
	CrmDataBody struct {
		CrmData CrmData `json:"crm_data"`
	}
	//CrmUserSignUp represents civi crm signup request
	CrmUserSignUp struct {
		CrmData CrmData `json:"crm_data" validate:"required"`
		CrmUser CrmUser     `json:"crm_user" validate:"required"`
		Mail    CrmEmail        `json:"mail" validate:"required"`
		Offset  Offset      `json:"offset" validate:"required"`
	}
	//Offset represents additional information about user
	Offset struct {
		KnownFrom  string `json:"known_from" validate:"required"`
		Newsletter bool   `json:"newsletter"`
	}
	CrmData struct {
		CampaignID int    `json:"campaign_id" validate:"required"`
		DropsID    string `json:"drops_id" validate:"required"`
		Activity   string `json:"activity" validate:"required"`
		Country    string `json:"country,omitempty"`
		Created    int64  `json:"created"`
	}
	CrmEmail struct {
		Email string `json:"email"`
		Link  string `json:"link"`
	}
	CrmEmailBody struct {
		CrmData CrmData  `json:"crm_data"`
		Mail    CrmEmail `json:"mail"`
	}

)
