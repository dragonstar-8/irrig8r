package scheduler

import (
	"irrig8r/internal/scheduler/events"
)

func Init() {
	events.Start()
}

func Schedule(zone, start string, duration int) error {
	return events.Schedule(zone, start, duration)
}
