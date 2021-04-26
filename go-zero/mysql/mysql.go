package mysql

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	"strings"
)

type (
	UserModel struct {
		Conn  sqlx.SqlConn
		Table string
	}
	User struct {
		ID   string `db:"id"`
		Name string `db:"name"`
		Age  int    `db:"age"`
	}
)

type UserModelOptions struct {
	Table string
}
type UserModelOption func(o *UserModelOptions)

func WithTable(t string) UserModelOption {
	return func(o *UserModelOptions) {
		o.Table = t
	}
}

func newUserModelOptions() UserModelOptions {
	return UserModelOptions{
		Table: "user",
	}
}
func NewUserModel(conn sqlx.SqlConn, opts ...UserModelOption) UserModel {
	options := newUserModelOptions()
	for _, o := range opts {
		o(&options)
	}
	return UserModel{
		Conn:  conn,
		Table: options.Table,
	}
}

func (um *UserModel) Insert(user User) (int64, error) {
	insertSql := `insert into ` + um.Table + `(name, age) values(?, ?)`
	res, err := um.Conn.Exec(insertSql, user.Name, user.Age)
	if err != nil {
		logx.Errorf("insert user error, err=%v", err)
		return -1, nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		logx.Errorf("insert user parse id error, err=%v", err)
		return -1, nil
	}
	return id, nil
}

var userBuilderQueryRows = strings.Join(builderx.FieldNames(&User{}), ",")

func (um *UserModel) FindOne(uid int64) (User, error) {
	var user User
	querySql := `select ` + userBuilderQueryRows + ` from ` + um.Table + ` where id=? limit 1`
	if err := um.Conn.QueryRow(&user, querySql, uid); err != nil {
		logx.Errorf("user fine one error, id=%d, err=%s", uid, err.Error())
		return User{}, err
	}
	return user, nil
}
