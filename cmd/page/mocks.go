package page

import (
	"time"
)

// MockMakeLoad mocks the function MakeLoad and the values returned by Load
func MockMakeLoad(err error) Load {
	return func(relativeURL string, timeout time.Duration) error {
		return err
	}
}
