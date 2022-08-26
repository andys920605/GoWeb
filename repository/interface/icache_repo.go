package repositories_interface

import (
	"context"
)

//go:generate mockgen -destination=../../test/mock/icache_mock_repository.go -package=mock GoWeb/repository/interface ICacheRepo
type ICacheRepo interface {
	DeleteCtx(context.Context, string) error
}
