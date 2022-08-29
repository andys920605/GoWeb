package service

import (
	"GoWeb/infras/configs"
	models_const "GoWeb/models"
	models_svc "GoWeb/models/service"
	rep "GoWeb/repository/interface"
	svc_interface "GoWeb/service/interface"
	"GoWeb/utils/crypto"
	"GoWeb/utils/errs"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginSrv struct {
	cfg       *configs.Config
	memberRep rep.IMemberRep
	cacheRep  rep.ICacheRep
}

// jwt secret key
var JwtSecret = []byte("secret")

func NewLoginSvc(config *configs.Config, IMemberRep rep.IMemberRep, ICacheRep rep.ICacheRep) svc_interface.ILoginSrv {
	return &LoginSrv{
		cfg:       config,
		memberRep: IMemberRep,
		cacheRep:  ICacheRep,
	}
}

func (svc *LoginSrv) Login(param *models_svc.LoginReq) (*models_svc.Scepter, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	// check token exist
	if result, _ := svc.cacheRep.GetTokenByIDCtx(context.Background(), models_const.CacheTokenClientId+param.Account); result != nil {
		token, _ := svc.hashClaims(result)
		if token != nil {
			return &models_svc.Scepter{
				AccessToken: *token,
				TokenType:   "Bearer",
			}, nil
		}
	}
	member, findErr := svc.memberRep.Find(ctx, &param.Account, nil)
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
	claims := &models_svc.Claims{
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
	token, err := svc.hashClaims(claims)
	if err != nil {
		return nil, err
	}
	redisModel := *claims
	svc.cacheRep.SetTokenCtx(ctx, models_const.CacheTokenClientId+param.Account, svc.getCacheTime(), &redisModel)
	return &models_svc.Scepter{
		AccessToken: *token,
		TokenType:   "Bearer",
	}, nil
}

// region private function

// get cache time
// return seconds
func (svc *LoginSrv) getCacheTime() int {
	return int((time.Duration(svc.cfg.Redis.CacheTTL) * time.Minute).Seconds())
}

// hash claims
func (svc *LoginSrv) hashClaims(claims *models_svc.Claims) (*string, *errs.ErrorResponse) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	if err != nil {
		return nil, &errs.ErrorResponse{
			Message: fmt.Sprintf("jwt err:%s", err.Error()),
		}
	}
	return &token, nil
}

// endregion
