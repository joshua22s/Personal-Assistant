package main

import (
	"fmt"
	"time"

	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
	traffic "github.com/joshua22s/Personal-Assistant/trafficsource"
)

func main() {
	calendar.Start()
	appointments := calendar.GetAppointments(time.Now(), time.Now().Add(time.Hour*24))
	fmt.Println(appointments)
	traffic.Start(getMapsKey())
	fmt.Println(traffic.GetTravelTime("Hoogstraat 39, Beringe", appointments[0].Location, "driving", appointments[0].StartTime))
	startWebServer()
}
