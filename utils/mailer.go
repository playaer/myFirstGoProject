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

func (self *Mailer) New(conf *config.Config) *Mailer {
	return &Mailer{conf}
}

type mail struct {
	From string
	To []string
	Message []byte
}

// Prepare email before send
func (self *Mailer) BuildRegistrationMail(user *models.User) *mail {
	config := self.config
	mailConfig := new(mail)

	mailConfig.From = config.EmailUser
	mailConfig.To = []string{user.Email}

	tplData := map[string]string {
		"From": config.EmailUser,
		"To": user.Email,
		"Subject": "Confirm registration",
		"Hash": user.Hash,
		"Domain": config.SiteDomain,
	}

	buffer := new(bytes.Buffer)
	t, err := template.ParseFiles("templates/mailer/confirm_registration.tmpl")
	CheckErr(err, "fff")
	t.Execute(buffer, tplData)

	mailConfig.Message = buffer.Bytes()

	return mailConfig
}

// Send email
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

