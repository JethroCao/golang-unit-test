package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var LogFile = "./unit-test.log"
var LogLevel = "debug"

func init() {
	initLogger()
}

var Logger *zap.Logger

/**
初始化logger日志对象
*/
func initLogger() {
	hook := lumberjack.Logger{
		Filename:   LogFile, // 日志文件路径
		MaxSize:    1024,    // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       //days
		Compress:   true,    // 是否压缩 disabled by default
	}
	writeSyncer := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(
		// 编码器配置
		zapcore.NewJSONEncoder(NewEncoderConfig()),
		// 打印到控制台和文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			writeSyncer),
		// 日志级别
		level,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "demoProject"))
	// 构造日志
	Logger = zap.New(core, caller, development, filed)
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
