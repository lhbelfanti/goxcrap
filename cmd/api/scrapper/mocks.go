package scrapper

import (
	"context"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/search/criteria"
)

// MockNew mocks New function
func MockNew(err error) New {
	return func(webDriver selenium.WebDriver) Execute {
		return MockExecute(err)
	}
}

// MockExecute mocks Execute function
func MockExecute(err error) Execute {
	return func(ctx context.Context, searchCriteria criteria.Type) error {
		return err
	}
}
