package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
	traffic "github.com/joshua22s/Personal-Assistant/trafficsource"
)

type AppointmentsByDate []calendar.Appointment

func (a AppointmentsByDate) Len() int           { return len(a) }
func (a AppointmentsByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AppointmentsByDate) Less(i, j int) bool { return a[j].StartTime.After(a[i].StartTime) }

func setupWakeUpTimeCalculator() {
	traffic.Start(getMapsKey())
}

func calculateWakeUpTime(day time.Time) (time.Time, time.Time, string) {
	appointments := calendar.GetAppointments(TimeToStartTime(day), TimeToEndTime(day))
	sort.Sort(AppointmentsByDate(appointments))
	travels := getAllUserTravels(1)
	contains := false
	var travelToUse Travel
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
	todos := getUserMorningTodosForDay(1, day.Weekday())
	wakeupTime := departureTime
	fmt.Println("todos")
	for _, t := range todos {
		fmt.Println(t)
		wakeupTime = wakeupTime.Add(-t.Duration)
	}
	fmt.Println("end todos")
	return wakeupTime, departureTime, travelToUse.TravelType
}

func TimeToStartTime(toConvert time.Time) time.Time {
	return time.Date(toConvert.Year(), toConvert.Month(), toConvert.Day(), 0, 0, 1, 0, toConvert.Location())
}

func TimeToEndTime(toConvert time.Time) time.Time {
	return time.Date(toConvert.Year(), toConvert.Month(), toConvert.Day(), 23, 59, 59, 0, toConvert.Location())
}

func FormatTimeHourMinute(toFormat time.Time) string {
	formatted := ""
	if toFormat.Hour() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Hour())
	} else {
		formatted += strconv.Itoa(toFormat.Hour()) + ""
	}
	formatted += ":"
	if toFormat.Minute() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Minute())
	} else {
		formatted += strconv.Itoa(toFormat.Minute())
	}
	return formatted
}

func FormatFullTime(toFormat time.Time) string {
	formatted := ""
	if toFormat.Hour() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Hour())
	} else {
		formatted += strconv.Itoa(toFormat.Hour()) + ""
	}
	formatted += ":"
	if toFormat.Minute() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Minute())
	} else {
		formatted += strconv.Itoa(toFormat.Minute())
	}
	formatted += "-"
	if toFormat.Day() < 10 {
		formatted += "0" + strconv.Itoa(toFormat.Day())
	} else {
		formatted += strconv.Itoa(toFormat.Day())
	}
	formatted += "-" + toFormat.Month().String() + "-" + strconv.Itoa(toFormat.Year())
	return formatted
}
