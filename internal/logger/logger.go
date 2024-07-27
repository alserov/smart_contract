package logger

type Logger interface {
	Info(value string)
	Warn(value string)
	Error(value string)
}
