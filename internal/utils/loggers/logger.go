package loggers

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

)

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
	fmt.Println("logeeee eeeeeee")
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