package service

import (
	rep "GoWeb/repository/postgredb"
	"time"
)

type ILoginSrv interface {
}

var (
	cancelTimeout time.Duration = 3 // default 3 second
)

type LoginSrv struct {
	MemberRepo rep.IMemberRepo
}

func NewLoginSrv(IMemberRepo rep.IMemberRepo) ILoginSrv {
	return &LoginSrv{
		MemberRepo: IMemberRepo,
	}
}
