package main

import (
	"GoWeb/database"
	rep "GoWeb/repository"
	"GoWeb/router"
	srv "GoWeb/service"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// new postgres db
	db, err := database.NewDb()
	if err != nil {
		return
	}
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
