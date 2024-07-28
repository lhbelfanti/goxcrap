package debug

import "errors"

var (
	FailedToTakeScreenshot       = errors.New("failed to take screenshot")
	FailedToCreateScreenshotFile = errors.New("failed to create screenshot file")
	FailedToSaveScreenshotFile   = errors.New("failed to save screenshot file")
)
