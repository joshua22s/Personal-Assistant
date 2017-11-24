package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/user"
	"strings"
	"time"

	philipshue "github.com/joshua22s/Personal-Assistant/PhilipsHue"
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
			alarmclocks = append(alarmclocks, testAlarmClock.NewTestAlarmClock(id, name, ipaddr))
			break
		case 2:
			lightings = append(lightings, philipshue.NewPhilipsHueController(id, name))
			break
		}
	}

	return alarmclocks, blinds, climates, lightings
}

func getUserMorningTodosForDay(userid int, day time.Weekday) []MorningTodo {
	var (
		morningtodos      []MorningTodo
		finalMorningTodos []MorningTodo
		id                int
		name              string
		dayofweek         []time.Weekday
		durationInSec     int
	)
	db := getConnection()
	defer db.Close()
	fmt.Println("tostring", strings.ToLower(day.String()))
	rows, err := db.Query("SELECT m.id, m.name, m.duration, d.daynumber FROM morningtodo m JOIN MorningTodo_Day md ON m.id = md.morningTodoId JOIN dayofweek d ON md.dayId = d.daynumber WHERE m.userid = ? AND d.name = ?", userid, strings.ToLower(day.String()))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &durationInSec, &dayofweek)
		morningtodos = append(morningtodos, MorningTodo{Id: id, Name: name, Duration: time.Duration(durationInSec * 1000000000)})
	}
	for _, m := range morningtodos {
		finalMorningTodos = append(finalMorningTodos, MorningTodo{Id: m.Id, Name: m.Name, Duration: m.Duration, Days: getDaysForMorningTodo(m.Id)})
	}
	return finalMorningTodos
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
		finalMorningTodos = append(finalMorningTodos, MorningTodo{Id: m.Id, Name: m.Name, Duration: m.Duration, Days: getDaysForMorningTodo(m.Id)})
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
		weekDays = append(weekDays, dayofweek)
	}
	return weekDays
}

func getDaysForTravel(travelId int) []time.Weekday {
	var (
		weekDays  []time.Weekday
		dayofweek time.Weekday
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT dayNumber FROM dayofweek WHERE dayNumber IN (SELECT dayId FROM Travel_Day WHERE travelid = ?)", travelId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dayofweek)
		weekDays = append(weekDays, dayofweek)
	}
	return weekDays
}

func getAllUserTravels(userid int) []Travel {
	var (
		travels      []Travel
		id           int
		name         string
		traveltype   string
		finalTravels []Travel
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT t.id, t.name, tt.name FROM Travel t JOIN traveltype tt ON tt.id = t.type WHERE t.user = ?", userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &traveltype)
		if err != nil {
			log.Fatal(err)
		}
		travels = append(travels, Travel{Id: id, Name: name, TravelType: traveltype})
	}
	fmt.Println(travels)
	for _, t := range travels {
		finalTravels = append(finalTravels, Travel{Id: t.Id, Name: t.Name, TravelType: t.TravelType, Days: getDaysForTravel(t.Id)})
	}
	return finalTravels
}
