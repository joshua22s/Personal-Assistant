package observer

type TriggerBase struct {
	listeners []Listener
}

func (this *TriggerBase) AddListener(listener Listener) {
	if this.listeners == nil || len(this.listeners) == 0 {
		this.listeners = make([]Listener, 0)
	}
	this.listeners = append(this.listeners, listener)
}

func (this *TriggerBase) AnnounceAll() {
	for _, listener := range this.listeners {
		listener.Announce(this)
	}
}
