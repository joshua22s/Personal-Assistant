package personalassistant

type Action interface {
	Activate() error
}
