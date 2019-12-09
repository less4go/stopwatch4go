package stopwatch4go

import (
	"time"
)

// Watch struct
type Watch struct {
	startTime time.Time
	endTime   time.Time
}

// StartTime func
func (w *Watch) StartTime() time.Time {
	return w.startTime
}

// EndTime func
func (w *Watch) EndTime() time.Time {
	return w.endTime
}

// Elapsed func
func (w *Watch) Elapsed() time.Duration {
	return w.endTime.Sub(w.startTime)
}
