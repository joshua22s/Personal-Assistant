package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/user"
	"strings"
	"time"

	testAlarmClock "github.com/joshua22s/Personal-Assistant/TestAlarmClock"
	_ "github.com/mattn/go-sqlite3"
)

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

func getMapsKey() string {
	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare("SELECT value FROM settings WHERE name = 'mapskey'")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var value string
	err = stmt.QueryRow().Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func getDevices() ([]IAlarmClockDevice, []IBlindDevice, []IClimateDevice, []ILightingDevice) {
	var (
		alarmclocks []IAlarmClockDevice
		blinds      []IBlindDevice
		climates    []IClimateDevice
		lightings   []ILightingDevice
		id          int
		name        string
		deviceType  int
		ipaddr      string
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM device")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&id, &name, &deviceType, &ipaddr)
		switch deviceType {
		case 1:
			alarmclocks = append(alarmclocks, testAlarmClock.TestAlarmClock{id, name, ipaddr})
		}
	}

	return alarmclocks, blinds, climates, lightings
}

func getUserMorningTodosForDay(userid int, day time.Weekday) []MorningTodo {
	var (
		morningtodos  []MorningTodo
		id            int
		name          string
		dayofweek     time.Weekday
		durationInSec int
	)
	db := getConnection()
	defer db.Close()
	fmt.Println("id: ", userid)
	fmt.Println("day: ", strings.ToLower(day.String()))
	rows, err := db.Query("SELECT m.id, m.name, m.duration, d.daynumber FROM morningtodo m JOIN dayofweek d ON m.day = d.id WHERE m.userid = ? AND d.name = ?", userid, strings.ToLower(day.String()))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &durationInSec, &dayofweek)
		//		var duration time.Duration = time.Duration(rand.Int31n(durationInSec)) * time.Second
		fmt.Println(durationInSec)
		morningtodos = append(morningtodos, MorningTodo{id, name, time.Duration(durationInSec * 1000000000), dayofweek})
	}
	return morningtodos
}
