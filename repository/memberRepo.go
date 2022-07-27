package repository

import (
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -destination=../../test/mock/member_mock_repository.go -package=mock GoWeb/repository IMemberRepo

type IMemberRepo interface {
	CreateMember() bool
	GetMember() bool
	UpdateMember() bool
	DeleteMember() bool
}

//var _ IMemberRepo = (*MemberRepo)(nil)

type MemberRepo struct {
	dB *gorm.DB
}

func NewMemberRepo(db *gorm.DB) IMemberRepo {
	return &MemberRepo{
		dB: db,
	}
}

func (m *MemberRepo) CreateMember() bool {
	return true
}
func (m *MemberRepo) GetMember() bool {
	return false
}
func (m *MemberRepo) UpdateMember() bool {
	return false
}
func (m *MemberRepo) DeleteMember() bool {
	return false
}
