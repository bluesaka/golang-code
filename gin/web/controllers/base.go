package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

var ()

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

// JsonReturn json response
func JsonReturn(g *gin.Context, code int, msg string, data interface{}) {
	g.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
}
