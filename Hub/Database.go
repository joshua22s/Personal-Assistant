package main

import (
	"database/sql"
	"log"
	"os/user"

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
