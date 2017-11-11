package devices

import (
	"time"
)

type IAlarmClockDevice interface {
	func setTime(time time.Time)
}
