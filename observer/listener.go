package observer

type Listener interface {
	Announce(trigger Trigger)
}
