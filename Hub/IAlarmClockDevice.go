package main

import (
	"time"
)

type IAlarmClockDevice interface {
	GetName() string
	SetTime(time time.Time)
}
