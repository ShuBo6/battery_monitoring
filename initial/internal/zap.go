package internal

import (
	"battery_monitoring/global"
	"time"

	"go.uber.org/zap/zapcore"
)

// GetEncoderConfig 获取zapcore.EncoderConfig
func GetEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.C.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.C.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.C.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.C.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.C.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return config
}

// GetEncoder 获取zapcore.Encoder
func GetEncoder() zapcore.Encoder {
	if global.C.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(GetEncoderConfig())
}

// GetEncoderCore 获取Encoder的zapcore.Core
func GetEncoderCore(writer zapcore.WriteSyncer, level zapcore.Level) (core zapcore.Core) {
	return zapcore.NewCore(GetEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.C.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
