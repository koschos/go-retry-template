package policy

import "github.com/koschos/go-retry-template/operations"

type SimpleRetryPolicy struct {
	MaxAttempts int
}

func (rp SimpleRetryPolicy) CanRetry(rc operations.RetryContext) bool {
	return rc.LastError == nil || rc.RetryCount <= rp.MaxAttempts
}

func (rp SimpleRetryPolicy) Open() operations.RetryContext {
	return operations.RetryContext{RetryCount: 0, LastError: nil}
}

func (rp SimpleRetryPolicy) Close(rc *operations.RetryContext) {
}

func (rp SimpleRetryPolicy) RegisterError(rc *operations.RetryContext, err error) {
	rc.RetryCount++
	rc.LastError = err
}
