package philipshue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gbbr.io/hue"
)

type PhilipsHue struct {
	settings Settings
	bridge   *hue.Bridge
}

type Settings struct {
	Lights []string `json:"lights"`
}

func NewPhilipsHue(settings string) *PhilipsHue {
	p := PhilipsHue{}
	p.settings = p.readSettings(settings)

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
	fmt.Println(this.settings.Lights)
	for _, l := range this.settings.Lights {
		light, err := this.bridge.Lights().Get(l)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := light.On(); err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println("Turn", light.Name, "on")
	}
	return nil
}

func (this *PhilipsHue) readSettings(settingsPath string) Settings {
	settings := Settings{}
	raw, err := ioutil.ReadFile("settings/philipshue.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(raw, &settings)
	if err != nil {
		fmt.Println(err.Error())
	}
	return settings
}
