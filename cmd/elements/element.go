package elements

import (
	"time"

	"github.com/tebeka/selenium"
)

type (
	// WaitAndRetrieve waits for a selenium.WebElement to be rendered in the page and retrieves it
	WaitAndRetrieve func(by, value string, timeout time.Duration) (selenium.WebElement, error)

	// WaitAndRetrieveCondition is the function that returns the SeleniumCondition
	WaitAndRetrieveCondition func(by, value string) SeleniumCondition

	// WaitAndRetrieveAll waits for a slice of selenium.WebElement to be rendered in the page and retrieve them
	WaitAndRetrieveAll func(by, value string, timeout time.Duration) ([]selenium.WebElement, error)
	
	// WaitAndRetrieveAllCondition is the function that returns the SeleniumCondition
	WaitAndRetrieveAllCondition func(by, value string) SeleniumCondition

	// SeleniumCondition is the condition that WaitAndRetrieve uses to check if it has to wait for the element
	SeleniumCondition func(drv selenium.WebDriver) (bool, error)
)

// MakeWaitAndRetrieve creates a new WaitAndRetrieve
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

// MakeWaitAndRetrieveAll creates a new WaitAndRetrieveAll
func MakeWaitAndRetrieveAll(driver selenium.WebDriver, condition WaitAndRetrieveAllCondition) WaitAndRetrieveAll {
	return func(by, value string, timeout time.Duration) ([]selenium.WebElement, error) {
		err := driver.WaitWithTimeout(selenium.Condition(condition(by, value)), timeout)
		if err != nil {
			return nil, NewElementError(FailedToExecuteWaitWithTimeout, err)
		}

		elements, err := driver.FindElements(by, value)
		if err != nil {
			return nil, NewElementError(FailedToRetrieveElements, err)
		}

		return elements, err
	}
}

// MakeWaitAndRetrieveAllCondition create a new WaitAndRetrieveAllCondition that returns SeleniumCondition
func MakeWaitAndRetrieveAllCondition() WaitAndRetrieveAllCondition {
	return func(by, value string) SeleniumCondition {
		return func(driver selenium.WebDriver) (bool, error) {
			elements, err := driver.FindElements(by, value)
			if err == nil {
				if len(elements) > 0 {
					return elements[0].IsDisplayed()
				}

				return false, nil
			}

			return false, nil
		}
	}
}
