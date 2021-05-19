package model

import (
	"my-gin/utils"
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

type HttpResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//Data interface{} `json:"data"`
	Data struct {
		List []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}
