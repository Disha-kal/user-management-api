package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger() error {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	Log, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

func GetLogger() *zap.Logger {
	return Log
}
