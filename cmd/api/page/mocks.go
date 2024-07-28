package page

import "time"

// MockLoad mocks Load function
func MockLoad(err error) Load {
	return func(relativeURL string, timeout time.Duration) error {
		return err
	}
}

// MockScroll mocks Scroll function
func MockScroll(err error) Scroll {
	return func() error {
		return err
	}
}
