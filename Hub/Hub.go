package main

import (
	"fmt"
	"time"

	calendar "github.com/joshua22s/Personal-Assistant/CalendarSource"
)

func main() {
	calendar.Start()
	fmt.Println(calendar.GetAppointments(time.Now(), time.Now().Add(time.Hour*24)))
}
