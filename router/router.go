package router

import (
	models_rep "GoWeb/models/repository"
	models_srv "GoWeb/models/service"
	srv "GoWeb/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	InitRouter() *gin.Engine
}

type Router struct {
	MemberSvc srv.IMemberSrv
	LoginSvc  srv.ILoginSrv
}

func NewRouter(IMemberSrv srv.IMemberSrv, ILoginSrv srv.ILoginSrv) IRouter {
	return &Router{
		MemberSvc: IMemberSrv,
		LoginSvc:  ILoginSrv,
	}
}

func (router *Router) InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 會員
	v1.POST("/members", router.createMember)
	v1.GET("/members", router.getMember)
	v1.PUT("/members/:account", router.updateMember)
	v1.DELETE("/members/:account", router.disableMember)
	// Login
	v1.POST("/login", router.login)
	return r
}

// region CRUD會員資料
// v1/members
func (router *Router) createMember(c *gin.Context) {
	var payload models_rep.Member
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	errRsp := router.MemberSvc.CreateMember(&payload)
	if errRsp != nil {
		c.JSON(http.StatusInternalServerError, errRsp)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func (router *Router) getMember(c *gin.Context) {
	account := c.Query("account")
	phone := c.Query("phone")
	if account != "" || phone != "" {
		result, errRsp := router.MemberSvc.GetMember(account, phone)
		if errRsp != nil {
			c.JSON(http.StatusInternalServerError, errRsp)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		result, errRsp := router.MemberSvc.GetAllMember()
		if errRsp != nil {
			c.JSON(http.StatusInternalServerError, errRsp)
			return
		}
		c.JSON(http.StatusOK, result)
	}

}
func (router *Router) updateMember(c *gin.Context) {
	account := c.Param("account")
	var payload models_rep.UpdateMember
	payload.Account = account
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	errRsp := router.MemberSvc.UpdateMember(&payload)
	if errRsp != nil {
		c.JSON(http.StatusInternalServerError, errRsp)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func (router *Router) disableMember(c *gin.Context) {
	account := c.Param("account")
	var payload models_rep.UpdateMember
	payload.Account = account
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	errRsp := router.MemberSvc.DisableMember(&payload)
	if errRsp != nil {
		c.JSON(http.StatusInternalServerError, errRsp)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// endregion

// region Login
// v1/login
func (router *Router) login(c *gin.Context) {
	var payload models_srv.LoginReq
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, errRsp := router.LoginSvc.Login(&payload)
	if errRsp != nil {
		c.JSON(http.StatusInternalServerError, errRsp)
		return
	}
	c.JSON(http.StatusOK, token)
}

// endregion
