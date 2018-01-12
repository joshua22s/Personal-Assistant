package personalassistant

import (
	"github.com/joshua22s/alarmclock"
	"github.com/joshua22s/philipshue"
	"github.com/joshua22s/testalarmclock"
	"github.com/joshua22s/timer"
)

var (
	events []Event
)

func Start() {
	events := make([]Event, 0)
	e := NewEvent("alarmtest")
	e.AddAction(alarmclock.NewAlarmClock("settings"))
	//	e := NewEvent("test")
	e.AddTrigger(timer.NewTimer("", e.GenerateId()))
	e.AddAction(testalarmclock.NewTestAlarmClock())
	e.AddAction(philipshue.NewPhilipsHue(""))
	events = append(events, *e)
	for true {

	}
}
