/**
此例子中没有依赖注入，不过相较于example1，db对象被塞到app对象中，比较安全。
但是在app内部产生了依赖，如果需要更换数据获取方式，如MySQL替换为Redis，需要修改很多地方。
*/

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type App struct {
	db *gorm.DB
}

// NewApp 初始化App
func NewApp() *App {
	return &App{
		db: NewMySQLClient("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local&timeout=5s&readTimeout=10s&writeTimeout=10s"),
	}
}

// NewMySQLClient 初始化MySQL连接
func NewMySQLClient(addr string) *gorm.DB {
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(err)
	}
	return db
}

// GetData 从MySQL获得数据
func (a *App) GetData(id int) {
	// a.db.Find("blah blah ...")
	fmt.Println("GetData from MySQL")
}

func main() {
	app := NewApp()
	app.GetData(1)
}
