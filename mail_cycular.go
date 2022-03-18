package vcago

type CycularMail struct {
	Email   string   `json:"email"`
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

func NewCycularMail(email string, emails []string, subject string, message string) *CycularMail {
	return &CycularMail{
		Email:   email,
		Emails:  emails,
		Subject: subject,
		Message: message,
	}
}
