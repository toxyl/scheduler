package scheduler

import (
	"time"
)

// Run schedules the execution of `fnCycle` at regular intervals.
//
// Start times will be aligned to multiples of the `every` parameter counted from the zero time,
// e.g. `every = 15 * time.Second` will execute on 00:00:15, 00:00:30, 00:00:45, 00:01:00, and so on.
//
// Use the `offset` parameter to shift the start times. Adding `5 * time.Second` offset would, for example,
// yield the start times 00:00:20, 00:00:35, 00:00:50, 00:01:05, and so on.
//
// The schedule will run in its own goroutine, you can use the returned `stop` function to cancel it.
// The cycle function can also cancel the goroutine by returning `true`. In that case the `fnStop` function,
// if non-nil, will be executed.
func Run(every, offset time.Duration, fnCycle func() (stop bool), fnStop func()) (stop func()) {
	var now time.Time
	var next time.Time
	var sleep time.Duration
	chStop := make(chan struct{})
	go func() {
		for {
			if fnCycle() {
				if fnStop != nil {
					fnStop()
				}
				return
			}
			now = time.Now()
			next = now.Round(every).Add(offset)
			sleep = next.Sub(now)
			if sleep < 0 {
				sleep += every
			}
			select {
			case <-time.After(sleep):
			case <-chStop:
				return
			}
		}
	}()
	return func() {
		select {
		case <-chStop:
		default:
			close(chStop)
		}
	}
}
