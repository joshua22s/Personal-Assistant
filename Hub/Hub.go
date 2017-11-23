package main

import (
	//	"fmt"
	//	"time"

	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
	traffic "github.com/joshua22s/Personal-Assistant/trafficsource"
)

func main() {
	calendar.Start()
	traffic.Start(getMapsKey())
	startWebServer()
	setupWakeUpTimeCalculator()
	//	fmt.Println("Je moet om ", calculateWakeUpTime(time.Now()), " gaan.")
}
