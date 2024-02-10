package scrapper

import (
	"errors"
	"fmt"
)

var (
	FailedToSetPageLoadTimeout     = "Failed to driver.SetPageLoadTimeout"
	FailedToRetrievePage           = "Failed to retrieve page"
	FailedToExecuteWaitWithTimeout = "Failed to execute driver.WaitWithTimeout"
	FailedToRetrieveElement        = "Failed to retrieve element"
	FailedToTakeScreenshot         = "Failed to take screenshot"
	FailedToCreateScreenshotFile   = "Failed to create screenshot file"
	FailedToSaveScreenshotFile     = "Failed to save screenshot file"
)

// NewScrapperError creates a new error based on a description and the error
// It adds the package name to identify easily where the error comes from
func NewScrapperError(description string, err error) error {
	newError := fmt.Sprintf("Scrapper: %s -> %v", description, err)
	return errors.New(newError)
}
