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

// MockOpenNewTab mocks OpenNewTab function
func MockOpenNewTab(err error) OpenNewTab {
	return func(ctx context.Context, page string, timeout time.Duration) error {
		return err
	}
}

// MockCloseOpenedTabs mocks CloseOpenedTabs function
func MockCloseOpenedTabs(err error) CloseOpenedTabs {
	return func(ctx context.Context) error {
		return err
	}
}
