package logutil

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	ZapLog *zap.Logger
	Logger *zap.SugaredLogger
)

func InitZapLog() {
	encoder := getZapEncoder()
	//writeSyncer := getFileWriter()
	writeSyncer := getLumberJackWriter()

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// zap.AddCaller() 将函数调用信息记录到日志中 {"level":"info","ts":1618567633.593997,"caller":"log/zap.go:72","msg":"go.uber.org/zap info: hello zap"}
	ZapLog = zap.New(core, zap.AddCaller())
	Logger = ZapLog.Sugar()
}

func getZapEncoder() zapcore.Encoder {
	// json格式的日志
	// {"level":"info","ts":1618567471.695061,"caller":"log/zap.go:24","msg":"go.uber.org/zap info: hello zap"}
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getFileWriter() zapcore.WriteSyncer {
	file, err := os.OpenFile("./zap.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	//file, err := os.OpenFile("/data/logs/zap.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

func getLumberJackWriter() zapcore.WriteSyncer {
	//Filename: 日志文件的位置
	//MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
	//MaxBackups：保留旧文件的最大个数
	//MaxAges：保留旧文件的最大天数
	//Compress：是否压缩/归档旧文件
	lumberJackLogger := &lumberjack.Logger{
		Filename: "./gin_lumber.log",
		//Filename:   "/data/logs/gin_lumber.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
