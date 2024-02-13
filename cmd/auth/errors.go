package auth

import (
	"errors"
	"fmt"
)

// NewAuthError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewAuthError(description string, err error) error {
	newError := fmt.Sprintf("Auth: %s -> %v", description, err)
	return errors.New(newError)
}
