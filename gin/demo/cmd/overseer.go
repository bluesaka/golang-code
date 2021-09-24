/**
overseer平滑重启
github.com/jpillora/overseer

结论：PID不变，也不会重复执行，符合平滑重启特性，也适合supervisor平滑重启

# 第一次构建项目
go build overseer.go

# 运行项目
./overseer &

# 请求项目，60s后返回
curl "http://127.0.0.1:5003/user/api/login?t=60"

# 修改内容，再次构建项目
go build overseer.go

# 发送USR2新新号重启，master进程pid获取可以通过 `./overseer &` 看到，ps -ef|grep overseer有多个pid注意区别
kill -USR2 {pid}}

# 新API请求
curl "http://127.0.0.1:5003/user/api/login?t=1"
*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"log"
	_ "my-gin/demo/boot"
	"my-gin/demo/utils"
	"my-gin/demo/web/routers"
	"net/http"
)

var (
	address = ":5003"
)

func main() {
	overseer.Run(overseer.Config{
		Address: address,
		Program: prog,
		Debug:   true,
	})
}

func prog(state overseer.State) {
	log.Printf("stateID:%s listening on %s", state.ID, address)
	r := gin.Default()
	routers.SetRouters(r)
	//srv := &http.Server{
	//	Addr:           address,
	//	Handler:        r,
	//	ReadTimeout:    20 * time.Second,
	//	WriteTimeout:   30 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}

	utils.InitZapLog()

	http.Serve(state.Listener, r)

	log.Printf("stateID: %s exiting...\n", state.ID)
}