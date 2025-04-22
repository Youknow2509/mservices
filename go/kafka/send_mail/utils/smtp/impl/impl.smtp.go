package impl

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	ssmtp "example.com/send_mail/utils/smtp"
)

type SmtpSendMailService struct {
	auth ssmtp.AuthSmtp
}

// Auth implements smtp.ISmtpSendMailService.
func (s SmtpSendMailService) Auth(authSmtp ssmtp.AuthSmtp) smtp.Auth {
	// Set up authentication information.
	authServer := smtp.PlainAuth(authSmtp.Identity, authSmtp.Username, authSmtp.Password, authSmtp.Host)
	// return nil, authServer
	return authServer
}

// SendHtml implements smtp.ISmtpSendMailService.
func (s SmtpSendMailService) SendHtml(to string, subject string, body string, htmlPath string) error {
	// Create authentication
	authServer := s.Auth(s.auth)

	// Get HTML template if htmlPath is provided
	var htmlBody string
	var err error

	if htmlPath != "" {
		// Create data map for template
		data := map[string]interface{}{
			"otpCode":  body,
			"userName": to,
		}

		htmlBody, err = getMailTemplate(htmlPath, data)
		if err != nil {
			return err
		}
	} else {
		htmlBody = body
	}

	// Format email with HTML content-type
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, htmlBody))

	// Send the email
	recipients := []string{to}
	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", s.auth.Host, s.auth.Port),
		authServer,
		s.auth.Username,
		recipients,
		msg,
	)

	return err
}

// SendText implements smtp.ISmtpSendMailService.
func (s SmtpSendMailService) SendText(to string, subject string, body string) error {
	// Create authentication
	authServer := s.Auth(s.auth)

	// Format plain text email
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body))

	// Send the email
	recipients := []string{to}
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", s.auth.Host, s.auth.Port),
		authServer,
		s.auth.Username,
		recipients,
		msg,
	)

	return err
}

// new ISmtpSendMail implement ISmtpSendMailService
func NewSmtpSendMailService(auth ssmtp.AuthSmtp) ssmtp.ISmtpSendMailService {
	return &SmtpSendMailService{
		auth: auth,
	}
}

// Help get template email html
func getMailTemplate(
	nameMailTemplate string,
	dataTemplate map[string]interface{},
) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(
		template.New(nameMailTemplate).ParseFiles("html-template/mail/" + nameMailTemplate))

	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}
