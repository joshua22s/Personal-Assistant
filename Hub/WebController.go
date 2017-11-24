package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"time"

	calendar "github.com/joshua22s/Personal-Assistant/calendarsource"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	appointments := calendar.GetAppointments(
		TimeToStartTime(time.Now().Add(time.Hour*24)),
		TimeToEndTime(time.Now().Add(time.Hour*24)))
	sort.Sort(AppointmentsByDate(appointments))
	var model HomeModel
	if len(appointments) > 0 {
		model = HomeModel{AppointmentModel{appointments[0].Title, appointments[0].Description, FormatFullTime(appointments[0].StartTime), FormatFullTime(appointments[0].EndTime), appointments[0].Location}, ""}
	} else {
		model = HomeModel{AppointmentModel{}, ""}
	}
	t, _ := template.ParseFiles("web/index.html")
	if r.Method == http.MethodGet {
		t.Execute(w, model)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		wakeUp, travelTime, travelType := calculateWakeUpTime(time.Now().Add(time.Hour * 24))
		model := HomePostModel{model.AppointmentTomorrow, FormatTimeHourMinute(wakeUp), TravelHomeModel{travelType, FormatTimeHourMinute(travelTime)}}
		fmt.Println(model)
		t.Execute(w, model)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("web/todos.html")
		if err != nil {
			fmt.Println(err)
		}
		model := MorningTodoModel{getAllUserMorningTodos(1)}
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
		model := TravelModel{getAllUserTravels(1)}
		t.Execute(w, model)
	}
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/devices.html")
	if err != nil {
		fmt.Println(err)
	}
	model := DeviceModel{alarmClockDevices, lightingDevices, blindDevices, climateDevices}
	if r.Method == http.MethodGet {
		for _, a := range alarmClockDevices {
			fmt.Println("foud")
			fmt.Println(a.GetName())
		}
		fmt.Println(alarmClockDevices)
		t.Execute(w, model)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		fmt.Println("light:")
		fmt.Println(r.Form["lightID"])
		turnOnHueLight(r.Form["lightID"][0])
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

func startWebServer() {
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
