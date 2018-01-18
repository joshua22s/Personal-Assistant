package personalassistant

import (

	//	"github.com/joshua22s/alarmclock"
	//	"github.com/joshua22s/midnighttimer"
	//	"github.com/joshua22s/philipshue"
	"github.com/joshua22s/testalarmclock"
	"github.com/joshua22s/timer"
)

var (
	events []Event
)

func Start() {
	events := make([]Event, 0)

	//	alarmStartEvent := NewEvent("alarmStartEvent")
	//	alarm := alarmclock.NewAlarmClock("settings", alarmStartEvent.GenerateId())
	//	alarmStartEvent.AddAction(alarm)
	//	alarmStartEvent.AddTrigger(midnighttimer.NewMidNightTimer(alarmStartEvent.GenerateId()))

	//	alarmEvent := NewEvent("alarm")
	//	alarmEvent.AddAction(philipshue.NewPhilipsHue("settings/philipshue.json"))
	//	alarmEvent.AddTrigger(alarm)

	randomEvent := NewEvent("randomEvent")
	randomEvent.AddAction(testalarmclock.NewTestAlarmClock())
	//	randomEvent.AddAction(philipshue.NewPhilipsHue("settings/philipshue.json"))
	//	randomEvent.AddTrigger(timer.NewTimer("settings/7sectimer.json", randomEvent.GenerateId()))
	randomEvent.AddTrigger(timer.NewTimer("settings/10sectimer.json", randomEvent.GenerateId()))

	//	events = append(events, *alarmStartEvent)
	//	events = append(events, *alarmEvent)
	events = append(events, *randomEvent)

	for true {

	}
}
