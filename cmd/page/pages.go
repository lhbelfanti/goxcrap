package page

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/tebeka/selenium"
)

const TwitterURL string = "https://x.com"

type (
	// Load waits until the page is fully loaded
	Load func(page string, timeout time.Duration) error

	// Scroll executes window.scrollTo, to scroll the page
	Scroll func() error
)

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

// MakeScroll creates a new Scroll
func MakeScroll(driver selenium.WebDriver) Scroll {
	return func() error {
		jsHeight := `return window.innerHeight;`
		height, err := driver.ExecuteScript(jsHeight, nil)
		if err != nil {
			slog.Error(err.Error())
			return FailedToGetInnerHeight
		}

		// TODO: change  %v * 2 by the exact amount it should scroll
		jsScroll := fmt.Sprintf("window.scrollBy(0, %v * 2);", height)
		_, err = driver.ExecuteScript(jsScroll, nil)
		if err != nil {
			slog.Error(err.Error())
			return FailedToScroll
		}

		return nil
	}
}
