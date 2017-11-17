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
		dayofweek     []time.Weekday
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
		morningtodos = append(morningtodos, MorningTodo{Id: id, Name: name, Duration: time.Duration(durationInSec * 1000000000)})
	}
	for _, m := range morningtodos {
		m.Days = getDaysForMorningTodo(m.Id)
	}
	return morningtodos
}

func getAllUserMorningTodos(userid int) []MorningTodo {
	var (
		morningtodos      []MorningTodo
		finalMorningTodos []MorningTodo
		id                int
		name              string
		durationInSec     int
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT id, name, duration FROM morningtodo WHERE userid = ?", userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &durationInSec)
		fmt.Println(durationInSec)
		morningtodos = append(morningtodos, MorningTodo{Id: id, Name: name, Duration: time.Duration(durationInSec * 1000000000)})
	}
	for _, m := range morningtodos {
		finalMorningTodos = append(finalMorningTodos, MorningTodo{Id: id, Name: name, Duration: time.Duration(durationInSec * 1000000000), Days: getDaysForMorningTodo(m.Id)})
		fmt.Println("m")
		fmt.Println(m)
	}
	fmt.Println("morningtodooos")
	fmt.Println(morningtodos)
	for _, d := range finalMorningTodos[0].Days {
		fmt.Println(d)
	}
	return finalMorningTodos
}

func getDaysForMorningTodo(morningTodoId int) []time.Weekday {
	var (
		weekDays  []time.Weekday
		dayofweek time.Weekday
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT dayNumber FROM dayofweek WHERE dayNumber IN (SELECT dayId FROM MorningTodo_Day WHERE morningTodoId = ?)", morningTodoId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dayofweek)
		fmt.Println(dayofweek)
		weekDays = append(weekDays, dayofweek)
	}
	fmt.Println(weekDays)
	return weekDays
}
