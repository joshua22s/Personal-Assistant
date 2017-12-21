package personalassistant

import (
	"github.com/joshua22s/testalarmclock"
	"github.com/joshua22s/timer"
)

var (
	events []Event
)

func Start() {
	events := make([]Event, 0)
	e := NewEvent("test")
	e.AddTrigger(timer.NewTimer(""))
	e.AddAction(testalarmclock.NewTestAlarmClock())
	events = append(events, *e)
	for true {
		for _, e := range events {
			e.checkDone()
		}
	}
}
