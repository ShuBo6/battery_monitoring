package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os/exec"
	"strings"
)

var (
	CmdBatteryInfo = `/org/freedesktop/UPower/devices/battery_BAT0`
	CmdPowerACInfo = `/org/freedesktop/UPower/devices/line_power_AC`
	CmdIsCharging  = `upower -i /org/freedesktop/UPower/devices/battery_BAT0 | grep state | awk '{print $NF}'`
)

func checkCmd(cmd ...string) bool {
	for _, c := range cmd {
		_, err := exec.LookPath(c)
		if err != nil {
			return false
		}
	}
	return true
}
func baseShellExec(cmd string, env map[string]string, args ...string) (string, error) {
	zap.L().Debug("baseShellExec", zap.String("cmd:", cmd))
	if !checkCmd(cmd) {
		//TODO 这里如果没有upower直接panic掉，后续可以优化下
		panic("upower工具未安装")
		return "", errors.New(fmt.Sprintf("cmd[%s] not found ", cmd))
	}
	c := exec.Command(cmd, args...)
	//fmt.Println(c)
	//path, _ := os.Getwd()
	c.Dir = "/tmp"
	for k, v := range env {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}
	output, err := c.CombinedOutput()
	if err != nil {
		zap.L().Error("baseShellExec", zap.Error(err))
		return string(output), err
	}
	zap.L().Debug("baseShellExec", zap.String("cmd output", string(output)))
	return string(output), nil
}
func ExecShell(cmd string, env map[string]string, args ...string) (string, error) {
	return baseShellExec(cmd, env, args...)
}
func IsCharging() bool {
	output, err := ExecShell("bash", nil, "-c", CmdIsCharging)
	if err != nil {
		zap.L().Error("IsCharging failed", zap.Error(err))
		return false
	}
	zap.L().Debug("IsCharging output:" + output)
	//这里经过测试打印出来的是yes\n懒得去处理，直接contains就给过
	return !strings.Contains(output, "discharging")
}
func GetACAdapterInfo() string {
	output, err := ExecShell("upower", nil, "-i", CmdPowerACInfo)
	if err != nil {
		zap.L().Error("GetACAdapterInfo failed", zap.Error(err))
		return err.Error()
	}
	return output
}
func GetBatteryInfo() string {
	output, err := ExecShell("upower", nil, "-i", CmdBatteryInfo)
	if err != nil {
		zap.L().Error("GetBatteryInfo failed", zap.Error(err))
		return err.Error()
	}
	return output
}
