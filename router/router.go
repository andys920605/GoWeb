package router

import (
	srv "GoWeb/service"

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
	v1.GET("/member", router.getMember)
	return r
}

func (router *Router) getMember(c *gin.Context) {
	check := router.MemberSvc.CreateMember()
	if check {
		c.JSON(200, gin.H{
			"message": "Ok",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "NotOk",
		})
	}
}
