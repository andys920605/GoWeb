package router

import (
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
	v1.GET("/member", router.getAllMember)
	return r
}

// v1/member
func (router *Router) getAllMember(c *gin.Context) {
	result, errRsp := router.MemberSvc.GetAllMember()
	if errRsp != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errRsp})
		return
	}
	c.JSON(http.StatusOK, result)
}
