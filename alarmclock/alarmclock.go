package alarmclock

import (
	"errors"
	"fmt"
	"time"

	"github.com/joshua22s/calendarsource"
	"github.com/joshua22s/observer"
	"github.com/joshua22s/trafficsource"
	"github.com/satori/go.uuid"
)

type AlarmClock struct {
	observer.TriggerBase
	settings string
}

func NewAlarmClock(settings string, id uuid.UUID) *AlarmClock {
	a := AlarmClock{settings: settings}
	a.Id = id
	trafficsource.Start("AIzaSyC8HP5o2pQpZ1D9KGJjUOBJXuw7LPg3VCs")
	return &a
}

func (this *AlarmClock) Activate() error {
	var departureTime time.Time
	found := false
	appointments := calendarsource.GetAppointments(time.Now(), time.Now().Add(time.Hour*24))
	index := 0
	for !found && index < len(appointments) {
		departureTime = trafficsource.GetTravelTime("Hoogstraat 39 Beringe", appointments[index].Location, "driving", appointments[index].StartTime)
		fmt.Println(departureTime)
		if departureTime.After(time.Now()) {
			found = true
		}
		index++
	}
	if found {
		this.startAlarmTimer(departureTime)
	} else {
		return errors.New("No upcoming events found")
	}
	return nil
}

func (this *AlarmClock) startAlarmTimer(alarmTime time.Time) {
	fmt.Println("starting alarmtimer:", alarmTime)
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for _ = range ticker.C {
			if time.Now().After(alarmTime) && time.Now().Before(alarmTime.Add(time.Second)) {
				this.alarm()
				return
			}
		}
	}()
}

func (this *AlarmClock) alarm() {
	fmt.Println("announce")
	this.AnnounceAll()
}
