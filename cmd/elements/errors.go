package elements

import (
	"errors"
	"fmt"
)

const (
	FailedToExecuteWaitWithTimeout string = "Failed to execute driver.WaitWithTimeout"
	FailedToRetrieveElement        string = "Failed to retrieve element"
	FailedToRetrieveElements       string = "Failed to retrieve elements"

	FailedToRetrieveInput string = "Failed to retrieve %s input"
	FailedToClickInput    string = "Failed to click %s input"
	FailedToFillInput     string = "Failed to fill input with %s"

	FailedToRetrieveButton string = "Failed to retrieve %s button"
	FailedToClickButton    string = "Failed to click %s button"
)

type (
	// ErrorCreator necessary function to create an error based on the current error and the specific package where it was triggered
	ErrorCreator func(description string, err error) error
)

// NewElementError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewElementError(description string, err error) error {
	newError := fmt.Sprintf("Element: %s -> %v", description, err)
	return errors.New(newError)
}
