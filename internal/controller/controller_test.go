package controller

import (
	"irrig8r/internal/scheduler/pause"
	"testing"
	"time"
)

func TestPause(t *testing.T) {
	Pause("5")
	pauseUntil := pause.GetPauseUntil()
	if time.Until(pauseUntil) < 4*time.Minute {
		t.Errorf("Pause too short: got %v", pauseUntil)
	}
}
