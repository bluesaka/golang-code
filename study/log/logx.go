package log

import "github.com/tal-tech/go-zero/core/logx"

func Logx() {
	logx.Infof("/tal-tech/go-zero/core/logx info: %s", "hello logx")
	str := `<xml>
				<Data>hello xml</Data>
			</xml>`
	logx.Info(str)
	logx.Infof("xml is %v", str)
}
