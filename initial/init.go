package initial

import "battery_monitoring/router"

func Init() {
	InitEnvVarious()
	InitViper()
	Zap.InitZap()
	InitRouter()

}
func InitService() {
	router.RegisterRouter()
}
