package page

import (
	"context"
	"time"
)

// MockLoad mocks Load function
func MockLoad(err error) Load {
	return func(ctx context.Context, relativeURL string, timeout time.Duration) error {
		return err
	}
}

// MockScroll mocks Scroll function
func MockScroll(err error) Scroll {
	return func(ctx context.Context) error {
		return err
	}
}
