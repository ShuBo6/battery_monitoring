package initial

import (
	"battery_monitoring/global"
	"go.uber.org/zap"
	"os"
)

func InitEnvVarious() {
	// 从环境变量拿要通知的邮箱地址
	toEmail := os.Getenv("TO_EMAILS")
	confPath := os.Getenv("CONF_PATH")
	if toEmail == "" {
		panic("请设置环境变量！！！\n示例：\nTO_EMAILS=xxx@163.com")
	}
	if confPath == "" {
		zap.L().Info("未配置CONF_PATH环境变量，使用默认conf/config.yaml配置文件")
		global.ConfPath = "conf/config.yaml"
	}
	global.ConfPath = confPath
	global.ToEmail = toEmail
}
