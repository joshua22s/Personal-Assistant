package personalassistant

import (
	"github.com/joshua22s/alarmclock"
	"github.com/joshua22s/philipshue"
	//	"github.com/joshua22s/testalarmclock"
	"github.com/joshua22s/timer"
)

var (
	events []Event
)

func Start() {
	events := make([]Event, 0)
	alarmclockEvent := NewEvent("alarmclock")
	alarmclockEvent.AddTrigger(timer.NewTimer("settings/alarmtimer.json", alarmclockEvent.GenerateId()))
	alarm := alarmclock.NewAlarmClock("settings")
	alarmclockEvent.AddAction(alarm)
	alarmclockResultEvent := NewEvent("alarmclockResult")
	alarmclockResultEvent.AddTrigger(alarm)
	alarmclockResultEvent.AddAction(philipshue.NewPhilipsHue("settings/philipshue.json"))
	//	e := NewEvent("alarmtest")
	//	e.AddAction(alarmclock.NewAlarmClock("settings"))
	//	//	e := NewEvent("test")
	//	e.AddTrigger(timer.NewTimer("settings/timer.json", e.GenerateId()))
	//	e.AddAction(testalarmclock.NewTestAlarmClock())
	//	e.AddAction(philipshue.NewPhilipsHue("settings/philipshue.json"))
	events = append(events, *alarmclockEvent)
	events = append(events, *alarmclockResultEvent)
	//	events = append(events, *e)
	for true {

	}
}
