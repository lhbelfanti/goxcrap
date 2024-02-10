package auth

import (
	"errors"
	"fmt"
)

const (
	FailedToRetrieveInput  string = "Failed to retrieve %s input"
	FailedToClickInput     string = "Failed to click %s input"
	FailedToFillInput      string = "Failed to fill input with %s"
	FailedToRetrieveButton string = "Failed to retrieve %s button"
	FailedToClickButton    string = "Failed to click %s button"
)

// NewAuthError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewAuthError(description string, err error) error {
	newError := fmt.Sprintf("Auth: %s -> %v", description, err)
	return errors.New(newError)
}
