package datastore

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	Rdb *redis.Client
	ctx = context.Background()
)

func InitRedisClient() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
	})
}