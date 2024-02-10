package scrapper

import (
	"time"

	"github.com/tebeka/selenium"
)

const TwitterURL string = "https://twitter.com"

// LoadPage waits until the page is fully loaded
type LoadPage func(page string, timeout time.Duration) error

// MakeLoadPage creates a new LoadPage
func MakeLoadPage(driver selenium.WebDriver) LoadPage {
	return func(relativeURL string, timeout time.Duration) error {
		err := driver.SetPageLoadTimeout(timeout)
		if err != nil {
			return NewScrapperError(FailedToSetPageLoadTimeout, err)
		}

		err = driver.Get(TwitterURL + relativeURL)
		if err != nil {
			return NewScrapperError(FailedToRetrievePage, err)
		}

		return nil
	}
}
