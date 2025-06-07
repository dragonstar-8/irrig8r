package executor

import (
	"fmt"
	"irrig8r/internal/mqtt"
	"irrig8r/internal/scheduler/pause"
	"time"
)

func ActivateZone(zone string, duration int) {
	if time.Now().Before(pause.GetPauseUntil()) {
		fmt.Printf("Skipping %s due to pause\n", zone)
		return
	}
	mqtt.PublishCommand(zone, duration)
}
