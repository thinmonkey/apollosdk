package util

import (
	"time"
)

func Min(x, y time.Duration) time.Duration {
	if x < y {
		return x
	}
	return y
}

