package search

import (
	"log/slog"
	"time"

	"goxcrap/cmd/api/page"
)

const (
	pageLoaderTimeout time.Duration = 10 * time.Second
)

// ExecuteAdvanceSearch is the first implementation of a search to retrieve tweets then
type ExecuteAdvanceSearch func(criterion Criterion) error

// MakeExecuteAdvanceSearch creates a new ExecuteAdvanceSearch
func MakeExecuteAdvanceSearch(loadPage page.Load) ExecuteAdvanceSearch {
	return func(criterion Criterion) error {
		queryString := criterion.ConvertIntoQueryString()
		err := loadPage("/search?"+queryString, pageLoaderTimeout)
		if err != nil {
			slog.Error(err.Error())
			return FailedToLoadAdvanceSearchPage
		}

		return nil
	}
}
