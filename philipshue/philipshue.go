package philipshue

import (
	"fmt"
	"log"

	"gbbr.io/hue"
)

type PhilipsHue struct {
	settings string
	bridge   *hue.Bridge
}

func NewPhilipsHue(settings string) *PhilipsHue {
	p := PhilipsHue{}
	p.settings = settings
	p.setup()
	return &p
}

func (this *PhilipsHue) setup() {
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

func (this *PhilipsHue) Activate() error {
	light, err := this.bridge.Lights().Get("Nachtlamp")
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err := light.On(); err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("TURN LIGHTS ON")

	return nil
	//	if false {
	//		return errors.New("Unable to turn on lights")
	//	}
	//	return nil
}
