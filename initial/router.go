package initial

import (
	"battery_monitoring/global"
	"battery_monitoring/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	global.DefaultRouter = gin.Default()
	global.DefaultRouter.Use(router.ZapLogger())
}
