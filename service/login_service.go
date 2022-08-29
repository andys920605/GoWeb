package service

import (
	"GoWeb/infras/configs"
	models_const "GoWeb/models"
	models_ext "GoWeb/models/externals"
	models_svc "GoWeb/models/service"
	rep "GoWeb/repository/interface"
	svc_interface "GoWeb/service/interface"
	"GoWeb/utils"
	"GoWeb/utils/crypto"
	"GoWeb/utils/errs"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/xorcare/pointer"
)

type LoginSrv struct {
	cfg       *configs.Config
	memberRep rep.IMemberRep
	cacheRep  rep.ICacheRep
	mailExt   rep.IMailExt
}

// jwt secret key
var JwtSecret = []byte("secret")

func NewLoginSvc(config *configs.Config, IMemberRep rep.IMemberRep, ICacheRep rep.ICacheRep, IMailExt rep.IMailExt) svc_interface.ILoginSrv {
	return &LoginSrv{
		cfg:       config,
		memberRep: IMemberRep,
		cacheRep:  ICacheRep,
		mailExt:   IMailExt,
	}
}

// 使用者登入
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
			ExpiresAt: now.Add(120 * time.Minute).Unix(),
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

// 使用者登出
func (svc *LoginSrv) Logout(account *string) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	if err := svc.cacheRep.DeleteCtx(ctx, models_const.CacheTokenClientId+*account); err != nil {
		return &errs.ErrorResponse{
			Message: fmt.Sprintf("Logout fail,%s", err.Error()),
		}
	}
	return nil
}

// 寄送信箱認證信
func (svc *LoginSrv) SendVerificationLetter(email *string) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	verifyCode := utils.Rand(8, utils.RAND_KIND_ALL)
	body := fmt.Sprintf("驗證碼:%s，五分鐘後失效。", verifyCode)
	mail := &models_ext.SendMail{
		TargetAddress: *email,
		Title:         "信箱驗證信",
		Body:          body,
	}
	if err := svc.mailExt.Send(mail); err != nil {
		return err
	}
	redisModel := *claims
	svc.cacheRep.SetTokenCtx(ctx, models_const.CacheTokenClientId+param.Account, svc.getCacheTime(), &redisModel)

}

// 給middleware驗證帳號憑證是否過期
func (svc *LoginSrv) CheckTokenExist(account string) *string {
	// check token exist
	if result, _ := svc.cacheRep.GetTokenByIDCtx(context.Background(), models_const.CacheTokenClientId+account); result == nil {
		return pointer.String("Invalid Token.")
	}
	return nil
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
