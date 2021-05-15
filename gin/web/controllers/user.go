package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"my-gin/logutil"
	"os"
	"time"
)

func UserLogin(g *gin.Context) {
	t := cast.ToInt(g.Query("t"))

	pid := os.Getpid()
	logrus.Infof("before pid: %d", pid)
	logutil.Logger.Infof("before pid: %d", pid)

	// sleep for test
	time.Sleep(time.Duration(t) * time.Second)

	pid = os.Getpid()
	logrus.Infof("after pid: %d", pid)
	logutil.Logger.Infof("after pid: %d", pid)

	JsonReturn(g, 0, "success", map[string]int{
		"code": t + 777,
	})
}
