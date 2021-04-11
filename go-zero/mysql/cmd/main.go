package main

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-code/go-zero/mysql"
	"log"
)

func main() {
	conn := sqlx.NewMysql("root:123456@tcp(localhost:3306)/test")
	userModel := mysql.NewUserModel(conn)

	//insert(userModel)
	//findOne(userModel)
	transaction(userModel)

}

func init() {
	logx.SetLevel(logx.ErrorLevel)
}

func insert(userModel mysql.UserModel) {
	user := mysql.User{
		Name: "aaa",
		Age:  3,
	}
	id, _ := userModel.Insert(user)
	log.Println(id)
}

func findOne(userModel mysql.UserModel) {
	user, _ := userModel.FindOne(7)
	logx.Infof("result: %+v\n", user)
}

func transaction(userModel mysql.UserModel) {
	insertSql := `insert into user(name, age) values (?, ?)`
	err := userModel.Conn.Transact(func(session sqlx.Session) error {
		stmt, err := session.Prepare(insertSql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(1, 1); err != nil {
			logx.Errorf("insert userinfo stmt exec: %s", err)
			return err
		}

		if _, err := session.Exec("insert into user(name, age) values (2, 2)"); err != nil {
			logx.Errorf("session exec: %s", err)
			return err
		}

		return nil
	})

	if err != nil {
		logx.Errorf("transaction error, err=%v", err)
	}
}
