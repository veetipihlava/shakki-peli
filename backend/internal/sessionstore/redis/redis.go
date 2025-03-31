package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/veetipihlava/shakki-peli/internal/sessionstore"
)

const (
	RedisAddr = "localhost:6379"
	RedisDB   = 0
)

type Redis struct {
	Client *redis.Client
	Ctx    context.Context
}

// Compile-level check that Redis does implement SessionStore
var _ sessionstore.SessionStore = (*Redis)(nil)

func InitializeRedis() (*Redis, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   RedisDB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{
		Client: client,
		Ctx:    ctx,
	}, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
