package search

import (
	"time"

	"github.com/rs/zerolog/log"

	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/search/criteria"
)

const (
	pageLoaderTimeout time.Duration = 10 * time.Second
)

// ExecuteAdvanceSearch is the first implementation of a search to retrieve tweets then
type ExecuteAdvanceSearch func(searchCriteria criteria.Type) error

// MakeExecuteAdvanceSearch creates a new ExecuteAdvanceSearch
func MakeExecuteAdvanceSearch(loadPage page.Load) ExecuteAdvanceSearch {
	return func(searchCriteria criteria.Type) error {
		queryString := searchCriteria.ConvertIntoQueryString()
		err := loadPage("/search?"+queryString, pageLoaderTimeout)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToLoadAdvanceSearchPage
		}

		return nil
	}
}
