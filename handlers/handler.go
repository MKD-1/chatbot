package handlers

import (
	// "strings"

	"github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
)

func HandleCommand(msg *openwechat.Message) bool {
	return true
}
func Handler(msg *openwechat.Message) {
	appLog := logger.New()
	sender, _ := msg.Sender()
	appLog.Info(sender.NickName + ": " + msg.Content)
	// if msg.IsText() {
	// 	content := strings.TrimSpace(msg.Content)
	// 	if strings.HasPrefix(content, "/") {
	// 		parts := strings.Fields(strings.TrimPrefix(content, "/"))
	// 		if len(parts) == 0 {
	// 			msg.ReplyText("空命令。发送\"/help\"以查看可用命令。")
	// 			return
	// 		}
	// 	}
	// }
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
		msg.IsJoinGroup() ||
		msg.IsSendBySelf() {
		return
	}

	if msg.IsSendByGroup() {
		GroupHandler(msg)
	}

	if msg.IsSendByFriend() {
		FriendHandler(msg)
	}

}
