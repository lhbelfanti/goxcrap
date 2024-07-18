package criteria

import "errors"

var (
	FailedToParseDate              = errors.New("failed to parse date")
	FailedToParseCriteriaSinceDate = errors.New("failed to parse criteria since date")
	FailedToParseCriteriaUntilDate = errors.New("failed to parse criteria until date")
)
