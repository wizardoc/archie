package services

import (
	"archie/utils"
	"archie/utils/configer"
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"os"
	"regexp"
)

const (
	SMTP_ADDR = "smtp.126.com"
	SMTP_PORT = 25
)

type TargetEmailInfo struct {
	Addr        string
	Subject     string
	Body        string
	ContentType string
}

func SendEmail(emailInfo TargetEmailInfo) error {
	mail := gomail.NewMessage()
	mailConfig := configer.LoadEmailConfig()

	mail.SetAddressHeader("From", mailConfig.Username, "Wizard Team")
	mail.SetHeader("To", mail.FormatAddress(emailInfo.Addr, "用户"))
	mail.SetHeader("Subject", emailInfo.Subject)
	mail.SetBody(emailInfo.ContentType, emailInfo.Body)

	dialer := gomail.NewDialer(SMTP_ADDR, SMTP_PORT, mailConfig.Username, mailConfig.Key)

	return dialer.DialAndSend(mail)
}

func SendContent(addr string) {

}

func SendVerifyCode(addr string) {
	file, err := os.Open("templates/email-valid.html")

	content, err := ioutil.ReadAll(file)
	defer file.Close()

	verifyCode := utils.CreateVerifyCode()

	r, _ := regexp.Compile("{{ verifyCode }}")

	err = SendEmail(TargetEmailInfo{
		Addr:        addr,
		Subject:     "请完成邮箱认证",
		Body:        r.ReplaceAllString(string(content), verifyCode),
		ContentType: "text/html",
	})

	/** 重新尝试发送 */
	if err != nil {
		SendVerifyCode(addr)
	}
}
