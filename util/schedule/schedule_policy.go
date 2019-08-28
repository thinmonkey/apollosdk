package schedule

import (
	"time"

	"github.com/thinmonkey/apollosdk/util"
)

type ExponentialSchedulePolicy struct {
	delayTimeLowerBound time.Duration
	delayTimeUpperBound time.Duration
	lastDelayTime       time.Duration
}

func NewExponentialSchedulePolicy(delayTimeLowerBound time.Duration, delayTimeUpperBound time.Duration) ExponentialSchedulePolicy {
	return ExponentialSchedulePolicy{
		delayTimeUpperBound: delayTimeUpperBound,
		delayTimeLowerBound: delayTimeLowerBound,
	}
}

func (exponentialSchedulePolicy *ExponentialSchedulePolicy) Fail() time.Duration {
	delayTime := exponentialSchedulePolicy.lastDelayTime

	if delayTime == 0 {
		delayTime = exponentialSchedulePolicy.delayTimeLowerBound
	} else {
		delayTime = util.Min(exponentialSchedulePolicy.lastDelayTime<<1, exponentialSchedulePolicy.delayTimeUpperBound)
	}

	exponentialSchedulePolicy.lastDelayTime = delayTime
	return delayTime
}

func (exponentialSchedulePolicy *ExponentialSchedulePolicy) Success() {
	exponentialSchedulePolicy.lastDelayTime = 0
}
