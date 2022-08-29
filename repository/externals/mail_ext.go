package externals

import (
	"GoWeb/infras/configs"
	models_ext "GoWeb/models/externals"
	rep_interface "GoWeb/repository/interface"
	"log"
	"net/smtp"
)

type MailExt struct {
	cfg *configs.Config
}

func NewMailExt(config *configs.Config) rep_interface.IMailExt {
	return &MailExt{
		cfg: config,
	}
}

func (rep *MailExt) Send(mail *models_ext.SendMail) bool {
	from := rep.cfg.Email.Account
	pass := rep.cfg.Email.Password
	to := mail.TargetAddress

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:" + mail.Title + "\n\n" +
		mail.Body

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
