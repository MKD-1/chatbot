package logger

import (
	"io"
	"log/slog"
	"os"
)

type Config struct {
	SystemWriter io.Writer
	Level        slog.Level
	// LogMessageContent   bool
	// MessagePreviewLimit int
}

type Logger struct {
	System       *SystemLogger
	MessageEvent *MessageEventLogger
}

func New(cfg Config) *Logger {
	if cfg.SystemWriter == nil {
		cfg.SystemWriter = os.Stderr
	}

	// if cfg.MessagePreviewLimit <= 0 {
	// 	cfg.MessagePreviewLimit = 50
	// }
	systemHandler := slog.NewTextHandler(cfg.SystemWriter, &slog.HandlerOptions{
		Level: cfg.Level,
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			switch attr.Key {
			case slog.TimeKey:
				return slog.String(
					slog.TimeKey,
					attr.Value.Time().Format("2006/01/02-15:04:05"),
				)
			default:
				return attr
			}
		},
	})
	// messageEventHandler:=slog.NewJSONHandler()
	return &Logger{
		System: NewSystemLogger(slog.New(systemHandler)),
	}
}
