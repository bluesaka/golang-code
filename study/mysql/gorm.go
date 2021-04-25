/**
@link https://learnku.com/docs/gorm/v2/index/9728
*/
package mysql

import (
	"database/sql"
	time2 "go-code/study/time"

	// 注意引用mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var db *gorm.DB

// 默认表名是 users
type User struct {
	//gorm.Model
	ID       uint   `json:"id" gorm:"type:int;not null;column:id;primaryKey"`
	Name     string `gorm:"type:varchar(255);"`
	Age      uint8
	Birthday time.Time
}

func (u User) TableName() string {
	return "user"
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
	db, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	// 如果设置禁用表名复数形式属性为 true，`User` 的表名将是 `user`
	db.SingularTable(true)

	// 是否启用Logger，显示详细日志
	db.LogMode(true)

	// 最大打开的连接数
	db.DB().SetMaxOpenConns(10)

	// 最大闲置连接数
	db.DB().SetMaxIdleConns(5)

	// 最大连接时间
	db.DB().SetConnMaxLifetime(3600 * time.Second)

	// 最大闲置时间
	db.DB().SetConnMaxIdleTime(1800 * time.Second)
}

func Insert() {
	user := User{Name: "a", Age: 18, Birthday: time.Now()}
	log.Printf("%+v", user)

	//result := db.Debug().Create(&user)
	result := db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}
	//log.Printf("%+v\n", result)
	log.Println(result.RowsAffected)
	log.Printf("%+v", user)
}

func Query() {
	user := User{}

	// SELECT * FROM user ORDER BY id ASC LIMIT 1;
	//db.First(&user)
	//log.Printf("%+v\n", user)
	//
	//// SELECT * FROM user LIMIT 1;
	//// 注意如果user里有相关，如ID=2，那么会带入到where查询条件中
	//db.Take(&user)
	//log.Printf("%+v\n", user)
	//
	//// SELECT * FROM user ORDER BY id DESC LIMIT 1;
	//db.Last(&user)
	//log.Printf("%+v\n", user)

	//if err := db.Table("user").
	//	Where("id = ?", 2).First(&user).Error; err != nil {
	//		log.Println(err)
	//}
	//log.Printf("%+v\n", user)

	if err := db.Where("id = ?", 2).First(&user).Error; err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", user)

	var sum = 0
	if err := db.Table("user").Select("sum(age) as sum").Row().Scan(&sum); err != nil {
		log.Println(err)
	}
	log.Printf("sum is %v\n", sum)
}

func Update() {
	//if err := db.Model(&User{}).Where("Id = ?", 2).Update("name", "update222").Error; err != nil {
	//	log.Println(err)
	//}

	// 会有两个update操作
	user := User{ID: 2}
	if err := db.Model(&user).
		Update(User{Name: "update333", Birthday: time.Now()}).
		Update("age", gorm.Expr("age+?", 66)).Error; err != nil {
		log.Println(err)
	}
}

func Delete() {
	user := User{ID: 3}
	db.Delete(&user)
	log.Printf("%+v\n", user)

	db.Where("name = ?", "a").Delete(user)
}

func Transaction() {
	//tx := db.Begin()
	//tx.Rollback()
	//tx.Commit()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		log.Println(tx.Error)
		return
	}
	if err := tx.Create(&User{Name: "c1", Birthday: time2.GetDatetime("2021-04-05")}).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	if err := tx.Create(&User{Name: "c2", Birthday: time2.GetDatetime("2021-04-05")}).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()
}

func Transaction2() {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&User{Name: "d1", Birthday: time2.GetDatetime("2021-04-05")}).Error; err != nil {
			return err
		}
		if err := tx.Create(&User{Name: "d2", Birthday: time2.GetDatetime("2021-04-05")}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
}
