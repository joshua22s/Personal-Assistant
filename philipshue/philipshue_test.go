package philipshue

import (
	"testing"
)

func TestLightsOn(t *testing.T) {
	hue := NewPhilipsHue("settings")
	err := hue.Activate()
	if err != nil {
		t.Error(err)
	}
}
