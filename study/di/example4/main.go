package main

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"time"
)

type DataSource interface {
	GetByID(id int)
}

type App struct {
	ds DataSource
}

// NewApp 初始化App
// 通过DI依赖注入的方式，由外部传入对象来创建模块
func NewApp(ds DataSource) *App {
	return &App{ds: ds}
}

func main() {
	app := NewApp(NewRedis("localhost:6379"))
	app.ds.GetByID(2)
}

// mysql
type mysql struct {
	m *gorm.DB
}

// NewRedis mysql
func NewMySQL(addr string) *mysql {
	return &mysql{m: NewMySQLClient(addr)}
}

// NewMySQLClient 初始化MySQL连接
func NewMySQLClient(addr string) *gorm.DB {
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(err)
	}
	return db
}

func (m *mysql) GetByID(id int) {
	fmt.Println("GetById from MySQL")
}


// redis
type redis struct {
	r *redigo.Pool
}

// NewRedis 初始化redis
func NewRedis(addr string) *redis {
	return &redis{r: NewRedisClient(addr)}
}

// NewRedisClient 初始化Redis连接
func NewRedisClient(addr string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     2,
		MaxActive:   5,
		IdleTimeout: 3600 * time.Second,
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", addr, redigo.DialDatabase(0), redigo.DialPassword(""))
		},
	}
}

func (r *redis) GetByID(id int) {
	fmt.Println("GetById from Redis")
}

