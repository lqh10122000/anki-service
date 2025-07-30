package service

import (
	"gopkg.in/gomail.v2"
)

type MailService interface {
	SendMail(to string, subject string, body string) error
}

type mailService struct {
	smtpHost string
	smtpPort int
	username string
	password string
	from     string
}

func NewMailService(host string, port int, username, password, from string) MailService {
	return &mailService{
		smtpHost: host,
		smtpPort: port,
		username: username,
		password: password,
		from:     from,
	}
}

func (m *mailService) SendMail(to string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.smtpHost, m.smtpPort, m.username, m.password)

	return dialer.DialAndSend(msg)
}
