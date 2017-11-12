package main

import (
	"fmt"
	"time"

	devices "github.com/joshua22s/Personal-Assistant/Devices"
	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
	traffic "github.com/joshua22s/Personal-Assistant/trafficsource"
)

func main() {
	calendar.Start()
	appointments := calendar.GetAppointments(time.Now(), time.Now().Add(time.Hour*24))
	fmt.Println(appointments)
	traffic.Start(getMapsKey())
	fmt.Println(traffic.GetTravelTime("Hoogstraat 39, Beringe", appointments[0].Location, "driving", appointments[0].StartTime))
	//startWebServer()
	var alarms []devices.IAlarmClockDevice
	alarms = append(alarms, devices.TestAlarmClockDevice{"test", "172.0.0.1:8080/alarm"})
	devices.SetupDevices(alarms, nil, nil, nil)
	for _, d := range devices.AlarmClockDevices {
		fmt.Printf("hub: %v\n", d.GetName())
	}
}
