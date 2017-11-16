package main

import (
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
	device := testAlarmClock.TestAlarmClock{1, "wekker", "http://localhost:8080"}
	alarmClockDevices = append(alarmClockDevices, device)
	alarmClockDevices[0].SetTime(time.Now())
}

func triggerAlarmClock(id int, timeToSet time.Time) {
	for _, alarm := range alarmClockDevices {
		if alarm.GetId() == id {
			alarm.SetTime(timeToSet)
		}
	}
}
