package main

import (
	"github.com/spf13/cast"
	"github.com/tal-tech/go-zero/core/logx"
	zredis "github.com/tal-tech/go-zero/core/stores/redis"
	"log"
	"reflect"
)

func main()  {
	redis := zredis.NewRedis("127.0.0.1:6379", zredis.NodeType, "")
	//get(redis)
	//Hget(redis)
	redisLock(redis)
}

func get(redis *zredis.Redis) {
	_ = redis.Set("test1", "111")
	v, err := redis.Get("test1")
	log.Println(v, err)
	log.Println(reflect.TypeOf(v))
	log.Println(cast.ToInt(v))

	v2, err := redis.Get("test2")
	log.Println(v2, err)
}

func Hget(redis *zredis.Redis) {
	r, err := redis.Hsetnx("testkey", "hashkey1", "11 1")
	log.Println(r, err)

	r2, err := redis.Hget("testkey", "hashkey1")
	log.Println(r2, err)
}

func redisLock(redis *zredis.Redis) bool {
	redisLock := zredis.NewRedisLock(redis, "redislockkey")
	//redisLock.SetExpire(10)
	if ok, err := redisLock.Acquire(); !ok || err != nil {
		//return nil, errors.New("lock failed")
		logx.Errorf("lock failed, ok=%v, err=%v", ok, err)
		return false
	}
	defer func() {
		recover()
		redisLock.Release()
	}()

	logx.Info("lock success")
	return true
}
