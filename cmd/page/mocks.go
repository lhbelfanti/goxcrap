package page

import (
	"time"
)

// MockLoad mocks the function MakeLoad and the values returned by Load
func MockLoad(err error) Load {
	return func(relativeURL string, timeout time.Duration) error {
		return err
	}
}
