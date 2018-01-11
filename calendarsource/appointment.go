package calendarsource

import (
	"time"
)

type Appointment struct {
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Location    string
}
