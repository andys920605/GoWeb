package service

import (
	models_srv "GoWeb/models/service"
	rep "GoWeb/repository/postgredb"
	"GoWeb/utils/crypto"
	"GoWeb/utils/errs"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type ILoginSrv interface {
	Login(*models_srv.LoginReq) (*models_srv.Scepter, *errs.ErrorResponse)
}

type LoginSrv struct {
	MemberRepo rep.IMemberRepo
}

// jwt secret key
var JwtSecret = []byte("secret")

func NewLoginSrv(IMemberRepo rep.IMemberRepo) ILoginSrv {
	return &LoginSrv{
		MemberRepo: IMemberRepo,
	}
}
func (svc *LoginSrv) Login(param *models_srv.LoginReq) (*models_srv.Scepter, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	member, findErr := svc.MemberRepo.Find(ctx, &param.Account, nil)
	if findErr != nil {
		return nil, &errs.ErrorResponse{
			Message: "查無此帳號",
		}
	}
	hash := crypto.NewSHA256([]byte(param.Password))
	if member.Password != fmt.Sprintf("%x", hash[:]) {
		return nil, &errs.ErrorResponse{
			Message: "密碼錯誤",
		}
	}
	now := time.Now()
	jwtId := param.Account + strconv.FormatInt(now.Unix(), 10)
	// set claims and sign
	claims := &models_srv.Claims{
		Account: param.Account,
		Role:    strconv.Itoa(member.Permission),
		StandardClaims: jwt.StandardClaims{
			Audience:  param.Account,
			ExpiresAt: now.Add(30 * time.Minute).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(1 * time.Second).Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	if err != nil {
		return nil, &errs.ErrorResponse{
			Message: fmt.Sprintf("jwt err:%s", err.Error()),
		}
	}
	return &models_srv.Scepter{
		AccessToken: token,
		TokenType:   "Bearer",
	}, nil
}
