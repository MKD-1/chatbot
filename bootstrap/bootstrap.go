package bootstrap

import (
	"github.com/MKD-1/chatbot/config"
	"github.com/MKD-1/chatbot/handlers"
	"github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
)

const LoginUrlPrefix = "https://login.weixin.qq.com/l/"

func Bot_ScanCallBack(appLog *logger.SystemLogger) func(body openwechat.CheckLoginResponse) {
	return func(body openwechat.CheckLoginResponse) {
		appLog.Info("扫码回调", "response", string(body))
	}
}
func Bot_LoginCallBack(appLog *logger.SystemLogger) func(body openwechat.CheckLoginResponse) {
	return func(body openwechat.CheckLoginResponse) {
		appLog.Info("登录回调", "response", string(body))
	}
}

func Run(cfg config.Config,
	systemLog *logger.SystemLogger,
	messageEventLog *logger.MessageEventLogger) {
	messageHandler := handlers.New(systemLog, messageEventLog)
	// 以桌面模式创建机器人.
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.ScanCallBack = Bot_ScanCallBack(systemLog)
	bot.LoginCallBack = Bot_LoginCallBack(systemLog)
	// bot.UUIDCallback = Bot_UUIDCallBack
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	bot.MessageHandler = messageHandler.Handle
	// 创建热存储容器对象
	// reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	// defer reloadStorage.Close()
	// 登录
	// appLog.Info("尝试热登录...")
	// err := bot.HotLogin(reloadStorage)
	// if err != nil {
	// 	appLog.Warn("热登录失败，尝试扫码登录...")
	// 	err = bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	// }

	if err := bot.Login(); err != nil {
		systemLog.Error("登录失败", "err", err)
	} else {
		systemLog.Info("登录成功")
	}
	if err := bot.Block(); err != nil {
		systemLog.Error("阻塞失败", "err", err)
	}
}
