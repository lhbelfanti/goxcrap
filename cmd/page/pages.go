package page

import (
	"log/slog"
	"time"

	"github.com/tebeka/selenium"
)

const TwitterURL string = "https://twitter.com"

// Load waits until the page is fully loaded
type Load func(page string, timeout time.Duration) error

// MakeLoad creates a new Load
func MakeLoad(driver selenium.WebDriver) Load {
	return func(relativeURL string, timeout time.Duration) error {
		err := driver.SetPageLoadTimeout(timeout)
		if err != nil {
			slog.Error(err.Error())
			return FailedToSetPageLoadTimeout
		}

		err = driver.Get(TwitterURL + relativeURL)
		if err != nil {
			slog.Error(err.Error())
			return FailedToRetrievePage
		}

		return nil
	}
}
