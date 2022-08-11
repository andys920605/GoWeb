package service

import (
	models_rep "GoWeb/models/repository"
	rep "GoWeb/repository/postgredb"
	"context"
	"fmt"
	"net/http"
	"time"

	"GoWeb/utils/errs"
)

type IMemberSrv interface {
	CreateMember(*models_rep.Member) *errs.ErrorResponse
	GetAllMember() (*[]models_rep.Member, *errs.ErrorResponse)
	GetMember(*string, *string) (*models_rep.Member, *errs.ErrorResponse)
	UpdateMember(*models_rep.UpdateMember) *errs.ErrorResponse
	DisableMember(*models_rep.UpdateMember) *errs.ErrorResponse
}

var (
	cancelTimeout time.Duration = 3 // default 3 second
)

type MemberSrv struct {
	MemberRepo rep.IMemberRepo
}

func NewMemberSrv(IMemberRepo rep.IMemberRepo) IMemberSrv {
	return &MemberSrv{
		MemberRepo: IMemberRepo,
	}
}

func (svc *MemberSrv) CreateMember(param *models_rep.Member) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	if err := svc.MemberRepo.Insert(ctx, param); err != nil {
		return &errs.ErrorResponse{
			Message: err.Error(),
		}
	}
	return nil
}
func (svc *MemberSrv) GetAllMember() (*[]models_rep.Member, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	result, err := svc.MemberRepo.FindAll(ctx)
	if err != nil {
		return nil, &errs.ErrorResponse{
			Message: err.Error(),
		}
	}
	return result, nil
}
func (svc *MemberSrv) GetMember(account *string, phone *string) (*models_rep.Member, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	result, err := svc.MemberRepo.Find(ctx, account, phone)
	if err != nil {
		return nil, &errs.ErrorResponse{
			Message: err.Error(),
		}
	}
	return result, nil
}
func (svc *MemberSrv) UpdateMember(param *models_rep.UpdateMember) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	if _, err := svc.MemberRepo.Find(ctx, &param.Account, nil); err != nil {
		return &errs.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("The account %s does not exits", param.Account),
		}
	}
	if err := svc.MemberRepo.Updates(ctx, param); err != nil {
		return &errs.ErrorResponse{
			Message: err.Error(),
		}
	}
	return nil
}
func (svc *MemberSrv) DisableMember(param *models_rep.UpdateMember) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	if _, err := svc.MemberRepo.Find(ctx, &param.Account, nil); err != nil {
		return &errs.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("The account %s does not exits", param.Account),
		}
	}
	if err := svc.MemberRepo.Disable(ctx, param); err != nil {
		return &errs.ErrorResponse{
			Message: err.Error(),
		}
	}
	return nil
}
