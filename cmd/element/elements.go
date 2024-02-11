package element

import (
	"time"

	"github.com/tebeka/selenium"
)

type (
	// WaitAndRetrieve waits for a selenium.WebElement to be rendered in the page and retrieves it
	WaitAndRetrieve func(by, value string, timeout time.Duration) (selenium.WebElement, error)
)

// MakeWaitAndRetrieve creates a new WaitForElement
func MakeWaitAndRetrieve(driver selenium.WebDriver) WaitAndRetrieve {
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
			return nil, NewElementError(FailedToExecuteWaitWithTimeout, err)
		}

		if element == nil {
			return nil, NewElementError(FailedToRetrieveElement, err)
		}

		return element, err
	}
}
