package repositories_interface

import (
	models_ext "GoWeb/models/externals"
	"GoWeb/utils/errs"
)

//go:generate mockgen -destination=../../test/mock/imail_mock_repository.go -package=mock GoWeb/repository/interface IMailRep
type IMailExt interface {
	Send(*models_ext.SendMail) *errs.ErrorResponse
}
