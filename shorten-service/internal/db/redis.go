package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic("Failed to Connect to Redis " + err.Error())
	}

}
