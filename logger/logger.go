package logger

import (
	"log"
	"os"

	"github.com/eatmoreapple/openwechat"
)

type Logger struct {
	info    *log.Logger
	error   *log.Logger
	warn    *log.Logger
	message *log.Logger
}

func New() *Logger {
	return &Logger{
		info:    log.New(os.Stdout, "\033[32m[INFO]\033[0m ", log.Ldate|log.Ltime),
		warn:    log.New(os.Stdout, "\033[33m[WARN]\033[0m ", log.Ldate|log.Ltime),
		error:   log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.Ldate|log.Ltime|log.Lshortfile),
		message: log.New(os.Stdout, "\033[36m[MESSAGE]\033[0m", log.Ldate|log.Ltime),
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
func (l *Logger) Message(msg *openwechat.Message) {
	l.message.SetPrefix("\033[36m[MESSAGE]\033[0m")
	if msg == nil {
		return
	}
	switch {
	case msg.IsText():
		l.message.SetPrefix("\033[36m[文本]\033[0m")
		l.message.Println(msg.Content)
	case msg.IsPicture():
		l.message.SetPrefix("\033[36m[图片]\033[0m")
		l.message.Println("[图片消息]")
	case msg.IsVoice():
		l.message.SetPrefix("\033[36m[语音]\033[0m")
		l.message.Println("[语音消息]")
	default:
	}
}
