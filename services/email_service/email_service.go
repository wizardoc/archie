package email_service

import (
	"archie/connection/redis_conn"
	"archie/services/email_sender"
	"archie/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const (
	VALID_TIMING = 60 * 15
)

type EmailService struct {
	Code  string
	Email string
}

// get the verify code from Redis
func (emailService *EmailService) GetCode() {
	redis_conn.GetRedisConnMust(func(conn redis.Conn) error {
		key := genCodeKey(emailService.Email)
		code, err := redis.String(conn.Do("GET", key))
		if err != nil {
			emailService.Code = ""
			return err
		}

		emailService.Code = code

		return nil
	})

	return
}

// delete the key from Redis
func (emailService *EmailService) DelCode() {
	redis_conn.GetRedisConnMust(func(conn redis.Conn) error {
		if _, err := conn.Do("DEL", genCodeKey(emailService.Email)); err != nil {
			return err
		}

		emailService.Code = ""

		return nil
	})
}

// save the verify code to redis and attach expired time
// cause when the timing out redis will remove it from the container
// so this is a simple timing task
func (emailService *EmailService) SaveCode() {
	fmt.Println(emailService.Code)

	redis_conn.GetRedisConnMust(func(conn redis.Conn) error {
		_, err := conn.Do("SETEX", genCodeKey(emailService.Email), VALID_TIMING, emailService.Code)

		return err
	})
}

// send the verify code to target email
func (emailService *EmailService) SendVerifyCode() error {
	file, err := os.Open("templates/email-valid.html")
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	verifyCode := utils.CreateVerifyCode()

	r, _ := regexp.Compile("{{ verifyCode }}")

	sender := email_sender.EmailSender{
		Addr:        emailService.Email,
		Subject:     "请完成邮箱认证",
		Body:        r.ReplaceAllString(string(content), verifyCode),
		ContentType: "text/html",
	}

	if err := sender.SendEmail(); err != nil && err != io.EOF {
		fmt.Println(err == io.EOF, err)

		return err
	}

	emailService.Code = verifyCode

	return nil
}

// generate a key that as a key of redis
func genCodeKey(email string) string {
	return fmt.Sprintf("verify_code_%s", email)
}
