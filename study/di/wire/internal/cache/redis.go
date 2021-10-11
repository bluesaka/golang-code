package cache

import (
	"context"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"go-code/study/di/wire/internal/config"
	"time"
)

var (
	Provider  = wire.NewSet(NewRedis)
	redisPool *redigo.Pool
)

func NewRedis(cfg *config.Config) *redigo.Pool {
	redisPool = &redigo.Pool{
		MaxIdle:     cfg.Redis.MaxIdle,
		MaxActive:   cfg.Redis.MaxActive,
		IdleTimeout: time.Duration(cfg.Redis.IdleTimeout) * time.Second,
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", cfg.Redis.Addr, redigo.DialDatabase(cfg.Redis.DB), redigo.DialPassword(cfg.Redis.Password))
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Redis.DialTimeout)*time.Second)
	defer cancel()
	done := make(chan struct{}, 1)

	go func() {
		redisConn := GetRedis()
		if r, _ := redigo.String(redisConn.Do("PING")); r != "PONG" {
			panic("redis connect failed")
		}
		redisConn.Close()
		done <- struct{}{}
	}()

	select {
	case <-done:
		return redisPool
	case <-ctx.Done():
		panic("redis connect failed")
	}
}

// GetRedis returns redis connection
func GetRedis() redigo.Conn {
	return redisPool.Get()
}

// CloseRedis close redis pool
func CloseRedis() error {
	if redisPool != nil {
		return redisPool.Close()
	}
	return nil
}
