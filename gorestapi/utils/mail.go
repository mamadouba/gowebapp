package utils

import (
	"bytes"
	"fmt"
	"gorestapi/config"
	"gorestapi/logger"
	"html/template"
	"net/smtp"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendGmail(tmpl string, subject string, to []string, data interface{}) error {
	conf := config.Configuration
	authAddr := strings.Split(conf.MailHost, ":")[0]
	auth := smtp.PlainAuth("", conf.MailUser, conf.MailPass, authAddr)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type:text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, headers)))
	err = t.Execute(&body, data)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = smtp.SendMail(conf.MailHost, auth, conf.MailUser, to, body.Bytes())
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("Mail sent to %s", to)
	return nil
}
func SendMail(tmpl string, subject string, to string, data interface{}) error {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	var htmlContent bytes.Buffer
	err = t.Execute(&htmlContent, data)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	from := mail.NewEmail("From", os.Getenv("SENDGRID_USER"))
	message := mail.NewSingleEmail(from, subject, mail.NewEmail("To", to), "Hello", htmlContent.String())
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_KEY"))
	_, err = client.Send(message)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("Mail sent")
	return nil
}
