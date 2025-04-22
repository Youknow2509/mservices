package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"example.com/send_mail/global"
	"example.com/send_mail/model"
	"example.com/send_mail/service"
	isSendMail "example.com/send_mail/service/impl"
	"example.com/send_mail/utils/smtp"
	"example.com/send_mail/utils/smtp/impl"
)

// reader and process
func ReaderAndProcess() {
	r := global.ReaderOTPAuth
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
			break
		}

		key := strings.TrimSpace(string(m.Key))
		var value model.Message
		if err := json.Unmarshal(m.Value, &value); err != nil {
			log.Fatal("failed to unmarshal message:", err)
			continue
		}
		// Process message here
		fmt.Printf("Consumed message: %s, Key: %s, Value: %+v\n", string(m.Topic), key, value)

		switch key {
		case "smtp":
			var smtpService smtp.ISmtpSendMailService

			smtpService = smtp.NewISmtpSendMailService(impl.NewSmtpSendMailService(smtp.AuthSmtp{
				Identity: "",
				Username: global.Config.Smtp.Username,
				Password: global.Config.Smtp.Password,
				Host:     global.Config.Smtp.Host,
				Port:     global.Config.Smtp.Port,
			}))

			err := smtpService.SendHtml(value.To, "TEST", value.Data, "otp-auth-2.html")
			if err != nil {
				fmt.Println("Error sending smtp email:", err)
			} else {
				fmt.Println("Email sent with smtp successfully")
			}
		case "sendgrid":
			implServiceSendMail := isSendMail.NewSendMailImpl()
			service.NewSendMailService(implServiceSendMail)
			sSendMail := service.GetSendMailService()
			err = sSendMail.SendMail(value)
			if err != nil {
				log.Fatal("failed to send mail:", err)
				continue
			}
		default:
			log.Println("Unknown key:", key)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
