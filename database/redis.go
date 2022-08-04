package database

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("RD_ADDR"),
		MinIdleConns: 200,
		PoolSize:     12000,
		PoolTimeout:  240,
		Password:     os.Getenv("RD_PASSWORD"), // no password set
		DB:           0,                        // use default DB
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	} else {
		log.Printf("Redis connected")
	}
	return client, nil
}
