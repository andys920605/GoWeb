package repositories_interface

import (
	models_svc "GoWeb/models/service"
	"context"
)

//go:generate mockgen -destination=../../test/mock/icache_mock_repository.go -package=mock GoWeb/repository/interface ICacheRep
type ICacheRep interface {
	GetTokenByIDCtx(context.Context, string) (*models_svc.Claims, error)
	SetTokenCtx(context.Context, string, int, *models_svc.Claims) error

	GetEmailCodeCtx(context.Context, string) ([]string, error)
	// @ctx@key@sec@item
	SetEmailCodeCtx(context.Context, string, int, string) error

	DeleteCtx(context.Context, string) error
}
