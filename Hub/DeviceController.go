package main

import (
	"time"

	//	testAlarmClock "github.com/joshua22s/Personal-Assistant/TestAlarmClock"
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
	philipshuecontrol := philipshue.PhilipsHueController{Id: 1, Name: "philipshue"}
	philipshuecontrol.Setup()
	philipshuecontrol.ToggleLight(name)

}
