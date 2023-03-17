package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MyLogger struct {
	logger *zap.Logger
}

func GetLogger() *MyLogger {
	return &MyLogger{
		logger: fileLogger("logs.log"),
	}
}

func (mlog *MyLogger) Info(msg string) {
	mlog.logger.Info(msg)
}

func (mlog *MyLogger) Warn(msg string) {
	mlog.logger.Warn(msg)
}

func (mlog *MyLogger) Error(msg string) {
	mlog.logger.Error(msg)
}

func fileLogger(filename string) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
