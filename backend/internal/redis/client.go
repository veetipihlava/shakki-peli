package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	RedisAddr = "localhost:6379"
	RedisDB   = 0
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func InitializeRedis() *RedisClient {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   RedisDB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &RedisClient{
		Client: client,
		Ctx:    ctx,
	}
}
