# Go Retry

<b>Motivation</b> is to port spring retry template into Go lang.  

Spring retry https://github.com/spring-projects/spring-retry

## Quick start

Let's say you have potentially dangerous peace of code

```go
// Some not safe code
resp, err := http.Get("http://example.com")
```

1. Install retry-template
```go
 go get github.com/koschos/go-retry-template
```

2. Create simple retry template

```go
rt := retry.RetryTemplate {
	RetryPolicy: policy.SimpleRetryPolicy{MaxAttempts: 2},
	BackoffPolicy: backoff.FixedBackoffPolicy{BackoffPeriodMs: 200},
}
```

3. Run your not safe code inside this retry template

```go
// Call not safe function under this retry logic
result, err := rt.Execute(func(rc operations.RetryContext) (interface{}, error) {
    fmt.Println("Called. Retry Attempt: " + strconv.Itoa(rc.RetryCount))
    return http.Get("http://example.com")
})

// Check error
if err != nil {
    fmt.Println("Error: " + err.Error())
} else {
    fmt.Println("OK ")
}
```

output
```bash
Called. Retry Attempt: 0
Called. Retry Attempt: 1
OK
```

## Extend
### Add New Retry Policy

Example of infinite retry policy

```go
type NeverRetryPolicy struct {}

func (rp NeverRetryPolicy) CanRetry(rc operations.RetryContext) bool {
    return false
}

func (rp NeverRetryPolicy) Open() operations.RetryContext {
    return retry.RetryContext{RetryCount: 0, LastError: nil}
}

func (rp NeverRetryPolicy) Close(rc *operations.RetryContext) {
}

func (rp NeverRetryPolicy) RegisterError(rc *operations.RetryContext, err error) {
}
```

### Add New Backoff Policy

```go
type CustomBackoffPolicy struct{}

func (f CustomBackoffPolicy) Start(rc operations.RetryContext) {
	// your logic here
}

func (f CustomBackoffPolicy) BackOff(rc operations.RetryContext) {
	// your logic here, for example, Fibanacci backoff 
}
```