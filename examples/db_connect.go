package examples

import (
	"database/sql"
	"github.com/lib/pq"
	"go-retry-template"
	"go-retry-template/backoff"
	"go-retry-template/operations"
	"go-retry-template/policy"
	"os"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pqUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		// Fail
	}

	rt := retry_template.RetryTemplate{
		RetryPolicy:   policy.SimpleRetryPolicy{MaxAttempts: 2},
		BackoffPolicy: backoff.FixedBackoffPolicy{BackoffPeriodMs: 200},
	}

	result, err := rt.Execute(func(rc operations.RetryContext) (interface{}, error) {
		return sql.Open("postgres", pqUrl)
	})
	if err != nil {
		// Fail
	}

	// Assert type
	db = result.(*sql.DB)

	err = db.Ping()
	if err != nil {
		// Fail
	}

	return db
}
