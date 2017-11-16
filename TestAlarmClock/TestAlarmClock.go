package testalarmclock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TestAlarmClock struct {
	Id        int
	Name      string
	IpAddress string
}

func (this TestAlarmClock) GetId() int {
	return this.Id
}

func (this TestAlarmClock) GetName() string {
	return this.Name
}

func (this TestAlarmClock) SetTime(timeToSend time.Time) {
	fmt.Println("Time to start:", timeToSend)
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
