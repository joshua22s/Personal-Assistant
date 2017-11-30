package hub

import (
	"fmt"
	"time"

	models "github.com/joshua22s/Personal-Assistant/Models"
)

func StartAlarmTimer() {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			if time.Now().Hour() == 19 && time.Now().Minute() == 51 {
				alarm := CalculateWakeUpTime(time.Now().Add(time.Hour * 24))
				if time.Now().After(alarm.WakeUpTime.Add(-time.Hour * 2)) {
					//start alarm timer now
					startWakeUpTimer(alarm)
				} else {
					//start 2 hour before timer
					startRecalculateTimer(alarm.WakeUpTime.Add(-time.Hour*2), alarm)
				}
			}
		}
	}()
}

func startRecalculateTimer(recalculateTime time.Time, alarm models.Alarm) {
	timer := time.NewTimer(recalculateTime.Sub(time.Now()))
	go func() {
		<-timer.C
		newAlarm := CalculateWakeUpTime(time.Now())
		startWakeUpTimer(newAlarm)
	}()
}

func startWakeUpTimer(alarm models.Alarm) {
	timer := time.NewTimer(alarm.WakeUpTime.Sub(time.Now()))
	go func() {
		<-timer.C
		//add triggers for devices
		fmt.Println("wakeup")
	}()
}
