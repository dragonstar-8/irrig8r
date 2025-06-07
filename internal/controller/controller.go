package controller

import (
	"irrig8r/internal/scheduler/events"
	"irrig8r/internal/scheduler/pause"
	"time"
)

func AddSchedule(zone, start string, duration int) error {
	return events.Schedule(zone, start, duration)
}

func Pause(minutes string) {
	d, err := time.ParseDuration(minutes + "m")
	if err == nil {
		pause.SetPauseUntil(time.Now().Add(d))
	} else {
		pause.SetPauseUntil(time.Time{}) // clear pause
	}
}
