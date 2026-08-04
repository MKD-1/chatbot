package handlers

import (
	"strings"

	"github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
)

type Handler struct {
	systemLog  *logger.SystemLogger
	messageLog *logger.MessageEventLogger
}

func New(
	systemLog *logger.SystemLogger,
	messageLog *logger.MessageEventLogger,
) *Handler {
	return &Handler{
		systemLog:  systemLog,
		messageLog: messageLog,
	}
}

func (h *Handler) Handle(msg *openwechat.Message) {
	h.systemLog.Info("收到消息")
}
func IsCommand(content string) bool {
	return true
}
func HandleCommand(msg *openwechat.Message) bool {
	content := strings.TrimSpace(msg.Content)
	if strings.HasPrefix(content, "/") {
		parts := strings.Fields(strings.TrimPrefix(content, "/"))
		if len(parts) == 0 {
			msg.ReplyText("空命令。发送\"/help\"以查看可用命令。")
			return true
		}
		command := strings.ToLower(parts[0])
		// args := parts[1:]

		switch command {
		case "help":
			msg.ReplyText("")
		case "tg":
			// msg.ReplyText(gtp.Talk(msg, "prompts\\group_chat\\tg.md"))
		}
		return true
	} else {
		return false
	}
}
