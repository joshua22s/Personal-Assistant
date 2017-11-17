package main

import (
	"time"
)

type MorningTodo struct {
	Id       int
	Name     string
	Duration time.Duration
	Day      time.Weekday
}
