package service

import (
	"battery_monitoring/global"
	"battery_monitoring/utils"
	"go.uber.org/zap"
	"html/template"
	"strings"
	"time"
)

func Monitoring() {
	for {
		if !IsCharging() {
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

var mailContent = `
	<div style="display: flex;justify-content: center;">
    <div style="padding: 32px;color: #333;max-width: 580px;">
      <div style="margin-top: 60px;font-size: 24px;font-weight: 500;line-height: 34px;">监测到电源适配器停止供电。
      </div>
      <div style="display: flex;justify-content: center;flex-direction: column;align-items: center;margin-top: 40px;">
        <div style="font-size: 16px;line-height: 22px;">
		电源适配器状态：
		{{ range .AcAdapterLine}}<br> {{ . }} {{end}}
		<br> 电池状态：
		{{ range .BatteryLine}}<br> {{ . }} {{end}}
		</div>
      </div>
          </div>
    </div>
  </div>
`

func genNotify(toEmail, ac, b string) (err error) {
	obj := struct {
		AcAdapterLine []string
		BatteryLine   []string
	}{
		AcAdapterLine: strings.Split(ac, "\n"),
		BatteryLine:   strings.Split(b, "\n"),
	}
	var content string
	content, err = renderTemplate(mailContent, obj)
	if err != nil {
		return err
	}
	err = utils.SendMail(strings.Split(toEmail, ","), "电池状态监测", content)
	if err != nil {
		return
	}
	return
}
func renderTemplate(tpl string, obj interface{}) (ret string, err error) {
	var t *template.Template
	t, err = template.New("renderTemplate").Parse(tpl)
	if err != nil {
		zap.L().Error("renderTemplate err", zap.Error(err))
		return
	}
	b := strings.Builder{}
	err = t.Execute(&b, obj)
	if err != nil {
		zap.L().Error("renderTemplate err", zap.Error(err))
		return
	}
	ret = b.String()
	return
}
