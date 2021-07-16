package examples

import (
	"errors"
	"fmt"
	"math/rand"
	"go-retry-template"
	"go-retry-template/backoff"
	"go-retry-template/operations"
	"go-retry-template/policy"
	"strconv"
	"time"
)

func main() {
	// Not safe function
	NotSafe := func() (string, error) {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		// Randomly fail
		if r.Intn(2) == 0 {
			return "OK", nil
		} else {
			return "", errors.New("Bad thing happened")
		}
	}

	// Create very simple retry template
	rt := retry_template.RetryTemplate{
		RetryPolicy:   policy.SimpleRetryPolicy{MaxAttempts: 2},
		BackoffPolicy: backoff.FixedBackoffPolicy{BackoffPeriodMs: 200},
	}

	// Call not safe function under this retry logic
	_, err := rt.Execute(func(rc operations.RetryContext) (interface{}, error) {
		fmt.Println("NotSafe called. Retry Attempt: " + strconv.Itoa(rc.RetryCount))
		return NotSafe()
	})

	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println("OK ")
	}
}
