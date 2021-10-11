package main

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"go-code/study/di/wire/internal/cache"
	"go-code/study/di/wire/internal/config"
	"log"
)

type App struct {
	redis *redigo.Pool
}

func NewApp(redis *redigo.Pool) *App {
	return &App{redis: redis}
}

func main() {
	c, _ := config.NewConfig()
	cache.NewRedis(c)

	app, err := InitApp()
	if err != nil {
		panic(err)
	}

	redisConn := app.redis.Get()
	defer redisConn.Close()

	fmt.Println(redisConn.Do("info memory"))

	log.Println("init success")
}
