package search

import "errors"

var (
	FailedToLoadAdvanceSearchPage = errors.New("failed to load advance search page")

	FailedToParseDate               = errors.New("failed to parse date")
	FailedToParseCriterionSinceDate = errors.New("failed to parse criterion.Since date")
	FailedToParseCriterionUntilDate = errors.New("failed to parse criterion.until date")
)
