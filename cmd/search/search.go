package search

import "goxcrap/cmd/page"

const (
	searchPageRelativeURL string = "/test"
)

// Search is the first implementation of a search to retrieve tweets then
type Search func(searchCriteria Criteria) error

// MakeSearch creates a new Search
func MakeSearch(loadPage page.Load) Search {
	return func(searchCriteria Criteria) error {
		return nil
	}
}
