package router

import (
	models_rep "GoWeb/models/repository"
	models_srv "GoWeb/models/service"
	"GoWeb/router/middlewares"
	svc "GoWeb/service/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	InitRouter() *gin.Engine
}

type Router struct {
	MemberSvc svc.IMemberSrv
	LoginSvc  svc.ILoginSrv
}

func NewRouter(IMemberSrv svc.IMemberSrv, ILoginSrv svc.ILoginSrv) IRouter {
	return &Router{
		MemberSvc: IMemberSrv,
		LoginSvc:  ILoginSrv,
	}
}

func (router *Router) InitRouter() *gin.Engine {
	r := gin.Default()
	g1 := r.Group("/g1")                    // 不用token
	g2 := r.Group("/g2")                    // 要token
	g2.Use(middlewares.JWTAuthMiddleware()) // use the Bearer Authentication middleware
	g1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 會員
	g1.POST("/member", router.createMember)
	g2.GET("/member", router.getMember)
	g2.PUT("/member/:account", router.updateMember)
	g2.DELETE("/member/:account", router.disableMember)
	// Login
	g1.POST("/login", router.login)
	return r
}

// region CRUD會員資料
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
	result := c.MustGet("account").(*models_srv.Claims)
	println(result)
	account := c.Query("account")
	phone := c.Query("phone")
	if account != "" || phone != "" {
		result, errRsp := router.MemberSvc.GetMember(&account, &phone)
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
		c.JSON(errRsp.StatusCode, errRsp)
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
