package handlers

import (
	gtp "github.com/MKD-1/chatbot/ai"
	"github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
)

func GroupHandler(msg *openwechat.Message) {
	appLog := logger.New()
	if
	// msg.IsText() ||
	msg.IsPicture() ||
		msg.IsLocation() ||
		msg.IsVoice() ||
		msg.IsFriendAdd() ||
		msg.IsCard() ||
		msg.IsVideo() ||
		msg.IsRecalled() ||
		msg.IsTransferAccounts() ||
		msg.IsSendRedPacket() ||
		msg.IsReceiveRedPacket() ||
		// msg.IsTickled()||
		// msg.IsTickledMe()||
		msg.IsJoinGroup() ||
		msg.IsSendBySelf() {
		return
	}
	if msg.IsRecalled() {
		appLog.Infof("撤回消息: %v", msg.Content)
	} else if msg.IsTickledMe() {
		msg.ReplyText("捅死你喵")
	} else if msg.IsTickled() {

	} else if msg.IsText() {
		msg.ReplyText(gtp.Talk(msg, "prompts\\group_chat\\system.md"))
	}
}
