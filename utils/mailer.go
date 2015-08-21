package utils

import (
	"github.com/playaer/myFirstGoProject/models"
	"net/smtp"
	"github.com/playaer/myFirstGoProject/config"
	"bytes"
	"html/template"
	"fmt"
)

type Mailer struct {
	config *config.Config
}

type mail struct {
	From string
	To []string
	Message []byte
}

func (self *Mailer) BuildRegistrationMail(user *models.User) *mail {
	config := self.config
	mailConfig := new(mail)

	mailConfig.From = config.EmailUser
	mailConfig.To = []string{user.Email}

	tplData := map[string]string {
		"From": config.EmailUser,
		"To": user.Email,
		"Subject": "Confirm registration",
		"Message": "your link",
	}

	buffer := new(bytes.Buffer)

	t, _ := template.ParseFiles("mailer/view.html")
	t.Execute(buffer, tplData)

	mailConfig.Message = buffer.Bytes()

	return mailConfig
}

func (self *Mailer) Send(mailConfig *mail) {
	config := self.config
	auth := smtp.PlainAuth("", config.EmailUser, config.EmailPassword, config.EmailHost)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", config.EmailHost, config.EmailPort),
		auth,
		config.EmailUser,
		mailConfig.To,
		mailConfig.Message)

	CheckErr(err, "ERROR: attempting to send a mail ")
}

