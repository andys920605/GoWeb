package database

import (
	model_com "GoWeb/models/commons"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedis(Opt *model_com.Options) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         Opt.Config.Redis.RedisAddr,
		MinIdleConns: Opt.Config.Redis.MinIdleConns,
		PoolSize:     Opt.Config.Redis.PoolSize,
		PoolTimeout:  time.Duration(Opt.Config.Redis.PoolTimeout) * time.Second,
		Password:     Opt.Config.Redis.Password, // no password set
		DB:           Opt.Config.Redis.DB,       // use default DB
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	} else {
		log.Printf("Redis connected")
	}
	return client, nil
}
