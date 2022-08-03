package router

import (
	models_rep "GoWeb/models/repository"
	srv "GoWeb/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	InitRouter() *gin.Engine
}

type Router struct {
	MemberSvc srv.IMemberSrv
}

func NewRouter(IMemberSrv srv.IMemberSrv) IRouter {
	return &Router{
		MemberSvc: IMemberSrv,
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
	v1.POST("/members", router.createMember)
	v1.GET("/members", router.getMember)
	v1.PUT("/members/:account", router.updateMember)
	return r
}

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
