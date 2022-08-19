package main

import (
	"GoWeb/database"
	"GoWeb/infras/configs"
	rep "GoWeb/repository/postgredb"
	"GoWeb/router"
	srv "GoWeb/service"
	"GoWeb/utils"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

var (
	redisBool = true
	cfgTemp   *configs.Config
)

func main() {
	config := ProvideConfig()

	// new postgres db
	db, postgreErr := database.NewDb(config)
	if postgreErr != nil {
		return
	}
	//new redis db
	if redisBool {
		_, redisErr := database.NewRedis(config)
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

// new config
func ProvideConfig() *configs.Config {
	buffer, err := configs.LoadConfig(utils.GetConfigPath())
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfgTemp, err = configs.ParseConfig(buffer)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	return cfgTemp
}
