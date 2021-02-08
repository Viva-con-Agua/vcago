package vmod

type (
	//MailBase base model for request mail-backend via nats
	MailBase struct {
		To           string      `json:"to" bson:"to"`
		JobCase string `json:"job_case"`
		JobScope string `json:"job_scope"`
		Country string `json:"country"`
	}
	//MailCode used for request mail-backend via nats
	MailCode struct {
		MailBase
		UserName string `json:"user_name"`
		Code string `json:"code"`
	}
)

//NewMailCode initial new MailCode from given variables.
func NewMailCode(to string, jobCase string, jobScope string, userName string, code string) MailCode{
	return MailCode{MailBase: MailBase{To: to, JobCase: jobCase, JobScope: jobScope}, UserName: userName, Code: code}
}
