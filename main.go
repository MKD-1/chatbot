package main

import (
	"github.com/MKD-1/chatbot/bootstrap"
	"github.com/MKD-1/chatbot/config"
	"github.com/MKD-1/chatbot/logger"
)

func main() {
	appLog := logger.New()
	cfg, err := config.Load("config/config.json")
	if err != nil {
		appLog.Errorf("加载配置失败: %v", err)
		return
	}

	bootstrap.Run()
}
