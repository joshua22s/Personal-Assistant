package personalassistant

import (
	"fmt"

	"github.com/joshua22s/observer"
)

type Event struct {
	Name     string
	triggers map[*observer.Trigger]bool
	actions  []Action
}

func NewEvent(name string) *Event {
	e := Event{Name: name}
	e.triggers = make(map[*observer.Trigger]bool, 0)
	e.actions = make([]Action, 0)
	return &e
}

func (this *Event) AddAction(action Action) {
	this.actions = append(this.actions, action)
}

func (this *Event) AddTrigger(trigger observer.Trigger) {
	trigger.AddListener(this)
	this.triggers[&trigger] = false
}

func (this *Event) Announce(announcer observer.Trigger) {
	fmt.Println(len(this.triggers))
	for trigger, _ := range this.triggers {
		fmt.Println(trigger)
		fmt.Println(&announcer)
		if trigger == &announcer {
			fmt.Println("announce")
			this.triggers[trigger] = true
			return
		}
	}
}

func (this *Event) checkDone() {
	for _, done := range this.triggers {
		if !done {
			return
		}
	}
	fmt.Println("actions")
	fmt.Println(this.actions)
	for _, action := range this.actions {
		action.Activate()
	}

	for trigger, _ := range this.triggers {
		this.triggers[trigger] = false
	}
}
