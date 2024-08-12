package search

import (
	"context"
	
	"goxcrap/cmd/api/search/criteria"
)

// MockExecuteAdvanceSearch mocks ExecuteAdvanceSearch function
func MockExecuteAdvanceSearch(err error) ExecuteAdvanceSearch {
	return func(ctx context.Context, searchCriteria criteria.Type) error {
		return err
	}
}
