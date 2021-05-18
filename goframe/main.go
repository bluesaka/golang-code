/**
GoFrame
@link https://goframe.org/pages/viewpage.action?pageId=1114155
*/
package main

import (
	"github.com/gogf/gf/frame/g"
	_ "my-goframe/boot"
	_ "my-goframe/logutil"
	_ "my-goframe/router"
)

func main() {
	//runtime.SetMutexProfileFraction(1) // (非必需)开启对锁调用的跟踪
	//runtime.SetBlockProfileRate(1)     // (非必需)开启对阻塞操作的跟踪

	s := g.Server()

	// @link https://goframe.org/pages/viewpage.action?pageId=1114350 开启pprof性能监控
	// go tool pprof http://localhost:8188/debug/pprof/profile
	//s.EnablePProf()

	// @link https://goframe.org/pages/viewpage.action?pageId=1114220 开启服务debug
	//s.EnableAdmin()

	// config/config.toml server.address
	//s.SetPort(8188)

	s.Run()
}
