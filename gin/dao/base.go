package dao

import (
	"github.com/jinzhu/gorm"
	"my-gin/datasource"
)

type Dao struct {
	db *gorm.DB
}

func NewDao() *Dao {
	return &Dao{
		db: datasource.GetDB(),
	}
}
