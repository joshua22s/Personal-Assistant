package main

import (
	"fmt"
	"time"

	testAlarmClock "github.com/joshua22s/Personal-Assistant/TestAlarmClock"
)

var (
	alarmClockDevices []IAlarmClockDevice
	lightingDevices   []ILightingDevice
	blindDevices      []IBlindDevice
	climateDevices    []IClimateDevice
)

func initializeDevices() {
	fmt.Println("device name")

	device := testAlarmClock.TestAlarmClock{"wekker", "http://localhost:8080"}
	fmt.Println(device.GetName())
	alarmClockDevices = append(alarmClockDevices, device)
	for _, d := range alarmClockDevices {
		d.SetTime(time.Now().Add(time.Hour * 1))
	}
	//	alarmClockDevices[0].SetTime(time.Now())
}
