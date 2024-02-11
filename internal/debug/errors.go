package debug

import (
	"errors"
	"fmt"
)

const (
	FailedToTakeScreenshot       string = "Failed to take screenshot"
	FailedToCreateScreenshotFile string = "Failed to create screenshot file"
	FailedToSaveScreenshotFile   string = "Failed to save screenshot file"
)

// NewScrapperError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewScrapperError(description string, err error) error {
	newError := fmt.Sprintf("Scrapper: %s -> %v", description, err)
	return errors.New(newError)
}
