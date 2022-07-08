package router

import (
	"battery_monitoring/global"
	"battery_monitoring/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

var handler = response.Handler{}

func RegisterRouter() {
	global.DefaultRouter.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})
	_ = global.DefaultRouter.Run("0.0.0.0:5080")

}
