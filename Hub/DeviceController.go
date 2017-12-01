package hub

import (
	"time"

	models "github.com/joshua22s/Personal-Assistant/Models"
	philipshue "github.com/joshua22s/Personal-Assistant/PhilipsHue"
)

var (
	alarmClockDevices []models.IAlarmClockDevice
	lightingDevices   []models.ILightingDevice
	blindDevices      []models.IBlindDevice
	climateDevices    []models.IClimateDevice
)

func InitializeDevices() {
	alarmClockDevices, blindDevices, climateDevices, lightingDevices = GetDevices()
}

func triggerAlarmClock(id int, timeToSet time.Time) {
	for _, alarm := range alarmClockDevices {
		if alarm.GetId() == id {
			alarm.SetTime(timeToSet)
		}
	}
}

func GetAlarmClockDevices() []models.IAlarmClockDevice {
	return alarmClockDevices
}

func GetLightingDevies() []models.ILightingDevice {
	return lightingDevices
}

func GetBlindDevices() []models.IBlindDevice {
	return blindDevices
}

func GetClimateDevices() []models.IClimateDevice {
	return climateDevices
}

func TurnOnHueLight(name string) {
	philipshuecontrol := philipshue.NewPhilipsHueController(1, "Nachtlamp")
	philipshuecontrol.Setup()
	philipshuecontrol.TurnOn()
}
