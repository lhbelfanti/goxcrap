package page

import (
	"errors"
	"fmt"
)

const (
	FailedToSetPageLoadTimeout string = "Failed to driver.SetPageLoadTimeout"
	FailedToRetrievePage       string = "Failed to retrieve page"
)

// newPageError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func newPageError(description string, err error) error {
	newError := fmt.Sprintf("Page: %s -> %v", description, err)
	return errors.New(newError)
}
