package externals

import (
	"GoWeb/infras/configs"
	model_com "GoWeb/models/commons"
	models_ext "GoWeb/models/externals"
	rep_interface "GoWeb/repository/interface"
	"GoWeb/utils/errs"
	"fmt"
	"log"
	"net/smtp"
)

type MailExt struct {
	cfg *configs.Config
}

func NewMailExt(opt *model_com.Options) rep_interface.IMailExt {
	return &MailExt{
		cfg: opt.Config,
	}
}

func (rep *MailExt) Send(mail *models_ext.SendMail) *errs.ErrorResponse {
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
		return &errs.ErrorResponse{
			Message: fmt.Sprintf("smtp error: %s", err.Error()),
		}
	}

	return nil
}
