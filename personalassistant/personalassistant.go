package personalassistant

import (
	"github.com/joshua22s/alarmclock"
	"github.com/joshua22s/philipshue"
	//	"github.com/joshua22s/testalarmclock"
	"github.com/joshua22s/midnighttimer"
	//	"github.com/joshua22s/timer"
)

var (
	events []Event
)

func Start() {
	events := make([]Event, 0)
	//	alarmclockEvent := NewEvent("alarmclock")
	//	alarm := alarmclock.NewAlarmClock("settings", alarmclockEvent.GenerateId())
	//	alarmclockEvent.AddAction(alarm)
	//	alarmclockEvent.AddTrigger(timer.NewTimer("settings/alarmtimer.json", alarmclockResultEvent.GenerateId()))
	//	alarmclockResultEvent := NewEvent("alarmclockResult")
	//	alarmclockResultEvent.AddAction(philipshue.NewPhilipsHue("settings/philipshue.json"))
	//	alarmclockResultEvent.AddTrigger(alarm)
	//	events = append(events, *alarmclockEvent)
	//	events = append(events, *alarmclockResultEvent)
	//	events = append(events, *e)
	alarmStartEvent := NewEvent("alarmStartEvent")
	alarm := alarmclock.NewAlarmClock("settings", alarmStartEvent.GenerateId())
	alarmStartEvent.AddAction(alarm)
	alarmStartEvent.AddTrigger(midnighttimer.NewMidNightTimer(alarmStartEvent.GenerateId())) //TODO add midnight timer

	alarmEvent := NewEvent("alarm")
	alarmEvent.AddAction(philipshue.NewPhilipsHue("settings/philipshue.json"))
	alarmEvent.AddTrigger(alarm)
	events = append(events, *alarmStartEvent)
	events = append(events, *alarmEvent)
	for true {

	}
}
