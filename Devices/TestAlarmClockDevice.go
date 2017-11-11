package devices

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//inherits Device.go, implements IAlarmClockDevice.go
type TestAlarmClockDevice struct {
	Device
	ipAddress string
}

func (this *TestAlarmClockDevice) setTime(timeToSend time.Time) {
	jsonData := timeToSend.Format(time.RFC3339)
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("Unable to parse %v to json\n", jsonData)
	}
	resp, err := http.Post(this.ipAddress+"/alarm", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalf("Http post request to %s failed with error %v", this.ipAddress, err)
	}
}
