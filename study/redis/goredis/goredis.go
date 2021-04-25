package goredis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var client *redis.Client

func GetRedis() *redis.Client {
	return client
}

func StatRedis() *redis.PoolStats {
	return client.PoolStats()
}

func CloseRedis() error {
	if client != nil {
		return client.Close()
	}
	return nil
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     8,
		MinIdleConns: 2,
		PoolTimeout:  3 * time.Second,
		MaxConnAge:   3600 * time.Second,
	})
}
