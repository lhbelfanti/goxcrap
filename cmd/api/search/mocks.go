package search

import "goxcrap/cmd/api/search/criteria"

// MockExecuteAdvanceSearch mocks ExecuteAdvanceSearch function
func MockExecuteAdvanceSearch(err error) ExecuteAdvanceSearch {
	return func(searchCriteria criteria.Type) error {
		return err
	}
}
