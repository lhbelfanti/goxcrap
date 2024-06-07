package search

import "errors"

var (
	FailedToLoadAdvanceSearchPage = errors.New("failed to load advance search page")

	FailedToParseDate              = errors.New("failed to parse date")
	FailedToParseCriteriaSinceDate = errors.New("failed to parse criteria.Since date")
	FailedToParseCriteriaUntilDate = errors.New("failed to parse criteria.until date")
)
