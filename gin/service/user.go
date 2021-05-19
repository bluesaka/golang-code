package service

import (
	"errors"
	"my-gin/defs"
	"my-gin/model"
	utils "my-gin/utils/logger"
)

func (s *Service) GetUserInfo(id int) (ret model.UserInfoResp, err error) {
	user, err := s.dao.GetUserInfo(model.User{
		BaseModel: model.BaseModel{ID: id},
	})
	if err != nil {
		utils.Log.Errorf("[service/user]GetUserInfo, get err: %v", err)
		err = errors.New(defs.ErrUserNotFound)
		return
	}

	//utils.Log.Infof("user: %+v", user)
	ret = model.UserInfoResp{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}
	return
}
