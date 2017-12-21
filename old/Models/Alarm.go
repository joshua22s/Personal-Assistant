package models

import (
	"time"
)

type Alarm struct {
	WakeUpTime        time.Time
	DepartureTime     time.Time
	Travel            Travel
	MorningTodos      []MorningTodo
	AlarmClockDevices []IAlarmClockDevice
	BlindDevices      []IBlindDevice
	ClimateDevices    []IClimateDevice
	LightingDevices   []ILightingDevice
}
