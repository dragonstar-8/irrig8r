package events

import (
	"fmt"
	"irrig8r/internal/executor"
	"strings"

	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func Start() {
	c = cron.New()
	c.Start()
}

func Schedule(zone, start string, duration int) error {
	parts := strings.Split(start, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid time format")
	}
	hour := parts[0]
	minute := parts[1]
	spec := fmt.Sprintf("%s %s * * *", minute, hour)

	_, err := c.AddFunc(spec, func() {
		executor.ActivateZone(zone, duration)
	})
	return err
}
