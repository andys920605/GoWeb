package externals

import (
	"GoWeb/infras/configs"
	rep_interface "GoWeb/repository/interface"
	"log"
	"net/smtp"
)

type MailRep struct {
	cfg *configs.Config
}

func NewMailRep(config *configs.Config) rep_interface.IMailRep {
	return &MailRep{
		cfg: config,
	}
}

func (rep *MailRep) Send(body string) bool {
	from := rep.cfg.Email.Account
	pass := rep.cfg.Email.Password
	to := "andys920605@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	log.Print("sent, visit mail")
	return true
}
