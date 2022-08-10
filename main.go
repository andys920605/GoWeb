package main

import (
	"GoWeb/database"
	rep "GoWeb/repository/postgredb"
	"GoWeb/router"
	srv "GoWeb/service"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

var (
	redisBool = true
)

func main() {
	// new postgres db
	db, postgreErr := database.NewDb()
	if postgreErr != nil {
		return
	}
	//new redis db
	if redisBool {
		_, redisErr := database.NewRedis()
		if redisErr != nil {
			return
		}
	}
	app := di(db)
	server := app.InitRouter()
	server.Run(":8070")
}

func di(db *gorm.DB) router.IRouter {
	//Repo
	MemberRepo := rep.NewMemberRepo(db)
	//Srv
	MemberSrv := srv.NewMemberSrv(MemberRepo)
	LoginSrv := srv.NewLoginSrv(MemberRepo)
	//Router
	Router := router.NewRouter(MemberSrv, LoginSrv)

	return Router

}
