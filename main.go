package main

import (
	"GoWeb/database"
	"GoWeb/infras/configs"
	"GoWeb/infras/logger"
	model_com "GoWeb/models/commons"
	rep_ext "GoWeb/repository/externals"
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
	//redisBool = false
	cfgTemp *configs.Config
)

func main() {
	// init config and logger
	config := ProvideConfig()
	apiLogger := logger.NewApiLogger(config)
	options := &model_com.Options{
		Config: config,
		Logger: apiLogger,
	}
	// new postgres db
	db, postgreErr := database.NewDb(options)
	if postgreErr != nil {
		return
	}
	// new redis db
	redis, redisErr := database.NewRedis(options)
	if redisErr != nil {
		return
	}
	app := di(options, db, redis)
	server := app.InitRouter()
	server.Run(":8070")
}

func di(opt *model_com.Options, db *gorm.DB, redis *redis.Client) router.IRouter {
	// Repo
	MailRep := rep_ext.NewMailExt(opt)
	MemberRep := rep_db.NewMemberRep(db)
	CacheRep := rep_redis.NewCacheRepository(redis)
	// Svc
	MemberSvc := srv.NewMemberSvc(MemberRep)
	LoginSvc := srv.NewLoginSvc(opt, MemberRep, CacheRep, MailRep)
	// Router
	Router := router.NewRouter(MemberSvc, LoginSvc)

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
