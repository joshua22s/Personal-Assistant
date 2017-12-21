package observer

type Trigger interface {
	AddListener(listener Listener)
}
