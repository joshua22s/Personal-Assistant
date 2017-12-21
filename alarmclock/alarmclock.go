package alarmclock

import (
	"errors"

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
	//refresh alarm timer
	return errors.New("Unable to refresh timer")
}

func (this *AlarmClock) alarm() {
	this.AnnounceAll()
}
