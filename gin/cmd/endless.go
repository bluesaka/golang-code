/**
gin + endless 平滑重启
github.com/fvbock/endless

结论：重启时未完成的请求会重复执行，PID会变化，不如grace

# 第一次构建项目
go build endless.go

# 运行项目，这时就可以做内容修改了
./endless &

# 请求项目，60s后返回
ps -ef | grep endless
PID: 10372
curl "http://127.0.0.1:6002/api/login?t=30"

# 再次构建项目，这里是新内容
go build endless.go

# 重启，22072为pid
kill -1 10372

# 新API请求
curl "http://127.0.0.1:6002/api/login?t=1"
*/
package main

import (
	"context"
	"flag"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"my-gin/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := flag.String("port", "6002", "port")
	flag.Parse()

	if true {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	routers.SetRouters(r)

	srv := endless.NewServer(":"+*port, r)

	InitLog()

	log.Println("server start on " + *port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown: ", err)
	}

	log.Println("server exit")
}

func InitLog() {
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
