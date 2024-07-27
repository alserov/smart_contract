package logger

import (
	log "log/slog"
	"os"
)

func NewSlog() Logger {
	return &slog{log.New(log.NewJSONHandler(os.Stdout, &log.HandlerOptions{
		Level: log.LevelInfo,
	}))}
}

type slog struct {
	l *log.Logger
}

func (s slog) Info(value string) {
	s.l.Info(value)
}

func (s slog) Warn(value string) {
	s.l.Warn(value)
}

func (s slog) Error(value string) {
	s.l.Error(value)
}
