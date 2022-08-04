package main

import (
	"GoWeb/database"
	rep "GoWeb/repository/postgredb"
	"GoWeb/router"
	srv "GoWeb/service"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// new postgres db
	db, postgreErr := database.NewDb()
	if postgreErr != nil {
		return
	}
	// new redis db
	_, redisErr := database.NewRedis()
	if redisErr != nil {
		return
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
	//Router
	Router := router.NewRouter(MemberSrv)

	return Router

}
