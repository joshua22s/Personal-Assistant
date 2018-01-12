package midnighttimer

import (
	//	"fmt"
	"time"

	"github.com/joshua22s/observer"
	"github.com/satori/go.uuid"
)

type MidNightTimer struct {
	observer.TriggerBase
}

func NewMidNightTimer(id uuid.UUID) *MidNightTimer {

	t := MidNightTimer{}
	t.Id = id
	t.Start()
	return &t
}

func (this *MidNightTimer) Start() {
	ticker := time.NewTicker(time.Second)
	go func() {
		this.tick()
		for _ = range ticker.C {
			curTime := time.Now()
			if curTime.Hour() == 23 && curTime.Minute() == 59 && curTime.Second() == 59 {
				this.tick()
			}
		}
	}()
}

func (this *MidNightTimer) tick() {
	this.AnnounceAll()
}
