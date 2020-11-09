package vmod

type (
	//Mail represents the email backend model
	Mail struct {
		From         string      `json:"from" bson:"from"`
		To           string      `json:"to" bson:"to"`
		TemplateData interface{} `json:"template_data" bson:"template_data"`
		TemplateID   string      `json:"template_id" bson:"template_id"`
	}

	//MailTemplate represents a template model using for database storage.
	MailTemplate struct {
		ID       string `bson:"_id" json:"id" validate:"required"`
		Name     string `bson:"name" json:"name" validate:"required"`
		Template string `bson:"template" json:"template" `
	}
)
