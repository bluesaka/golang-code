package log

import (
	"github.com/tal-tech/go-zero/core/logx"
	"time"
)

func Logx() {
	logx.SetUp(logx.LogConf{
		Mode: "file",
		Path: "/tmp/logs/logx-test",
		//Level: "info",
	})
	//logx.Disable()
	//logx.SetLevel(logx.ErrorLevel)
	logx.Infof("/tal-tech/go-zero/core/logx info: %s", "hello logx")
	logx.Error("/tal-tech/go-zero/core/logx error")

	// need time to write content to file
	time.Sleep(time.Second)
}
