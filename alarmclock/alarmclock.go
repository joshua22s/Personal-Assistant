package alarmclock

import (
	//	"errors"

	//	"github.com/joshua22s/calendarsource"
	"github.com/joshua22s/observer"
)

type AlarmClock struct {
	observer.TriggerBase
	settings string
}

func NewAlarmClock(settings string) *AlarmClock {
	a := AlarmClock{settings: settings}
	return &a
}

func (this *AlarmClock) Activate() error {
	//	appointments := calendarsource.GetAppointments(time.Time{}, time.Time{}.Add(time.Hour*24))
	//departureTime := trafficsource.GetTravelTime("Hoogstraat 39 Beringe", "Eindhoven", "driving", time.Now())
	//fmt.Println(departureTime)

	return nil
}

func (this *AlarmClock) alarm() {
	this.AnnounceAll()
}
