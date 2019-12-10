package stopwatch4go

import (
	"time"
)

// Watcher struct
type Watcher struct {
	startTime time.Time
	endTime   time.Time
}

// StartTime func
func (w *Watcher) StartTime() time.Time {
	return w.startTime
}

// EndTime func
func (w *Watcher) EndTime() time.Time {
	return w.endTime
}

// Elapsed func
func (w *Watcher) Elapsed() time.Duration {
	return w.endTime.Sub(w.startTime)
}
