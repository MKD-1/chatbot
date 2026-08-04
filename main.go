package main

import (
	"log/slog"

	"github.com/MKD-1/chatbot/bootstrap"
	"github.com/MKD-1/chatbot/config"
	"github.com/MKD-1/chatbot/logger"
	"github.com/MKD-1/chatbot/prompts"
)

func main() {
	appLog := logger.New(logger.Config{
		Level: slog.LevelDebug,
	})
	cfg, err := config.Load("config/config.json")
	if err != nil {
		appLog.System.Error("配置加载失败", "err", err)
		return
	}
	// promptStore, err := prompts.Load(cfg.PromptPaths)
	_, err = prompts.Load(cfg.PromptPaths)
	if err != nil {
		appLog.System.Error("提示词加载失败", "err", err)
		return
	}

	appLog.System.Info("你好")
	bootstrap.Run(cfg, appLog.System, appLog.MessageEvent)
}
