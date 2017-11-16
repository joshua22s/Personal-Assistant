package devices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//type IAlarmClockDevice interface {
//	getName() string
//	setTime(time time.Time)
//}

//inherits Device.go, implements IAlarmClockDevice.go
type TestAlarmClockDevice struct {
	ID        int
	Name      string
	IpAddress string
}

func (this TestAlarmClockDevice) GetName() string {
	return this.Name
}

func (this TestAlarmClockDevice) SetTime(timeToSend time.Time) {
	fmt.Println("name: ", this.Name, " ip: ", this.IpAddress, " time: ", timeToSend)
	jsonData := timeToSend.Format(time.RFC3339)
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("Unable to parse %v to json\n", jsonData)
	}
	_, err = http.Post(this.IpAddress+"/alarm", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalf("Http post request to %s failed with error %v", this.IpAddress, err)
	}
}
