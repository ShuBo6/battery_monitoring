package main

import (
	"battery_monitoring/initial"
	"battery_monitoring/service"
)

func main() {
	initial.Init()
	go service.Monitoring()
	initial.InitService()
}
