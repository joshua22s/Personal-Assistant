package main

import (
	"time"

	philipshue "github.com/joshua22s/Personal-Assistant/PhilipsHue"
)

var (
	alarmClockDevices []IAlarmClockDevice
	lightingDevices   []ILightingDevice
	blindDevices      []IBlindDevice
	climateDevices    []IClimateDevice
)

func initializeDevices() {
	alarmClockDevices, blindDevices, climateDevices, lightingDevices = getDevices()
}

func triggerAlarmClock(id int, timeToSet time.Time) {
	for _, alarm := range alarmClockDevices {
		if alarm.GetId() == id {
			alarm.SetTime(timeToSet)
		}
	}
}

func turnOnHueLight(name string) {
	philipshuecontrol := philipshue.NewPhilipsHueController(1, "philipshue")
	philipshuecontrol.Setup()
	philipshuecontrol.ToggleLight(name)
}
