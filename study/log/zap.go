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
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger, logger2, logger3 *zap.Logger
var slogger, slogger2, slogger3 *zap.SugaredLogger

func Zap() {
	initZap()
	logger.Info("go.uber.org/zap info: hello zap")
	logger.Info("go.uber.org/zap info:", zap.String("msg", "hello zap"))
	slogger.Infof("go.uber.org/zap sugar info: %s", "hello zap sugar")
}

func initZap() {
	// zap.NewProduction()  zap.NewDevelopment() zap.NewExample()
	logger, _ = zap.NewProduction()
	slogger = logger.Sugar()
}

func initZapWithCore() {
	encoder := getEncoder()

	file, err := os.Create("./zap.log")
	if err != nil {
		panic(err)
	}
	writeSyncer := zapcore.AddSync(file)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// logger2 = zap.New(core)
	// zap.AddCaller() 将函数调用信息记录到日志中 {"level":"info","ts":1618567633.593997,"caller":"log/zap.go:72","msg":"go.uber.org/zap info: hello zap"}
	logger2 = zap.New(core, zap.AddCaller())
	slogger2 = logger2.Sugar()
}

func initZapWithLumber() {
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	// logger2 = zap.New(core)
	// zap.AddCaller() 将函数调用信息记录到日志中 {"level":"info","ts":1618567633.593997,"caller":"log/zap.go:72","msg":"go.uber.org/zap info: hello zap"}
	logger3 = zap.New(core, zap.AddCaller())
	slogger3 = logger3.Sugar()
}

func getEncoder() zapcore.Encoder {
	// json格式的日志
	// {"level":"info","ts":1618567471.695061,"caller":"log/zap.go:24","msg":"go.uber.org/zap info: hello zap"}
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getEncoder2() zapcore.Encoder {
	// 命令行格式的日志 1.6185674716960819e+09	info	go.uber.org/zap info: hello zap
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

func getEncoder3() zapcore.Encoder {
	// 命令行格式的日志 2021-04-16T18:05:13.149+0800	INFO	go.uber.org/zap info: hello zap
	c := zap.NewProductionEncoderConfig()
	c.EncodeTime = zapcore.ISO8601TimeEncoder
	c.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(c)
}

func Zap2() {
	initZapWithCore()
	logger2.Info("go.uber.org/zap info: hello zap")
	logger2.Info("go.uber.org/zap info:", zap.String("msg", "hello zap"))
	slogger2.Infof("go.uber.org/zap sugar info: %s", "hello zap sugar")
}

func Zap3() {
	initZapWithLumber()
	for i := 0; i < 10; i++ {
		logger3.Info("go.uber.org/zap lumber info: hello zap")
		logger3.Info("go.uber.org/zap lumber info:", zap.String("msg", "hello zap"))
		slogger3.Infof("go.uber.org/zap lumber sugar info: %s", "hello zap sugar")
	}
}

func getLogWriter() zapcore.WriteSyncer {
	//Filename: 日志文件的位置
	//MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
	//MaxBackups：保留旧文件的最大个数
	//MaxAges：保留旧文件的最大天数
	//Compress：是否压缩/归档旧文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./lumber.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
