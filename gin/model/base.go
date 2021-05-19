package model

import (
	"github.com/jinzhu/gorm"
	"my-gin/utils"
	"time"
)

type BaseModel struct {
	ID        int            `gorm:"primary_key" json:"id"`
	CreatedAt utils.JsonTime `json:"created_at"`
	UpdatedAt utils.JsonTime `json:"updated_at"`
}

// gorm hook
func (m BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now())
	scope.SetColumn("updated_at", time.Now())
	return nil
}

func (m BaseModel) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", time.Now())
	return nil
}

type Error struct {
	Err error
	Msg string
}
