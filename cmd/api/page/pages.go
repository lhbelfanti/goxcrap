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

	// GoBack executes window.history.go(-1) to go to the previous page in the history
	GoBack func(ctx context.Context) error

	// OpenNewTab opens a new tab and navigates to the given page
	OpenNewTab func(ctx context.Context, page string, timeout time.Duration) error

	// CloseOpenedTabs closes all the opened tabs leaving active only the first one in the slice
	CloseOpenedTabs func(ctx context.Context) error
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

// MakeOpenNewTab creates a new OpenNewTab
func MakeOpenNewTab(wd selenium.WebDriver, loadPage Load) OpenNewTab {
	return func(ctx context.Context, page string, timeout time.Duration) error {
		jsOpenNewTab := "window.open('', '_blank');"
		_, err := wd.ExecuteScript(jsOpenNewTab, nil)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToOpenNewTab
		}

		handles, err := wd.WindowHandles()
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToObtainWindowHandles
		}

		err = wd.SwitchWindow(handles[len(handles)-1])
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToSwitchWindow
		}

		err = loadPage(ctx, page, timeout)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToLoadPageOnTheNewTab
		}

		return nil
	}
}

// MakeCloseOpenedTabs creates a new CloseOpenedTabs
func MakeCloseOpenedTabs(wd selenium.WebDriver) CloseOpenedTabs {
	return func(ctx context.Context) error {
		handles, err := wd.WindowHandles()
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToObtainWindowHandles
		}

		for i := 1; i < len(handles); i++ {
			err = wd.SwitchWindow(handles[i])
			if err != nil {
				log.Info(ctx, err.Error())
				return FailedToSwitchWindow
			}

			err = wd.Close()
			if err != nil {
				log.Info(ctx, err.Error())
				return FailedToCloseWindow
			}
		}

		err = wd.SwitchWindow(handles[0])
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToSwitchToMainWindow
		}

		return nil
	}
}
