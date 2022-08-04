package service

import (
	rep "GoWeb/repository/postgredb"
)

type ILoginSrv interface {
}

type LoginSrv struct {
	MemberRepo rep.IMemberRepo
}

func NewLoginSrv(IMemberRepo rep.IMemberRepo) ILoginSrv {
	return &LoginSrv{
		MemberRepo: IMemberRepo,
	}
}
