package loggers

import (
	"time"
	"bytes"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

)

type LogContext struct {
	Created    time.Time
	Log        bytes.Buffer
	Logger     zap.Logger
	StackTrace interface{}
}

var logLevel = map[string]zapcore.Level{
	"DEBUG":         zap.DebugLevel,
	"INFO":          zap.InfoLevel,
	"WARN":          zap.WarnLevel,
	"ERROR":         zap.ErrorLevel,
	"D_PANIC_LEVEL": zap.DPanicLevel,
	"PANIC_LEVEL":   zap.PanicLevel,
	"FATAL_LEVEL":   zap.FatalLevel,
}

var logger *zap.Logger

func Init(){
	var err error
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return logger
}

func getLogLevel(logLevelValue string) zapcore.Level {
	if val, ok := logLevel[logLevelValue]; ok {
		return val
	}

	return logLevel["INFO"]
}