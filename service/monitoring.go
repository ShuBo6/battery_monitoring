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
			ac, b := GetACAdapterInfo(), GetBatteryInfo()
			zap.L().Info("PowerSupply no", zap.String("ac", ac), zap.String("b", b))
			err := genNotify(global.ToEmail, ac, b)
			if err != nil {
				zap.L().Error("genNotify error", zap.Error(err))
			}
		} else {
			zap.L().Debug("PowerSupply yes")
		}

		time.Sleep(time.Duration(global.C.System.MonitoringInterval * 1000 * 1000 * 1000))
	}
}
func genNotify(toEmail, ac, b string) (err error) {
	mailContent := `
	<div style="display: flex;justify-content: center;">
    <div style="padding: 32px;color: #333;max-width: 580px;">
      <div style="margin-top: 60px;font-size: 24px;font-weight: 500;line-height: 34px;">监测到电源适配器停止供电。
      </div>
      <div style="display: flex;justify-content: center;flex-direction: column;align-items: center;margin-top: 40px;">
        <div style="font-size: 16px;line-height: 22px;">
		电源适配器状态：
		<br> %s
		电池状态：<br>
		%s
		</div>
      </div>
          </div>
    </div>
  </div>
`

	err = utils.SendMail(strings.Split(toEmail, ","), "电池状态监测", fmt.Sprintf(mailContent, ac, b))
	if err != nil {
		return
	}
	return
}
