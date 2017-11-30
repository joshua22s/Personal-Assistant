package hub

import (
	"database/sql"
	_ "fmt"
	"log"
	"os/user"
	"strings"
	"time"

	models "github.com/joshua22s/Personal-Assistant/Models"
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

func GetDevices() ([]models.IAlarmClockDevice, []models.IBlindDevice, []models.IClimateDevice, []models.ILightingDevice) {
	var (
		alarmclocks []models.IAlarmClockDevice
		blinds      []models.IBlindDevice
		climates    []models.IClimateDevice
		lightings   []models.ILightingDevice
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

func GetUserMorningTodosForDay(userid int, day time.Weekday) []models.MorningTodo {
	var (
		morningtodos      []models.MorningTodo
		finalMorningTodos []models.MorningTodo
		id                int
		name              string
		dayofweek         []time.Weekday
		durationInSec     int
	)
	db := getConnection()
	defer db.Close()
	rows, err := db.Query("SELECT m.id, m.name, m.duration, d.daynumber FROM morningtodo m JOIN MorningTodo_Day md ON m.id = md.morningTodoId JOIN dayofweek d ON md.dayId = d.daynumber WHERE m.userid = ? AND d.name = ?", userid, strings.ToLower(day.String()))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &durationInSec, &dayofweek)
		morningtodos = append(morningtodos, models.MorningTodo{Id: id, Name: name, Duration: time.Duration(durationInSec * 1000000000)})
	}
	for _, m := range morningtodos {
		finalMorningTodos = append(finalMorningTodos, models.MorningTodo{Id: m.Id, Name: m.Name, Duration: m.Duration, Days: GetDaysForMorningTodo(m.Id)})
	}
	return finalMorningTodos
}

func GetAllUserMorningTodos(userid int) []models.MorningTodo {
	var (
		morningtodos      []models.MorningTodo
		finalMorningTodos []models.MorningTodo
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
		morningtodos = append(morningtodos, models.MorningTodo{Id: id, Name: name, Duration: time.Duration(durationInSec * 1000000000)})
	}
	for _, m := range morningtodos {
		finalMorningTodos = append(finalMorningTodos, models.MorningTodo{Id: m.Id, Name: m.Name, Duration: m.Duration, Days: GetDaysForMorningTodo(m.Id)})
	}
	return finalMorningTodos
}

func GetDaysForMorningTodo(morningTodoId int) []time.Weekday {
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

func GetDaysForTravel(travelId int) []time.Weekday {
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

func GetAllUserTravels(userid int) []models.Travel {
	var (
		travels      []models.Travel
		id           int
		name         string
		traveltype   string
		finalTravels []models.Travel
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
		travels = append(travels, models.Travel{Id: id, Name: name, TravelType: traveltype})
	}
	for _, t := range travels {
		finalTravels = append(finalTravels, models.Travel{Id: t.Id, Name: t.Name, TravelType: t.TravelType, Days: GetDaysForTravel(t.Id)})
	}
	return finalTravels
}
