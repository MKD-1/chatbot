package handlers

import (
	"os"
	"strings"

	gtp "github.com/MKD-1/chatbot/ai"
	"github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
)

func IsCommand(content string) bool {

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
			msg.ReplyText(gtp.Talk(msg, "prompts\\group_chat\\tg.md"))
		}
		return true
	} else {
		return false
	}
}
func Handler(msg *openwechat.Message) {
	appLog := logger.New()
	if msg.IsSendBySelf() {
		return
	}
	sender, _ := msg.Sender()
	appLog.Info(sender.NickName + ": " + msg.Content)
	if msg.IsText() && HandleCommand(msg) {
		return
	}
	if
	// msg.IsText() ||
	msg.IsPicture() ||
		msg.IsLocation() ||
		msg.IsVoice() ||
		msg.IsFriendAdd() ||
		msg.IsCard() ||
		msg.IsVideo() ||
		// msg.IsRecalled() ||
		msg.IsTransferAccounts() ||
		msg.IsSendRedPacket() ||
		msg.IsReceiveRedPacket() ||
		// msg.IsTickled()||
		// msg.IsTickledMe()||
		msg.IsJoinGroup() {
		return
	}

	if msg.IsSendByGroup() {
		GroupHandler(msg)
	}

	if msg.IsSendByFriend() {
		FriendHandler(msg)
	}
}
func loadPrompt(path string) string {
	content, _ := os.ReadFile(path)
	return string(content)
}
