package service_interface

import (
	models_rep "GoWeb/models/repository"
	"GoWeb/utils/errs"
)

type IMemberSrv interface {
	CreateMember(*models_rep.Member) *errs.ErrorResponse
	GetAllMember() (*[]models_rep.Member, *errs.ErrorResponse)
	GetMember(*string, *string) (*models_rep.Member, *errs.ErrorResponse)
	UpdateMember(*models_rep.UpdateMember) *errs.ErrorResponse
	DisableMember(*models_rep.UpdateMember) *errs.ErrorResponse
}
