package stopwatch4go

import (
	"time"
)

// Stopwatch struct
type Stopwatch struct {
	start time.Time
}

// Start func
func (s *Stopwatch) Start() {
	s.start = time.Now()
}

// Stop func
func (s *Stopwatch) Stop() *Watcher {
	return &Watcher{
		startTime: s.start,
		endTime:   time.Now(),
	}
}
