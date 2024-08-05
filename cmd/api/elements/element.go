package elements

import (
	"time"

	"github.com/rs/zerolog/log"
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
	SeleniumCondition func(wd selenium.WebDriver) (bool, error)
)

// MakeWaitAndRetrieve creates a new WaitAndRetrieve
func MakeWaitAndRetrieve(wd selenium.WebDriver, condition WaitAndRetrieveCondition) WaitAndRetrieve {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		err := wd.WaitWithTimeout(selenium.Condition(condition(by, value)), timeout)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, FailedToExecuteWaitWithTimeout
		}

		element, err := wd.FindElement(by, value)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, FailedToRetrieveElement
		}

		return element, err
	}
}

// MakeWaitAndRetrieveCondition create a new WaitAndRetrieveCondition that returns SeleniumCondition
func MakeWaitAndRetrieveCondition() WaitAndRetrieveCondition {
	return func(by, value string) SeleniumCondition {
		return func(wd selenium.WebDriver) (bool, error) {
			element, err := wd.FindElement(by, value)
			if err == nil {
				return element.IsDisplayed()
			}

			return false, nil
		}
	}
}

// MakeWaitAndRetrieveAll creates a new WaitAndRetrieveAll
func MakeWaitAndRetrieveAll(wd selenium.WebDriver, condition WaitAndRetrieveAllCondition) WaitAndRetrieveAll {
	return func(by, value string, timeout time.Duration) ([]selenium.WebElement, error) {
		err := wd.WaitWithTimeout(selenium.Condition(condition(by, value)), timeout)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, FailedToExecuteWaitWithTimeout
		}

		elements, err := wd.FindElements(by, value)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, FailedToRetrieveElements
		}

		return elements, err
	}
}

// MakeWaitAndRetrieveAllCondition create a new WaitAndRetrieveAllCondition that returns SeleniumCondition
func MakeWaitAndRetrieveAllCondition() WaitAndRetrieveAllCondition {
	return func(by, value string) SeleniumCondition {
		return func(wd selenium.WebDriver) (bool, error) {
			elements, err := wd.FindElements(by, value)
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
