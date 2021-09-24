/**
gin + grace 平滑重启
github.com/facebookgo/grace

结论：
重启时未完成的请求PID不会变，还是重启前的PID，也不会重复执行。新的请求PID会变，不适用于supervisor平滑重启

// 构建项目
go build grace.go

// 运行项目
./grace &

// 请求服务，等待60s
curl "http://localhost:5001/sleep?duration=60s"

// 修改内容
blah blah

// 重新构建
go build grace.go

// 重启服务
kill -USR2 `ps -ef | grep grace | grep -v grep | awk '{print $2}'`
*/
package main

import (
	"flag"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/sirupsen/logrus"
	"io"
	_ "my-gin/demo/boot"
	"my-gin/demo/utils"
	"my-gin/demo/web/routers"
	"net/http"
	"os"
	"time"
)

var (
	zkReporter reporter.Reporter
)

func main() {
	port := flag.String("port", "6003", "port")
	flag.Parse()

	r := gin.Default()
	if true {
		gin.SetMode(gin.ReleaseMode)
	}

	// zipkin
	if err := initZipkin(r); err != nil {
		panic(err)
	}
	defer zkReporter.Close()

	routers.SetRouters(r)

	srv := &http.Server{
		Addr:           ":" + *port,
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	initLog()
	utils.InitZapLog()

	if err := gracehttp.Serve(srv); err != nil {
		panic(fmt.Sprintf("listen error: %s\n", err))
	}
}

func initZipkin(engine *gin.Engine) error {
	zkReporter = zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint("test-service", "127.0.0.1:6003")
	if err != nil {
		return err
	}

	nativeTracer, err := zipkin.NewTracer(zkReporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return err
	}
	tracer := zkOt.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(tracer)

	engine.Use(func(c *gin.Context) {
		span := tracer.StartSpan(c.FullPath())
		defer span.Finish()
		c.Next()
	})
	return nil
}

func initLog() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("gin_logrus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	//file, err := os.OpenFile("/data/logs/gin_logrus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
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
