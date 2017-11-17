package main

import (
	"time"
)

type Travel struct {
	Id         int
	Name       string
	TravelType string
	Days       []time.Weekday
}
