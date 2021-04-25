package redigo

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

func GetRedis() redis.Conn {
	return redisPool.Get()
}

func StatRedis() redis.PoolStats {
	return redisPool.Stats()
}

func CloseRedis() error {
	if redisPool != nil {
		return redisPool.Close()
	}
	return nil
}

func init() {
	redisPool = &redis.Pool{
		MaxIdle:     2,
		MaxActive:   5,
		IdleTimeout: 3600 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0), redis.DialPassword(""))
		},
	}

	conn := GetRedis()
	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		panic("redis connect failed")
	}
}
