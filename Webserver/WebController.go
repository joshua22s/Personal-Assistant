package webserver

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"time"

	hub "github.com/joshua22s/Personal-Assistant/Hub"
	//	models "github.com/joshua22s/Personal-Assistant/Models"
	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	appointments := calendar.GetAppointments(
		hub.TimeToStartTime(time.Now().Add(time.Hour*24)),
		hub.TimeToEndTime(time.Now().Add(time.Hour*24)))
	sort.Sort(hub.AppointmentsByDate(appointments))
	var model HomeModel
	if len(appointments) > 0 {
		model = HomeModel{AppointmentModel{appointments[0].Title, appointments[0].Description, hub.FormatFullTime(appointments[0].StartTime), hub.FormatFullTime(appointments[0].EndTime), appointments[0].Location}, ""}
	} else {
		model = HomeModel{AppointmentModel{}, ""}
	}
	t, _ := template.ParseFiles("web/index.html")
	if r.Method == http.MethodGet {
		t.Execute(w, model)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		alarmToUse := hub.CalculateWakeUpTime(time.Now().Add(time.Hour * 24))
		model := HomePostModel{model.AppointmentTomorrow,
			hub.FormatTimeHourMinute(alarmToUse.WakeUpTime),
			TravelHomeModel{alarmToUse.Travel.TravelType, hub.FormatTimeHourMinute(alarmToUse.DepartureTime)}}
		t.Execute(w, model)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("web/todos.html")
		if err != nil {
			fmt.Println(err)
		}
		model := MorningTodoModel{hub.GetAllUserMorningTodos(1)}
		t.Execute(w, model)
	} else if r.Method == http.MethodPost {

	}
}

func travelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("web/travels.html")
		if err != nil {
			fmt.Println(err)
		}
		model := TravelModel{hub.GetAllUserTravels(1)}
		t.Execute(w, model)
	}
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/devices.html")
	if err != nil {
		fmt.Println(err)
	}
	model := DeviceModel{hub.GetAlarmClockDevices(), hub.GetLightingDevies(), hub.GetBlindDevices(), hub.GetClimateDevices()}
	if r.Method == http.MethodGet {
		t.Execute(w, model)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		hub.TurnOnHueLight(r.Form["lightID"][0])
		t.Execute(w, model)
	}
}

func alarmHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/alarm.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func StartWebServer() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/todos", todoHandler)
	http.HandleFunc("/travels", travelHandler)
	http.HandleFunc("/devices", deviceHandler)
	http.HandleFunc("/alarms", alarmHandler)
	//	http.HandleFunc("/login", loginHandler)
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)
	http.Handle("/fonts/", fs)
	http.Handle("/img/", fs)
	fmt.Println("Started webserver, listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
