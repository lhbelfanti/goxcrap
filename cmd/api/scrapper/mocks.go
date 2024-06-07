package scrapper

import "time"

// MockExecute mocks Execute function
func MockExecute(err error) Execute {
	return func(waitTimeAfterLogin time.Duration) error {
		return err
	}
}
