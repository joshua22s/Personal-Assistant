package philipshue

import (
	"fmt"
	//	"fmt"
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
	//	var err error
	b, err := hue.Discover()
	if err != nil {
		log.Fatal(err)
	}
	if !b.IsPaired() {
		// link button must be pressed for non-error response
		if err := b.Pair(); err != nil {
			log.Fatal(err)
		}
	}
	light, err := b.Lights().Get("Nachtlamp")
	if err != nil {
		log.Fatal(err)
	}
	if err := light.On(); err != nil {
		log.Fatal(err)
	}
	err = light.Set(&hue.State{
		TransitionTime: 20,
		Brightness:     255,
		XY:             &[2]float64{0.438746, 0.501925},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (this PhilipsHueController) ToggleLight(name string) {
	//	fmt.Println(this.bridge)
	//	light, err := this.bridge.Lights().Get(name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	if err := light.On(); err != nil {
	//		log.Fatal(err)
	//	}
}
