package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"my-gin/demo/defs"
	"my-gin/demo/model"
	"my-gin/demo/utils"
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

func (s *Service) HttpGet() (ret model.HttpResp, err error) {
	params := gin.H{
		"name":  "test",
		"value": "test value",
	}
	resp, err := utils.HttpGetWithParam("http://localhost:8885/huya/sms", params)

	var httpGetResp model.HttpResp
	err = jsoniter.Unmarshal(resp, &httpGetResp)
	if err != nil {
		utils.Log.Error("[service/user]HttpGet, unmarshal error:", err)
	}

	utils.Log.Infof("httpGet resp: %+v", httpGetResp)
	utils.Log.Infof("httpGet resp total: %v", httpGetResp.Data.Total)
	return
}

func (s *Service) HttpPost() (ret model.HttpResp, err error) {
	params := gin.H{
		"name":  "test",
		"value": "test value",
	}
	//resp, err := utils.HttpPost("http://localhost:8885/huya/notify", params)

	json, err := jsoniter.Marshal(params)
	if err != nil {
		utils.Log.Error("[service/user]HttpPost, marshal error:", err)
	}
	resp, err := utils.HttpPostJson("http://localhost:8885/huya/notify", json)

	var httpResp model.HttpResp
	err = jsoniter.Unmarshal(resp, &httpResp)
	if err != nil {
		utils.Log.Error("[service/user]HttpPost, unmarshal error:", err)
	}

	utils.Log.Infof("HttpPost resp: %+v", httpResp)
	utils.Log.Infof("HttpPost resp total: %v", httpResp.Data.Total)
	return
}
