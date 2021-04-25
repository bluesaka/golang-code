package main

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
	"go-code/study/redis/goredis"
	"go-code/study/redis/rank"
	"go-code/study/redis/redigo"
	"log"
	"reflect"
	"time"
)

func main() {
	//redigo1()
	//go_redis1()
	rank.SortBucket()
}

func redigo1() {
	redisConn := redigo.GetRedis()
	defer redigo.CloseRedis()

	reply, err := redisConn.Do("SET", "key-a", "value a", "EX", 100, "NX")
	if err != nil {
		log.Println("redigo set error:", err)
	}
	// OK with success
	log.Println("redigo set reply:", reply)
	log.Println("redigo set reply:", reflect.TypeOf(reply))

	reply, err = redisConn.Do("GET", "key-a")
	if err != nil {
		log.Println("redigo get error:", err)
	} else {
		log.Println("redigo get reply:", reply)
	}

	reply, err = redis.String(reply, err)
	if err != nil {
		log.Println("redigo get string error:", err)
	} else {
		log.Println("redigo get string reply:", reply)
	}

	log.Printf("redigo stat, %+v\n", redigo.StatRedis())
}

func go_redis1() {
	redisClient := goredis.GetRedis()
	defer goredis.CloseRedis()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := redisClient.SetNX(ctx, "key-b", "value b", 100*time.Second).Err(); err != nil {
		log.Println("goredis setnx error:", err)
	}

	val, err := redisClient.Get(ctx, "key-b").Result()
	if err != nil {
		log.Println("goredis get error:", err)
	} else {
		log.Println("goredis get value:", val)
	}

	val2, err := redisClient.Get(ctx, "key-c").Result()
	if err == redis2.Nil {
		log.Println("goredis key-c not exist")
	} else if err != nil {
		log.Println("goredis get error:", err)
	} else {
		log.Println("goredis get value:", val2)
	}

	log.Printf("goredis stat, %+v\n", goredis.StatRedis())
}
