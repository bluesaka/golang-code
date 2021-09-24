package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"io"
	"my-gin/demo/defs"
	"my-gin/demo/service"
	"net/http"
	"os"
	"strings"
)

type Controller struct {
	srv *service.Service
}

var ctl Controller

func Init() {
	ctl = Controller{
		srv: service.Init(),
	}
}

// JsonReturn json response
//func JsonReturn(g *gin.Context, code int, msg string, data interface{}) {
//	g.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
//}

// JsonReturn json response
func JsonReturn(g *gin.Context, code string, data interface{}) {
	s := strings.Split(code, "|")
	if len(s) != 2 {
		s[0] = cast.ToString(defs.CommonCode)
		s[1] = code
	}
	g.JSON(http.StatusOK, gin.H{"code": cast.ToInt(s[0]), "msg": s[1], "data": data})
}

func init2() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("gin_logrus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	writers := []io.Writer{
		file,
		//os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	logrus.SetOutput(writer)
	logrus.SetLevel(logrus.InfoLevel)
}
