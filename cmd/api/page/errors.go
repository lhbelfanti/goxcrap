package page

import "errors"

var (
	FailedToSetPageLoadTimeout = errors.New("failed to driver.SetPageLoadTimeout")
	FailedToRetrievePage       = errors.New("failed to retrieve page")

	FailedToGetInnerHeight = errors.New("failed to execute JavaScript to get innerHeight")
	FailedToScroll         = errors.New("failed to execute JavaScript to scroll")

	FailedToGoBack = errors.New("failed to execute JavaScript to go back to the previous page")
)
