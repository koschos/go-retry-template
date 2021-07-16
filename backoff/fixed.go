package backoff

import (
	"go-retry-template/operations"
	"time"
)

type FixedBackoffPolicy struct {
	BackoffPeriodMs int64
}

func (f FixedBackoffPolicy) Start(rc operations.RetryContext) {
}

func (f FixedBackoffPolicy) BackOff(rc operations.RetryContext) {
	time.Sleep(time.Duration(f.BackoffPeriodMs) * time.Millisecond)
}
