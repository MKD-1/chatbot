package logger

import (
	"log/slog"
)

type SystemLogger struct {
	logger *slog.Logger
}

func NewSystemLogger(logger *slog.Logger) *SystemLogger {
	return &SystemLogger{
		logger: logger,
	}
}
func (l *SystemLogger) With(args ...any) *SystemLogger {
	return &SystemLogger{
		logger: l.logger.With(args...),
	}
}

func (l *SystemLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SystemLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SystemLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}
func (l *SystemLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
