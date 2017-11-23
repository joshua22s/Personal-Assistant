package main

type HomeModel struct {
	AppointmentTomorrow AppointmentModel
	WakeUpTime          string
}

type HomePostModel struct {
	AppointmentTomorrow AppointmentModel
	WakeUpTime          string
	Travel              TravelHomeModel
}

type TravelHomeModel struct {
	Type          string
	DepartureTime string
}

type AppointmentModel struct {
	Title       string
	Description string
	StartTime   string
	EndTime     string
	Location    string
}

type MorningTodoModel struct {
	Todos []MorningTodo
}

type TravelModel struct {
	Travels []Travel
}
