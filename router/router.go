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
	v1.GET("/members", router.getMembers)
	v1.POST("/members", router.createMember)
	return r
}

// v1/member
func (router *Router) getMembers(c *gin.Context) {
	result, errRsp := router.MemberSvc.GetAllMember()
	if errRsp != nil {
		c.JSON(http.StatusInternalServerError, errRsp)
		return
	}
	c.JSON(http.StatusOK, result)
}
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
