package element

import (
	"time"

	"github.com/tebeka/selenium"
)

type (
	// WaitAndRetrieve waits for a selenium.WebElement to be rendered in the page and retrieves it
	WaitAndRetrieve func(by, value string, timeout time.Duration) (selenium.WebElement, error)

	WaitAndRetrieveCondition func(by, value string) SeleniumCondition

	SeleniumCondition func(drv selenium.WebDriver) (bool, error)
)

// MakeWaitAndRetrieve creates a new WaitForElement
func MakeWaitAndRetrieve(driver selenium.WebDriver, condition WaitAndRetrieveCondition) WaitAndRetrieve {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		err := driver.WaitWithTimeout(selenium.Condition(condition(by, value)), timeout)
		if err != nil {
			return nil, NewElementError(FailedToExecuteWaitWithTimeout, err)
		}

		element, err := driver.FindElement(by, value)
		if err != nil {
			return nil, NewElementError(FailedToRetrieveElement, err)
		}

		return element, err
	}
}

// MakeWaitAndRetrieveCondition create a new WaitAndRetrieveCondition that returns SeleniumCondition
func MakeWaitAndRetrieveCondition() WaitAndRetrieveCondition {
	return func(by, value string) SeleniumCondition {
		return func(driver selenium.WebDriver) (bool, error) {
			element, err := driver.FindElement(by, value)
			if err == nil {
				return element.IsDisplayed()
			}

			return false, nil
		}
	}
}
