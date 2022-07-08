package initial

import (
	"battery_monitoring/global"
	"os"
)

func InitEnvVarious() {
	// 从环境变量拿要通知的邮箱地址
	toEmail := os.Getenv("TO_EMAILS")
	if toEmail == "" {
		panic("请设置环境变量！！！\n示例：\nTO_EMAILS=xxx@163.com")
	}
	global.ToEmail = toEmail
}
