package search

import (
	"errors"
	"fmt"
)

const FailedToLoadAdvanceSearchPage string = "Failed to load advance search page"

// NewSearchError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewSearchError(description string, err error) error {
	newError := fmt.Sprintf("Search: %s -> %v", description, err)
	return errors.New(newError)
}
