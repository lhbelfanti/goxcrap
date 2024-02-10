package scrapper

import (
	"time"

	"github.com/tebeka/selenium"
)

type (
	// WaitAndRetrieveElement waits for a selenium.WebElement to be rendered in the page and retrieves it
	WaitAndRetrieveElement func(by, value string, timeout time.Duration) (selenium.WebElement, error)
)

// MakeWaitAndRetrieveElement creates a new WaitForElement
func MakeWaitAndRetrieveElement(driver selenium.WebDriver) WaitAndRetrieveElement {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		var element selenium.WebElement
		err := driver.WaitWithTimeout(func(drv selenium.WebDriver) (bool, error) {
			element, _ = drv.FindElement(by, value)
			if element != nil {
				return element.IsDisplayed()
			}

			return false, nil
		}, timeout)

		if err != nil {
			return nil, NewScrapperError(FailedToExecuteWaitWithTimeout, err)
		}

		if element == nil {
			return nil, NewScrapperError(FailedToRetrieveElement, err)
		}

		return element, err
	}
}
