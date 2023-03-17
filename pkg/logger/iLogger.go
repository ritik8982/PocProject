package logger

type loggerMethods interface {
	Info(string)
	Warn(string)
	Error(string)
}
