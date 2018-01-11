package observer

import (
	"github.com/satori/go.uuid"
)

type TriggerBase struct {
	listeners []Listener
	id        uuid.UUID
}

func (this *TriggerBase) AddListener(listener Listener) {
	if this.listeners == nil || len(this.listeners) == 0 {
		this.listeners = make([]Listener, 0)
	}
	this.listeners = append(this.listeners, listener)
}

func (this *TriggerBase) GetId() uuid.UUID {
	return this.id
}

func (this *TriggerBase) AnnounceAll() {
	for _, listener := range this.listeners {
		listener.Announce(this)
	}
}
