package service

import (
	models_rep "GoWeb/models/repository"
	rep "GoWeb/repository"
	"context"
	"time"

	"GoWeb/utils/errs"
)

type IMemberSrv interface {
	CreateMember() bool
	GetAllMember() (*[]models_rep.Member, *errs.ErrorResponse)
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
func (svc *MemberSrv) CreateMember() bool {
	return svc.MemberRepo.Insert()
}
func (svc *MemberSrv) GetAllMember() (*[]models_rep.Member, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	result, err := svc.MemberRepo.FindAll(ctx)
	if err != nil {
		return nil, errs.ErrNotFound
	}
	return result, nil
}
