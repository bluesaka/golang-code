package dao

import (
	"my-gin/demo/model"
	"my-gin/demo/utils"
)

func (d *Dao) GetUserInfo(params model.User) (user model.User, err error) {
	if err = d.db.Where(&params).First(&user).Error; err != nil {
		utils.Log.Error("[dao/user]GetUserInfo, get error:", err)
		return
	}
	return
}
