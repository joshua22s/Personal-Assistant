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
			if time.Now().Hour() == 11 && time.Now().Minute() == 06 {
				//				alarm := CalculateWakeUpTime(time.Now().Add(time.Hour * 24))
				alarm := CalculateWakeUpTime(time.Now())
				for _, l := range lightingDevices {
					fmt.Println(l.GetName(), " jaaa hier wel?")
				}
				fmt.Println(alarm)
				if time.Now().After(alarm.WakeUpTime.Add(-time.Hour * 2)) {
					//start alarm timer now
					fmt.Println("starting timer now")
					startWakeUpTimer(alarm)
				} else {
					//start 2 hour before timer
					fmt.Println("starting 2 hours before timer")
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
	fmt.Println("wakeuptimer starting")
	timer := time.NewTimer(alarm.WakeUpTime.Sub(time.Now()))
	go func() {
		<-timer.C
		//add triggers for devices
		for _, a := range alarm.AlarmClockDevices {
			a.SetTime(time.Now())
		}
		fmt.Println(alarm)
		for _, l := range alarm.LightingDevices {
			fmt.Println(l.GetName())
			l.Setup()
			l.TurnOn()
		}
		//TODO add implementation for other devices
		fmt.Println("wakeup")
	}()
}
