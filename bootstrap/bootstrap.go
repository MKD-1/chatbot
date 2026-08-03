package bootstrap

import (
	"github.com/MKD-1/chatbot/handlers"
	"github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

const LoginUrlPrefix = "https://login.weixin.qq.com/l/"

func Bot_ScanCallBack(body openwechat.CheckLoginResponse) {
	appLog := logger.New()
	appLog.Infof("扫码回调: %v", string(body))
}

func Bot_LoginCallBack(body openwechat.CheckLoginResponse) {
	appLog := logger.New()
	appLog.Infof("登录回调: %v", string(body))
}

func Bot_UUIDCallBack(uuid string) {
	appLog := logger.New()
	q, err := qrcode.New(LoginUrlPrefix+uuid, qrcode.Low)
	if err != nil {
		appLog.Errorf("Error generating QR code: %v", err)
		return
	}
	appLog.Infof("%v", q.ToString(true))
}

func Run() {
	appLog := logger.New()
	// 以桌面模式创建机器人.
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.ScanCallBack = Bot_ScanCallBack
	bot.LoginCallBack = Bot_LoginCallBack
	// bot.UUIDCallback = Bot_UUIDCallBack
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	bot.MessageHandler = handlers.Handler
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	// 登录
	// appLog.Info("尝试热登录...")
	// err := bot.HotLogin(reloadStorage)
	// if err != nil {
	// 	appLog.Warn("热登录失败，尝试扫码登录...")
	// 	err = bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	// }
	if err := bot.Login(); err != nil {
		appLog.Errorf("登录失败: %v", err)
	} else {
		appLog.Info("登陆成功")
	}
	if err := bot.Block(); err != nil {
		appLog.Errorf("阻塞失败: %v", err)
	}
}
