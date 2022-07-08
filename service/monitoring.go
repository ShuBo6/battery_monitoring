package service

import (
	"battery_monitoring/global"
	"battery_monitoring/utils"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"
)

func Monitoring() {
	for {
		if !ISPowerSupply() {
			zap.L().Info("PowerSupply no")
			ac, b := GetACAdapterInfo(), GetBatteryInfo()
			err := genNotify(global.ToEmail, ac, b)
			if err != nil {
				zap.L().Error("genNotify error", zap.Error(err))
			}
		} else {
			zap.L().Info("PowerSupply yes")
		}

		time.Sleep(time.Duration(global.C.System.MonitoringInterval * 1000 * 1000 * 1000))
	}
}
func genNotify(toEmail, ac, b string) (err error) {
	mailContent := `监测到电源适配器停止供电。
电源适配器状态：
%s
电池状态：
%s`

	err = utils.SendMail(strings.Split(toEmail, ","), "电池状态监测", fmt.Sprintf(mailContent, ac, b))
	if err != nil {
		return
	}
	return
}
