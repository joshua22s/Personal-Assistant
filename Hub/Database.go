package main

import (
	"database/sql"
	"log"
	"os/user"

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
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM device")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		//TODO figure out how to read different devices from database
		err = rows.Scan()
	}

	return alarmclocks, blinds, climates, lightings
}
