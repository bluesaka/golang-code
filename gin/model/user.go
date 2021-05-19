package model

import (
	utils "my-gin/utils/time"
)

type User struct {
	BaseModel
	Name string `gorm:"type:varchar(20); not null;" json:"name"`
	Age  int    `gorm:"column:age"`
	//CreatedAt time.Time `gorm:"column:created_at"`
	//CreatedAt utils.JsonTime `gorm:"column:created_at"`
}

type UserInfoReq struct {
	ID int `json:"id"`
}

type UserInfoResp struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	CreatedAt utils.JsonTime `json:"created_at"`
}
