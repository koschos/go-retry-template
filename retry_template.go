package retry_template

import "github.com/koschos/go-retry-template/operations"

type RetryTemplate struct {
	RetryPolicy   operations.RetryPolicy
	BackoffPolicy operations.BackoffPolicy
}

func (rt RetryTemplate) Execute(callback func(rc operations.RetryContext) (interface{}, error)) (interface{}, error) {
	context := rt.RetryPolicy.Open()
	defer rt.RetryPolicy.Close(&context)

	for rt.RetryPolicy.CanRetry(context) {
		result, err := callback(context)
		if err != nil {
			rt.RetryPolicy.RegisterError(&context, err)
			if rt.RetryPolicy.CanRetry(context) {
				rt.BackoffPolicy.BackOff(context)
			}
		} else {
			return result, nil
		}
	}

	return nil, context.LastError
}
