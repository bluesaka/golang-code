/**
面向接口编程

Python 里面有个概念叫鸭子类型(duck-typing)，就是如果你叫起来像鸭子，走路像鸭子，游泳像鸭子，那么你就是一只鸭子。
这里的叫、走路、游泳就是我们约定的鸭子接口，而你如果完整实现了这些接口，我们可以像对待一个鸭子一样对待你。
在之前的例子中，不论是 Mysql 实现还是 Redis 实现，他们都有个共同的功能：用一个 id，查一个数据出来，那么这就是共同的接口。
我们可以约定一个叫 DataSource 的接口，它必须有一个方法叫 GetByID，功能是要接收一个 id，返回数据。
*/

package main

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"time"
)

type App struct {
	ds DataSource
}

// NewApp 初始化App，但此处MySQL和Redis的addr格式不同，因此还需优化
func NewApp(addr string) *App {
	// 使用MySQL的时候
	//return &App{ds: NewMySQL(addr)}

	// 使用Redis的时候
	return &App{ds: NewRedis(addr)}
}

func main() {
	app := NewApp("localhost:6379")
	app.ds.GetByID(1)
}

type DataSource interface {
	GetByID(id int)
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