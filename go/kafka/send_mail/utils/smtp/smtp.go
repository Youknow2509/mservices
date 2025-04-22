package smtp

import "net/smtp"

// interface smtp send mail service
type ISmtpSendMailService interface {
	Auth(authSmtp AuthSmtp) smtp.Auth
	SendText(to string, subject string, body string) error
	SendHtml(to string, subject string, body string, htmlPath string) error
}

// struct smtp auth
type AuthSmtp struct {
	Identity string
	Username string
	Password string
	Host     string
	Port     string
}

var vISmtpSendMailService ISmtpSendMailService

// new ISmtpSendMail with auth
func NewISmtpSendMailService(v ISmtpSendMailService) ISmtpSendMailService {
	if vISmtpSendMailService == nil {
		vISmtpSendMailService = v
	} 
	return vISmtpSendMailService
}
