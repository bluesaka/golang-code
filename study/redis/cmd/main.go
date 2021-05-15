package main

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"go-code/study/redis/goredis"
	"go-code/study/redis/redigo"
	"log"
	"reflect"
	"time"
)

func main() {
	//redigo1()
	//go_redis1()
	//rank.SortBucket()
	redigo3()
}

type RR struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func redigo3() {
	redisConn := redigo.GetRedis()
	defer redigo.CloseRedis()

	s := make(map[string]string, 2)
	s["a"] = "a1"
	s["b"] = "b1"

	var r RR
	str := `{"a":"a1", "b":"b1"}`
	jsoniter.Unmarshal([]byte(str), &r)
	log.Printf("%+v\n", r)

	return

	b, _ := jsoniter.Marshal(s)
	_, err := redisConn.Do("SET", "test1", string(b), "EX", 600)
	if err != nil {
		log.Println("redis set error:", err)
	}

	b2, err := redis.Bytes(redisConn.Do("GET", "test1"))
	if err != nil {
		log.Println("redis get error:", err)
	}

	ss := make(map[string]string)
	jsoniter.Unmarshal(b2, &ss)
	log.Printf("%+v\n", ss)

	log.Println(ss["c"])

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
