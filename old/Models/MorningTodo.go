package models

import (
	"time"
)

type MorningTodo struct {
	Id       int
	Name     string
	Duration time.Duration
	Days     []time.Weekday
}
