package initial

import (
	"battery_monitoring/global"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitViper() {
	_v := viper.New()
	_v.SetConfigFile(global.ConfPath)
	if err := _v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf(`读取config.yaml文件失败, err: %v`, err))
	}
	err := _v.Unmarshal(&global.C)
	zap.L().Info("global", zap.Any("", global.C))
	if err != nil {
		panic(fmt.Sprintf(`读取config.yaml文件失败, err: %v`, err))
	}
}
