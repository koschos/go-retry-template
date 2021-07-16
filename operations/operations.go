package operations

type RetryContext struct {
	RetryCount int
	LastError  error
}

type RetryPolicy interface {
	CanRetry(rc RetryContext) bool
	Open() RetryContext
	Close(rc *RetryContext)
	RegisterError(rc *RetryContext, err error)
}

type BackoffPolicy interface {
	Start(rc RetryContext)
	BackOff(rc RetryContext)
}

type RetryOperations interface {
	Execute(callback func(rc RetryContext) (interface{}, error)) (interface{}, error)
}
