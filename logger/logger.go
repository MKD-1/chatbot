package logger

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
	warn  *log.Logger
}

func New() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "\033[32m[INFO]\033[0m ", log.Ldate|log.Ltime),
		warn:  log.New(os.Stdout, "\033[33m[WARN]\033[0m ", log.Ldate|log.Ltime),
		error: log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...any) {
	l.warn.Println(v...)
}
func (l *Logger) Error(v ...any) {
	l.error.Println(v...)
}
func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...any) {
	l.warn.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...any) {
	l.error.Printf(format, v...)
}
