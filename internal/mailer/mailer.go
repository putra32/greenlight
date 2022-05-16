package mailer

import (
	"embed"
	"fmt"

	"gopkg.in/gomail.v2"
)

var templateFS embed.FS

type Mailer struct {
	dialer *gomail.Dialer
	sender string
}

func New(host string, port int, username, password, sender string) Mailer {
	dialer := gomail.NewDialer(host, port, username, password)

	return Mailer{
		dialer: dialer,
		sender: sender,
	}
}

func (m Mailer) Send(recipient string, data map[string]interface{}) error {
	// tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	// if err != nil {
	// 	return err
	// }

	// subject := new(bytes.Buffer)
	// err = tmpl.ExecuteTemplate(subject, "subject", data)
	// if err != nil {
	// 	return err
	// }

	// plainBody := new(bytes.Buffer)
	// err = tmpl.ExecuteTemplate(plainBody, "plaintBody", data)
	// if err != nil {
	// 	return err
	// }

	// htmlBody := new(bytes.Buffer)
	// err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	// if err != nil {
	// 	return err
	// }

	bodyMessage := fmt.Sprintf("Hello, <b>have a nice day</b> set you token {%s}", data["activationToken"].(string))

	msg := gomail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", "test mail")
	msg.SetBody("text/html", bodyMessage)
	// msg.AddAlternative("text/html", htmlBody.String())

	err := m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}
	return nil
}
