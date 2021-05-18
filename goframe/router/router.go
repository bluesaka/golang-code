package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gproc"
	"my-goframe/app/api"
	"my-goframe/app/service"
	"my-goframe/logutil"
	"time"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Auth,
			//service.Middleware.CORS,
			service.Middleware.MyCORS,
		)
		//group.ALL("/user", api.User)
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.ALL("/find-by-name", api.User.FindByName)
		})

		group.Group("/car", func(group *ghttp.RouterGroup) {
			group.ALL("/info", api.Car.GetCarInfo)
		})
	})

	// 平滑重启
	// @link https://goframe.org/pages/viewpage.action?pageId=1114220
	s.BindHandler("/pid", func(r *ghttp.Request) {
		r.Response.Write(gproc.Pid())
	})

	s.BindHandler("/sleep", func(r *ghttp.Request) {
		pid := gproc.Pid()
		r.Response.Writeln(pid)
		// 使用gf的平滑重启，这里的log会重复执行一次，类似于endless，和grace的平滑重启不太一样
		logutil.Zap.Info("before pid:", pid)

		time.Sleep(60 * time.Second)

		pid = gproc.Pid()
		logutil.Zap.Info("after pid:", pid)
		r.Response.Writeln(gproc.Pid())
		r.Response.Writeln("hello 555")
	})
}
