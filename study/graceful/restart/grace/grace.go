/**
facebook grace 平滑重启

测试步骤：

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
kill -USR2 `ps -ef | grep grace | grep -v grep | awk 'print {$2}'`

可以发现，重启时未完成的请求会继续旧服务处理，新来的请求会使用新的服务，实现了平滑重启的功能

 */
package main

import (
	"github.com/facebookgo/grace/gracehttp"
	"go-code/study/log"
	"net/http"
	"os"
	"time"
)

func main() {
	gracehttp.Serve(
		&http.Server{Addr: ":5001", Handler: newGraceHandler()},
		&http.Server{Addr: ":5002", Handler: newGraceHandler()},
	)
}

func newGraceHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		duration, err := time.ParseDuration(r.FormValue("duration"))
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		pid := os.Getpid()
		log.ZapLogger.Info("before pid:", pid)

		time.Sleep(duration)

		pid = os.Getpid()
		log.ZapLogger.Info("after pid:", pid)

		w.Write([]byte("Hello 222"))
	})
	return mux
}
