package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"my-gin/defs"
	"my-gin/utils"
	"os"
	"time"
)

func UserLogin(g *gin.Context) {
	t := cast.ToInt(g.Query("t"))

	pid := os.Getpid()
	logrus.Infof("before pid: %d", pid)
	utils.Log.Infof("before pid: %d", pid)

	// sleep for test
	time.Sleep(time.Duration(t) * time.Second)

	pid = os.Getpid()
	logrus.Infof("after pid: %d", pid)
	utils.Log.Infof("after pid: %d", pid)

	JsonReturn(g, defs.SuccessCode, map[string]int{
		"code": t + 777,
	})
}

func UserInfo(g *gin.Context) {
	id := g.Query("id")
	if id == "" {
		JsonReturn(g, defs.ErrParam, nil)
		return
	}

	ret, err := ctl.srv.GetUserInfo(cast.ToInt(id))
	if err != nil {
		JsonReturn(g, err.Error(), nil)
		return
	}

	JsonReturn(g, defs.SuccessCode, ret)
}

func HttpGet(g *gin.Context) {
	ret, err := ctl.srv.HttpGet()
	if err != nil {
		JsonReturn(g, err.Error(), nil)
		return
	}

	JsonReturn(g, defs.SuccessCode, ret)
}

func HttpPost(g *gin.Context) {
	ret, err := ctl.srv.HttpPost()
	if err != nil {
		JsonReturn(g, err.Error(), nil)
		return
	}

	JsonReturn(g, defs.SuccessCode, ret)
}
