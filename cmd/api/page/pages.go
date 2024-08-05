package page

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/tebeka/selenium"
)

const twitterURL string = "https://x.com"

type (
	// Load waits until the page is fully loaded
	Load func(page string, timeout time.Duration) error

	// Scroll executes window.scrollTo, to scroll the page
	Scroll func() error
)

// MakeLoad creates a new Load
func MakeLoad(wd selenium.WebDriver) Load {
	return func(relativeURL string, timeout time.Duration) error {
		err := wd.SetPageLoadTimeout(timeout)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToSetPageLoadTimeout
		}

		pageURL := twitterURL + relativeURL
		log.Info().Msgf("Accessing page: %s", pageURL)
		err = wd.Get(pageURL)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToRetrievePage
		}

		return nil
	}
}

// MakeScroll creates a new Scroll
func MakeScroll(wd selenium.WebDriver) Scroll {
	return func() error {
		jsHeight := `return window.innerHeight;`
		height, err := wd.ExecuteScript(jsHeight, nil)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToGetInnerHeight
		}

		// TODO: change  %v * 2 by the exact amount it should scroll
		jsScroll := fmt.Sprintf("window.scrollBy(0, %v * 2);", height)
		_, err = wd.ExecuteScript(jsScroll, nil)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToScroll
		}

		return nil
	}
}
