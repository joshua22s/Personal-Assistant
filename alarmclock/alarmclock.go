package alarmclock

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os/user"
	"time"

	"github.com/joshua22s/calendarsource"
	"github.com/joshua22s/observer"
	"github.com/joshua22s/trafficsource"
	"github.com/satori/go.uuid"

	_ "github.com/mattn/go-sqlite3"
)

type AlarmClock struct {
	observer.TriggerBase
	settings string
}

func NewAlarmClock(settings string, id uuid.UUID) *AlarmClock {
	a := AlarmClock{settings: settings}
	a.Id = id
	trafficsource.Start("AIzaSyC8HP5o2pQpZ1D9KGJjUOBJXuw7LPg3VCs")
	return &a
}

func (this *AlarmClock) Activate() error {
	var departureTime time.Time
	found := false
	appointments := calendarsource.GetAppointments(time.Now(), time.Now().Add(time.Hour*24))
	index := 0
	for !found && index < len(appointments) {
		departureTime = trafficsource.GetTravelTime("Hoogstraat 39 Beringe", appointments[index].Location, "driving", appointments[index].StartTime)
		departureTime = departureTime.Add(this.getTodoTime(1, departureTime))
		if departureTime.After(time.Now()) {
			found = true
		}
		index++
	}
	if found {
		this.startAlarmTimer(departureTime)
	} else {
		return errors.New("No upcoming events found")
	}
	return nil
}

func (this *AlarmClock) startAlarmTimer(alarmTime time.Time) {
	fmt.Println("starting alarmtimer:", alarmTime)
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for _ = range ticker.C {
			if time.Now().After(alarmTime) && time.Now().Before(alarmTime.Add(time.Second)) {
				this.alarm()
				return
			}
		}
	}()
}

func (this *AlarmClock) alarm() {
	fmt.Println("announce")
	this.AnnounceAll()
}

func (this *AlarmClock) getTodoTime(userid int, alarmTime time.Time) time.Duration {
	var (
		duration      time.Duration
		durationInSec int
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query(`SELECT m.duration 
	FROM morningtodo m 
	JOIN MorningTodo_Day md ON m.id = md.morningTodoId 
	JOIN dayofweek d ON md.dayId = d.id 
	WHERE m.userid = ? 
	AND d.daynumber = ?`, userid, int(alarmTime.Weekday()))
	fmt.Println(int(alarmTime.Weekday()))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&durationInSec)
		if err != nil {
			log.Fatal(err)
		}
		duration += time.Duration(time.Second * time.Duration(durationInSec))
	}
	return duration
}

func getConnection() *sql.DB {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current user: %v", err)
	}
	db, err := sql.Open("sqlite3", usr.HomeDir+"/database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
