package page

import (
	"context"
	"fmt"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const twitterURL string = "https://x.com"

type (
	// Load waits until the page is fully loaded
	Load func(ctx context.Context, page string, timeout time.Duration) error

	// Scroll executes window.scrollTo, to scroll the page
	Scroll func(ctx context.Context) error
)

// MakeLoad creates a new Load
func MakeLoad(wd selenium.WebDriver) Load {
	return func(ctx context.Context, relativeURL string, timeout time.Duration) error {
		err := wd.SetPageLoadTimeout(timeout)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToSetPageLoadTimeout
		}

		pageURL := twitterURL + relativeURL
		ctx = log.With(ctx, log.Param("page_url", pageURL))
		log.Debug(ctx, fmt.Sprintf("Accessing page: %s", pageURL))

		err = wd.Get(pageURL)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToRetrievePage
		}

		return nil
	}
}

// MakeScroll creates a new Scroll
func MakeScroll(wd selenium.WebDriver) Scroll {
	return func(ctx context.Context) error {
		jsHeight := `return window.innerHeight;`
		height, err := wd.ExecuteScript(jsHeight, nil)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToGetInnerHeight
		}

		// TODO: change  %v * 2 by the exact amount it should scroll
		jsScroll := fmt.Sprintf("window.scrollBy(0, %v * 2);", height)
		_, err = wd.ExecuteScript(jsScroll, nil)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToScroll
		}

		return nil
	}
}
