package elements

import (
	"errors"
)

var (
	FailedToExecuteWaitWithTimeout = errors.New("failed to execute wd.WaitWithTimeout")
	FailedToRetrieveElement        = errors.New("failed to retrieve element")
	FailedToRetrieveElements       = errors.New("failed to retrieve elements")

	FailedToRetrieveInput = errors.New("failed to retrieve input")
	FailedToClickInput    = errors.New("failed to click input")
	FailedToFillInput     = errors.New("failed to fill input")

	FailedToRetrieveButton = errors.New("failed to retrieve button")
	FailedToClickButton    = errors.New("failed to click button")
)
