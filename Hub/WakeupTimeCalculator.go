package main

import (
	"fmt"
	"time"

	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
	traffic "github.com/joshua22s/Personal-Assistant/trafficsource"
)

func setupWakeUpTimeCalculator() {
	calendar.Start()
	traffic.Start(getMapsKey())
}

func calculateWakeUpTime(day time.Time) time.Time {
	appointments := calendar.GetAppointments(timeToStartTime(day), timeToEndTime(day))
	departureTime := traffic.GetTravelTime("Hoogstraat 39 Beringe", appointments[0].Location, "driving", appointments[0].StartTime)
	todos := getUserMorningTodosForDay(1, day.Weekday())
	wakeupTime := departureTime
	for _, t := range todos {
		fmt.Println(t.Duration)
		wakeupTime = wakeupTime.Add(-t.Duration)
	}
	return wakeupTime
}

func timeToStartTime(toConvert time.Time) time.Time {
	return time.Date(toConvert.Year(), toConvert.Month(), toConvert.Day(), 0, 0, 1, 0, toConvert.Location())
}

func timeToEndTime(toConvert time.Time) time.Time {
	return time.Date(toConvert.Year(), toConvert.Month(), toConvert.Day(), 23, 59, 59, 0, toConvert.Location())
}
