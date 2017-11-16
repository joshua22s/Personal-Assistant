package main

import (
	"time"
)

type IAlarmClockDevice interface {
	GetId() int
	GetName() string
	SetTime(time time.Time)
}
