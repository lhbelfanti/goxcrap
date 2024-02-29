package page

import (
	"errors"
)

var (
	FailedToSetPageLoadTimeout = errors.New("failed to driver.SetPageLoadTimeout")
	FailedToRetrievePage       = errors.New("failed to retrieve page")
)
