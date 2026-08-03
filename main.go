package main

import (
	"github.com/MKD-1/chatbot/bootstrap"
	// "github.com/MKD-1/chatbot/ai"
	"github.com/MKD-1/chatbot/logger"
	"github.com/joho/godotenv"
)

func main() {
	appLog := logger.New()
	if err := godotenv.Load(); err != nil {
		appLog.Errorf("无法读取 .env: %v", err)
		return
	}
	bootstrap.Run()
	// gtp.Test("回复我圆周率，保留小数点后101位")

}
