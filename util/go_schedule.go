package util

import (
	"time"
)

func ScheduleIntervalExecutor(refreshInterval time.Duration, f func()) {
	go func() {
		t2 := time.NewTimer(refreshInterval)
		//long poll for sync
		for {
			select {
			case <-t2.C:
				f()
				t2.Reset(refreshInterval)
			}
		}
	}()
}
