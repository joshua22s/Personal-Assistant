package personalassistant

import (
	"fmt"

	"github.com/joshua22s/observer"
	"github.com/satori/go.uuid"
)

type Event struct {
	Name     string
	triggers map[observer.Trigger]bool
	actions  []Action
}

func NewEvent(name string) *Event {
	e := Event{Name: name}
	e.triggers = make(map[observer.Trigger]bool, 0)
	e.actions = make([]Action, 0)
	return &e
}

func (this *Event) GenerateId() uuid.UUID {
	found := false
	succes := false
	var id uuid.UUID
	for !succes {
		id, _ = uuid.NewV4()
		for trigger, _ := range this.triggers {
			if trigger.GetId() == id {
				found = true
				break
			}
		}
		if !found {
			succes = true
		}
	}
	return id

}

func (this *Event) AddAction(action Action) {
	this.actions = append(this.actions, action)
}

func (this *Event) AddTrigger(trigger observer.Trigger) {
	trigger.AddListener(this)
	this.triggers[trigger] = false
}

func (this *Event) Announce(announcer observer.Trigger) {
	fmt.Println("Going to check it all")
	for trigger, _ := range this.triggers {
		if trigger.GetId() == announcer.GetId() {
			this.triggers[trigger] = true
			this.checkDone()
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
	for _, action := range this.actions {
		err := action.Activate()
		if err != nil {
			fmt.Println(err)
		}
	}
	for trigger, _ := range this.triggers {
		this.triggers[trigger] = false
	}
}
