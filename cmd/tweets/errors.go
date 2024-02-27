package tweets

import (
	"errors"
	"fmt"
)

const (
	FailedToRetrieveArticles string = "Failed to retrieve articles"

	FailedToObtainTweetTimestampElement string = "Failed to obtain tweet timestamp element"
	FailedToObtainTweetTimestamp        string = "Failed to obtain tweet timestamp"
	FailedToObtainTweetAuthorElement    string = "Failed to obtain tweet author element"
	FailedToObtainTweetAuthor           string = "Failed to obtain tweet author"

	// TODO: Review

	FailedToObtainTweetTextElement     string = "Failed to obtain tweet text element"
	FailedToObtainTweetTextParts       string = "Failed to obtain tweet text parts"
	FailedToObtainTweetTextPartTagName string = "Failed to obtain tweet text part tag name"
	FailedToObtainTweetTextFromSpan    string = "Failed to obtain tweet text from span"
	FailedToObtainTweetImagesElement   string = "Failed to obtain tweet images element"
	FailedToObtainTweetImages          string = "Failed to obtain tweet images"
	FailedToObtainTweetSrcFromImage    string = "Failed to obtain tweet src from image"
)

// NewTweetsError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewTweetsError(description string, err error) error {
	newError := fmt.Sprintf("Tweets: %s -> %v", description, err)
	return errors.New(newError)
}
