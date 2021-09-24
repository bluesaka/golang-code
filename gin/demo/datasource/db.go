package datasource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"time"
)

var db *gorm.DB


type gormLogger struct{}

func (*gormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		log.Println("sql", v)
	case "log":
		log.Println("log", v)
	}
}

// GetDB get db
func GetDB() *gorm.DB {
	return db
}

// StatsDB db stats
func StatsDB() sql.DBStats {
	return db.DB().Stats()
}

// CloseDb close db
func CloseDb() error {
	return db.DB().Close()
}

func init() {
	var err error
	db, err = gorm.Open("mysql", viper.GetString("mysql.dsn"))
	if err != nil {
		panic(err)
	}

	// 如果设置禁用表名复数形式属性为 true，`User` 的表名将是 `user`
	db.SingularTable(true)

	// 是否启用Logger，显示详细日志
	db.LogMode(viper.GetBool("mysql.debug"))

	// 最大打开的连接数
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_active_conn"))

	// 最大闲置连接数
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conn"))

	// 最大连接时间
	db.DB().SetConnMaxLifetime(viper.GetDuration("mysql.max_life_time") * time.Second)

	// 最大闲置时间
	db.DB().SetConnMaxIdleTime(viper.GetDuration("mysql.max_idle_time") * time.Second)

	// log
	db.SetLogger(&gormLogger{})
}
