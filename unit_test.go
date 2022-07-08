package main

import (
	"battery_monitoring/initial"
	"battery_monitoring/service"
	"go.uber.org/zap"
	"testing"
)

func TestGetInfo(t *testing.T) {
	initial.Init()
	ac, b := service.GetACAdapterInfo(), service.GetBatteryInfo()
	zap.L().Info("PowerSupply no", zap.String("ac", ac), zap.String("b", b))
}
