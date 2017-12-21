package philipshue

import (
	_ "fmt"
	"log"

	"gbbr.io/hue"
)

type PhilipsHueController struct {
	id     int
	name   string
	bridge *hue.Bridge
}

func NewPhilipsHueController(id int, name string) *PhilipsHueController {
	return &PhilipsHueController{id: id, name: name}
}

func (this PhilipsHueController) GetId() int {
	return this.id
}

func (this PhilipsHueController) GetName() string {
	return this.name
}

func (this *PhilipsHueController) Setup() {
	bridge, err := hue.Discover()
	if err != nil {
		log.Fatal(err)
	}
	if !bridge.IsPaired() {
		// link button must be pressed for non-error response
		if err := bridge.Pair(); err != nil {
			log.Fatal(err)
		}
	}
	this.bridge = bridge
}

func (this *PhilipsHueController) TurnOn() {
	light, err := this.bridge.Lights().Get(this.name)
	if err != nil {
		log.Fatal(err)
	}
	if err := light.On(); err != nil {
		log.Fatal(err)
	}
}
