package pause

import (
	"sync"
	"time"
)

var mu sync.RWMutex
var pauseUntil time.Time

func SetPauseUntil(t time.Time) {
	mu.Lock()
	defer mu.Unlock()
	pauseUntil = t
}

func GetPauseUntil() time.Time {
	mu.RLock()
	defer mu.RUnlock()
	return pauseUntil
}
