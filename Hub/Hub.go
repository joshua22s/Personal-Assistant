package main

import (
	//	"fmt"
	//	"time"

	traffic "github.com/joshua22s/Personal-Assistant/trafficsource"
)

func main() {
	traffic.Start(getMapsKey())
	initializeDevices()
	startWebServer()
	setupWakeUpTimeCalculator()
	//	fmt.Println("Je moet om ", calculateWakeUpTime(time.Now()), " gaan.")
}
