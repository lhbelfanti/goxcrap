package page

import "errors"

var (
	FailedToSetPageLoadTimeout = errors.New("failed to driver.SetPageLoadTimeout")
	FailedToRetrievePage       = errors.New("failed to retrieve page")

	FailedToGetInnerHeight = errors.New("failed to execute JavaScript to get innerHeight")
	FailedToScroll         = errors.New("failed to execute JavaScript to scroll")

	FailedToOpenNewTab          = errors.New("failed to open new tab")
	FailedToObtainWindowHandles = errors.New("failed to obtain window handles")
	FailedToSwitchWindow        = errors.New("failed to switch window")
	FailedToLoadPageOnTheNewTab = errors.New("failed to load page on the new tab")
	FailedToCloseWindow         = errors.New("failed to close window")
	FailedToSwitchToMainWindow  = errors.New("failed to switch to main window")
)
