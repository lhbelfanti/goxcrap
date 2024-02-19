package tweets

import (
	"errors"
	"fmt"
)

const (
	FailedToRetrieveArticles string = "Failed to retrieve articles"

	FailedToObtainTweetElement          = "Failed to obtain tweet element"
	FailedToObtainTweetTextElement      = "Failed to obtain tweet text element"
	FailedToObtainTweetTextParts        = "Failed to obtain tweet text parts"
	FailedToObtainTweetTextPartTagName  = "Failed to obtain tweet text part tag name"
	FailedToObtainTweetTextFromSpan     = "Failed to obtain tweet text from span"
	FailedToObtainTweetTimestampElement = "Failed to obtain tweet timestamp element"
	FailedToObtainTweetTimestamp        = "Failed to obtain tweet timestamp"
)

// NewTweetsError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewTweetsError(description string, err error) error {
	newError := fmt.Sprintf("Tweets: %s -> %v", description, err)
	return errors.New(newError)
}
