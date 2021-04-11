package main

import (
	"github.com/tal-tech/go-zero/core/bloom"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"log"
)

func main() {
	bloomFilter()
}

func bloomFilter() {
	store := redis.NewRedis("127.0.0.1:6379", redis.NodeType, "")
	key := "test_bloom_key"
	b := []byte("k1")
	//b = append(b, []byte("k1")...)
	//b = append(b, []byte("k2")...)
	filter := bloom.New(store, key, 1024)
	exist, _ := filter.Exists(b)
	if exist {
		log.Println("exist:", exist)
		return
	}

	if err := filter.Add(b); err != nil {
		log.Println("add err:", err)
		return
	}
	log.Println("add success")

	if err := store.Expire(key, 600); err != nil {
		log.Println("expire err:", err)
		return
	}
	log.Println("expire success")

}
