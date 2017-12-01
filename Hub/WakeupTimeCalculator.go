package hub

import (
	_ "fmt"
	"sort"
	"time"

	models "github.com/joshua22s/Personal-Assistant/Models"
	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
	traffic "github.com/joshua22s/Personal-Assistant/trafficsource"
)

type AppointmentsByDate []calendar.Appointment

func (a AppointmentsByDate) Len() int           { return len(a) }
func (a AppointmentsByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AppointmentsByDate) Less(i, j int) bool { return a[j].StartTime.After(a[i].StartTime) }

func SetupWakeUpTimeCalculator() {
	traffic.Start(getMapsKey())
}

func CalculateWakeUpTime(day time.Time) models.Alarm {
	appointments := calendar.GetAppointments(TimeToStartTime(day), TimeToEndTime(day))
	sort.Sort(AppointmentsByDate(appointments))
	travels := GetAllUserTravels(1)
	contains := false
	var travelToUse models.Travel
	for _, travel := range travels {
		for _, d := range travel.Days {
			if d == day.Weekday() {
				contains = true
				break
			}
		}
		if contains {
			travelToUse = travel
			break
		}
	}
	departureTime := traffic.GetTravelTime("Hoogstraat 39 Beringe", appointments[0].Location, travelToUse.TravelType, appointments[0].StartTime)
	todos := GetUserMorningTodosForDay(1, day.Weekday())
	wakeupTime := departureTime
	for _, t := range todos {
		wakeupTime = wakeupTime.Add(-t.Duration)
	}
	alarmClockDevices, blindDevices, climateDevices, lightingDevices := GetDevicesForDay(day.Weekday())
	alarm := models.Alarm{
		WakeUpTime:        wakeupTime,
		DepartureTime:     departureTime,
		Travel:            travelToUse,
		MorningTodos:      todos,
		AlarmClockDevices: alarmClockDevices,
		BlindDevices:      blindDevices,
		ClimateDevices:    climateDevices,
		LightingDevices:   lightingDevices}
	return alarm
}
