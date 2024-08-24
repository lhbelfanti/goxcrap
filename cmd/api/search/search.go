package search

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/search/criteria"
)

// ExecuteAdvanceSearch is the first implementation of a search to retrieve tweets then
type ExecuteAdvanceSearch func(ctx context.Context, searchCriteria criteria.Type) error

// MakeExecuteAdvanceSearch creates a new ExecuteAdvanceSearch
func MakeExecuteAdvanceSearch(loadPage page.Load) ExecuteAdvanceSearch {
	pageLoaderTimeoutValue, _ := strconv.Atoi(os.Getenv("SEARCH_PAGE_TIMEOUT"))
	pageLoaderTimeout := time.Duration(pageLoaderTimeoutValue) * time.Second

	return func(ctx context.Context, searchCriteria criteria.Type) error {
		queryString := searchCriteria.ConvertIntoQueryString()
		err := loadPage(ctx, "/search?"+queryString, pageLoaderTimeout)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToLoadAdvanceSearchPage
		}

		return nil
	}
}
