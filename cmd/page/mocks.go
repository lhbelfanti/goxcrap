package page

import (
	"time"
)

// MockLoad mocks Load function
func MockLoad(err error) Load {
	return func(relativeURL string, timeout time.Duration) error {
		return err
	}
}
