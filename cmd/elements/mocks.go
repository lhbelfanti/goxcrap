package elements

import (
	"time"

	"github.com/tebeka/selenium"
)

// MockWaitAndRetrieve mocks WaitAndRetrieve function
func MockWaitAndRetrieve(element selenium.WebElement, err error) WaitAndRetrieve {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		return element, err
	}
}

// MockWaitAndRetrieveCondition mocks WaitAndRetrieveCondition function
func MockWaitAndRetrieveCondition(elementFound bool) WaitAndRetrieveCondition {
	return func(by, value string) SeleniumCondition {
		return func(drv selenium.WebDriver) (bool, error) {
			return elementFound, nil
		}
	}
}

// MockWaitAndRetrieveAll mocks WaitAndRetrieveAll function
func MockWaitAndRetrieveAll(element selenium.WebElement, err error) WaitAndRetrieve {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		return element, err
	}
}

// MockWaitAndRetrieveAllCondition mocks WaitAndRetrieveAllCondition function
func MockWaitAndRetrieveAllCondition(elementFound bool) WaitAndRetrieveAllCondition {
	return func(by, value string) SeleniumCondition {
		return func(drv selenium.WebDriver) (bool, error) {
			return elementFound, nil
		}
	}
}

// MockRetrieveAndFillInput mocks RetrieveAndFillInput function
func MockRetrieveAndFillInput(err error, elementID string) RetrieveAndFillInput {
	return func(by, value, element, inputText string, timeout time.Duration, newError ErrorCreator) error {
		if elementID == element || elementID == "" {
			return err
		}

		return nil
	}
}

// MockRetrieveAndClickButton mocks RetrieveAndClickButton function
func MockRetrieveAndClickButton(err error, elementID string) RetrieveAndClickButton {
	return func(by, value, element string, timeout time.Duration, newError ErrorCreator) error {
		if elementID == element || elementID == "" {
			return err
		}

		return nil
	}
}
