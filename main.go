package main

import (
	"GoWeb/database"
	"GoWeb/infras/configs"
	rep_db "GoWeb/repository/postgredb"
	rep_redis "GoWeb/repository/redisdb"
	"GoWeb/router"
	srv "GoWeb/service"
	"GoWeb/utils"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

var (
	redisBool = false
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
	redis, redisErr := database.NewRedis(config)
	if redisErr != nil {
		return
	}
	app := di(config, db, redis)
	server := app.InitRouter()
	server.Run(":8070")
}

func di(cfg *configs.Config, db *gorm.DB, redis *redis.Client) router.IRouter {
	//Repo
	MemberRepo := rep_db.NewMemberRepo(db)
	CacheRep := rep_redis.NewCacheRepository(redis)
	//Srv
	MemberSrv := srv.NewMemberSrv(MemberRepo)
	LoginSrv := srv.NewLoginSrv(cfg, MemberRepo, CacheRep)
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
