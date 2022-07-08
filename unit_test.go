package main

import (
	"battery_monitoring/initial"
	"battery_monitoring/service"
	"fmt"
	"testing"
)

func TestGetInfo(t *testing.T) {
	initial.Init()
	fmt.Println(service.IsCharging())
}
