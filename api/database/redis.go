package database

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(addr, password string, db, poolSize int) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: poolSize,
	})
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	return err
}
