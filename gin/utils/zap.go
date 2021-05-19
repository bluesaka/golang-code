package utils

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	z   *zap.Logger
	Log *zap.SugaredLogger
)

func InitZapLog() {
	encoder := getZapEncoder()
	//writeSyncer := getFileWriter()
	writeSyncer := getLumberJackWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	// zap.AddCaller() 将函数调用信息记录到日志中 {"level":"info","ts":1618567633.593997,"caller":"log/zap.go:72","msg":"go.uber.org/zap info: hello zap"}
	z = zap.New(core, zap.AddCaller())
	Log = z.Sugar()
}

func getZapEncoder() zapcore.Encoder {
	// json格式的日志
	// {"level":"info","ts":1618567471.695061,"caller":"log/zap.go:24","msg":"go.uber.org/zap info: hello zap"}
	cfg := zap.NewProductionEncoderConfig()

	// {"ts":"2021-05-19T10:24:37+08:00"}
	//cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	// {"ts":"2021-05-19 10:24:37"}
	cfg.EncodeTime = MyZapTimeEncoder
	return zapcore.NewJSONEncoder(cfg)
}

func MyZapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
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
		Filename: viper.GetString("log.path"),
		//Filename:   "/data/logs/gin_lumber.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
