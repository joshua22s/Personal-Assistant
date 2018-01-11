package observer

import (
	"github.com/satori/go.uuid"
)

type Trigger interface {
	AddListener(listener Listener)
	GetId() uuid.UUID
}
