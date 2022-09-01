package repositories_interface

import (
	models_rep "GoWeb/models/repository"
	"context"
)

//go:generate mockgen -destination=../../test/mock/imember_mock_repository.go -package=mock GoWeb/repository/interface IMemberRep
type IMemberRep interface {
	Insert(context.Context, *models_rep.Member) error
	Find(context.Context, *string, *string) (*models_rep.Member, error)
	FindAll(context.Context) (*[]models_rep.Member, error)
	Updates(context.Context, *models_rep.UpdateMember) error
	Disable(context.Context, *models_rep.UpdateMember) error
	CheckAccountExist(context.Context, string) bool
	CheckEmailExist(context.Context, string) bool
	Close()
}
