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
	id        int
	name      string
	ipAddress string
}

func NewTestAlarmClock(id int, name string, ipaddress string) *TestAlarmClock {
	return &TestAlarmClock{id: id, name: name, ipAddress: ipaddress}
}

func (this TestAlarmClock) GetId() int {
	return this.id
}

func (this TestAlarmClock) GetName() string {
	return this.name
}

func (this TestAlarmClock) SetTime(timeToSend time.Time) {
	fmt.Println("Time to start:", timeToSend)
	jsonData := timeToSend.Format(time.RFC3339)
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("Unable to parse %v to json\n", jsonData)
	}
	_, err = http.Post(this.ipAddress+"/alarm", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalf("Http post request to %s failed with error %v", this.ipAddress, err)
	}
}
