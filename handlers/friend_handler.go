package handlers

import (
	// "github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
)

func FriendHandler(msg *openwechat.Message) {
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
		msg.IsTickled() ||
		msg.IsTickledMe() ||
		msg.IsJoinGroup() ||
		msg.IsSendBySelf() {
		return
	}
	if msg.IsText() {

	}
}
