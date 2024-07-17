package scrapper

import (
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/search"
)

// MockNew mocks New function
func MockNew(err error) New {
	return func(webDriver selenium.WebDriver) Execute {
		return MockExecute(err)
	}
}

// MockExecute mocks Execute function
func MockExecute(err error) Execute {
	return func(criteria search.Criteria, waitTimeAfterLogin time.Duration) error {
		return err
	}
}
