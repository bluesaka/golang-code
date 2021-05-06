/**
go.uber.org/zap

zap 提供了两种类型的日志记录器—和 Logger 和 Sugared Logger 。两者之间的区别是：
	- 在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。
	- 在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快 4-10 倍，并且支持结构化和 printf 风格的日志记录。
所以一般场景下我们使用 Sugared Logger 就足够了。

日志切割归档
go get -u github.com/natefinch/lumberjack

*/

package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	ZapLogger *zap.SugaredLogger
)

func init() {
	encoder := getZapEncoder()
	file, err := os.OpenFile("./zap.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writeSyncer := zapcore.AddSync(file)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// logger2 = zap.New(core)
	// zap.AddCaller() 将函数调用信息记录到日志中 {"level":"info","ts":1618567633.593997,"caller":"log/zap.go:72","msg":"go.uber.org/zap info: hello zap"}
	ZapLogger = zap.New(core, zap.AddCaller()).Sugar()
}

func getZapEncoder() zapcore.Encoder {
	// json格式的日志
	// {"level":"info","ts":1618567471.695061,"caller":"log/zap.go:24","msg":"go.uber.org/zap info: hello zap"}
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
