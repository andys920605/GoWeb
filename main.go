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
	app.InitRouter().Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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
