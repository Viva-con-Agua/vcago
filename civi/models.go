package civi

import "github.com/Viva-con-Agua/echo-pool/crm"

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
		CrmData crm.CrmData `json:"crm_data"`
	}
	//CrmUserSignUp represents civi crm signup request
	CrmUserSignUp struct {
		CrmData crm.CrmData `json:"crm_data" validate:"required"`
		CrmUser CrmUser     `json:"crm_user" validate:"required"`
		Mail    Mail        `json:"mail" validate:"required"`
		Offset  Offset      `json:"offset" validate:"required"`
	}
	//Offset represents additional information about user
	Offset struct {
		KnownFrom  string `json:"known_from" validate:"required"`
		Newsletter bool   `json:"newsletter"`
	}
	//Mail represents an email to share links
	Mail struct {
		Email string `json:"email" validate:"required"`
		Link  string `json:"link" validate:"required"`
	}
)
