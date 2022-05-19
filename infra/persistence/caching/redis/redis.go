package caching

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result, _ := rdb.Ping(ctx).Result()
	log.Println("Connectado", result)

	return rdb
}
