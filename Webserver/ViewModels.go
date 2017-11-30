package webserver

import (
	models "github.com/joshua22s/Personal-Assistant/Models"
)

type HomeModel struct {
	AppointmentTomorrow AppointmentModel
	WakeUpTime          string
}

type HomePostModel struct {
	AppointmentTomorrow AppointmentModel
	WakeUpTime          string
	Travel              TravelHomeModel
}

type TravelHomeModel struct {
	Type          string
	DepartureTime string
}

type AppointmentModel struct {
	Title       string
	Description string
	StartTime   string
	EndTime     string
	Location    string
}

type MorningTodoModel struct {
	Todos []models.MorningTodo
}

type TravelModel struct {
	Travels []models.Travel
}

type DeviceModel struct {
	AlarmClockDevices []models.IAlarmClockDevice
	LightingDevices   []models.ILightingDevice
	BlindDevices      []models.IBlindDevice
	ClimateDevices    []models.IClimateDevice
}
