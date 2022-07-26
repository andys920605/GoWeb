package service

import rep "GoWeb/repository"

type IMemberSrv interface {
	CreateMember() bool
	GetMember() bool
}

var _ IMemberSrv = (*MemberSrv)(nil)

type MemberSrv struct {
	MemberRepo rep.IMemberRepo
}

func NewMemberSrv(IMemberRepo rep.IMemberRepo) IMemberSrv {
	return &MemberSrv{
		MemberRepo: IMemberRepo,
	}
}
func (m *MemberSrv) CreateMember() bool {
	return m.MemberRepo.CreateMember()
}
func (m *MemberSrv) GetMember() bool {
	return false
}
