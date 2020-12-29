package email_sender

import (
	"archie/utils/configer"
	"github.com/go-gomail/gomail"
)

const (
	SMTP_ADDR = "smtp.163.com"
	SMTP_PORT = 25
)

type EmailSender struct {
	Addr        string
	Subject     string
	Body        string
	ContentType string
}

func (sender *EmailSender) SendEmail() error {
	mail := gomail.NewMessage()
	mailConfig := configer.LoadEmailConfig()

	mail.SetAddressHeader("From", mailConfig.Username, "Wizard Team")
	mail.SetHeader("To", mail.FormatAddress(sender.Addr, "用户"))
	mail.SetHeader("Subject", sender.Subject)
	mail.SetBody(sender.ContentType, sender.Body)

	dialer := gomail.NewDialer(SMTP_ADDR, SMTP_PORT, mailConfig.Username, mailConfig.Key)

	return dialer.DialAndSend(mail)
}
