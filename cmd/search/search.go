package search

import (
	"time"

	"goxcrap/cmd/page"
)

const (
	pageLoaderTimeout time.Duration = 10 * time.Second
)

// ExecuteAdvanceSearch is the first implementation of a search to retrieve tweets then
type ExecuteAdvanceSearch func(searchCriteria Criteria) error

// MakeExecuteAdvanceSearch creates a new ExecuteAdvanceSearch
func MakeExecuteAdvanceSearch(loadPage page.Load) ExecuteAdvanceSearch {
	return func(searchCriteria Criteria) error {
		queryString := searchCriteria.ConvertIntoQueryString()
		err := loadPage("/search?"+queryString, pageLoaderTimeout)
		if err != nil {
			return NewSearchError(FailedToLoadAdvanceSearchPage, err)
		}

		return nil
	}
}
