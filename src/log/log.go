package log

import (
	"errors"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
)

var (
	Zap  *zap.SugaredLogger
	once sync.Once
)

func Init(env string) {
	once.Do(func() {
		switch env {
		case "dev":
			logger, err := zap.NewDevelopment()
			if err != nil {
				panic(err)
			}
			Zap = logger.Sugar()
		case "test":
			core := zapcore.NewCore(encoder(), logWriter(), zapcore.DebugLevel)
			Zap = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
		case "pro":
			core := zapcore.NewCore(encoder(), logWriter(), zapcore.DebugLevel)
			Zap = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
		default:
			panic(errors.New("请输入环境变量"))
		}
	})
}

func logWriter() zapcore.WriteSyncer {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	lumberJackLogger := lumberjack.Logger{
		Filename:   filepath.Dir(ex) + "/logs/log.log", // 日志文件的位置
		MaxSize:    100,                                // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 60,                                 // 保留旧文件的最大个数
		MaxAge:     180,                                // 保留旧文件的最大天数
		Compress:   true,                               // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(&lumberJackLogger)
}

func encoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	return zapcore.NewConsoleEncoder(encoderConfig)
}
