package service

import (
	models_rep "GoWeb/models/repository"
	rep "GoWeb/repository/interface"
	svc_interface "GoWeb/service/interface"
	"GoWeb/utils/crypto"
	"GoWeb/utils/errs"
	"context"
	"fmt"
	"net/http"
	"time"
)

var (
	cancelTimeout time.Duration = 3 // default 3 second
)

type MemberSrv struct {
	MemberRepo rep.IMemberRep
}

func NewMemberSrv(IMemberRepo rep.IMemberRep) svc_interface.IMemberSrv {
	return &MemberSrv{
		MemberRepo: IMemberRepo,
	}
}

func (svc *MemberSrv) CreateMember(param *models_rep.Member) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	hash := crypto.NewSHA256([]byte(param.Password))
	param.Password = fmt.Sprintf("%x", hash[:])
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
