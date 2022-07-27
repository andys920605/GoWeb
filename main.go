package main

import (
	"GoWeb/database"
	rep "GoWeb/repository"
	"GoWeb/router"
	srv "GoWeb/service"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.NewDb()

	app := di()
	server := app.InitRouter()
	server.Run(":8070")
}

func di() router.IRouter {
	//Repo
	MemberRepo := rep.NewMemberRepo(database.DB)
	//Srv
	MemberSrv := srv.NewMemberSrv(MemberRepo)
	//Router
	Router := router.NewRouter(MemberSrv)

	return Router

}
