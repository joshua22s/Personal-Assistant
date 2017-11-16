package philipshue

import (
	"log"

	"gbbr.io/hue"
)

type PhilipsHueController struct {
	Id     int
	Name   string
	bridge *hue.Bridge
}

func (this PhilipsHueController) GetId() int {
	return this.Id
}

func (this PhilipsHueController) GetName() string {
	return this.Name
}

func (this PhilipsHueController) Setup() {
	var err error
	this.bridge, err = hue.Discover()
	if err != nil {
		log.Fatal(err)
	}
	if !this.bridge.IsPaired() {
		if err := this.bridge.Pair(); err != nil {
			log.Fatal(err)
		}
	}
}

func (this PhilipsHueController) ToggleLight(name string) {
	light, err := this.bridge.Lights().Get(name)
	if err != nil {
		log.Fatal(err)
	}
	if err := light.On(); err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	b, err := hue.Discover()
//	if err != nil {
//		log.Fatal(err)
//	}
//	if !b.IsPaired() {
//		// link button must be pressed before calling
//		if err := b.Pair(); err != nil {
//			log.Fatal(err)
//		}
//	}
//	light, err := b.Lights().Get("Desk")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err := light.On(); err != nil {
//		log.Fatal(err)
//	}
//}
