package service_interface

import (
	models_srv "GoWeb/models/service"
	"GoWeb/utils/errs"
)

type ILoginSrv interface {
	Login(*models_srv.LoginReq) (*models_srv.Scepter, *errs.ErrorResponse)
	Logout(*string) *errs.ErrorResponse
	CheckTokenExist(string) *string
}
