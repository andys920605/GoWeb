package database

import (
	"GoWeb/infras/configs"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedis(cfg *configs.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.RedisAddr,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password, // no password set
		DB:           cfg.Redis.DB,       // use default DB
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	} else {
		log.Printf("Redis connected")
	}
	return client, nil
}
