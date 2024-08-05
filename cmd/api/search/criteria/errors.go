package criteria

import "errors"

const (
	InvalidRequestBody  string = "Invalid request payload"
	FailedToEnqueueTask string = "Failed to enqueue task"
)

var (
	FailedToParseDate              = errors.New("failed to parse date")
	FailedToParseCriteriaSinceDate = errors.New("failed to parse criteria since date")
	FailedToParseCriteriaUntilDate = errors.New("failed to parse criteria until date")
)
