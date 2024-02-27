package elements

import (
	"time"

	"github.com/tebeka/selenium"
)

// MockWaitAndRetrieve mocks the function MakeWaitAndRetrieve and the values returned by WaitAndRetrieve
func MockWaitAndRetrieve(element selenium.WebElement, err error) WaitAndRetrieve {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		return element, err
	}
}

// MockWaitAndRetrieveCondition mocks the function MakeWaitAndRetrieveCondition and the values returned by WaitAndRetrieveCondition
func MockWaitAndRetrieveCondition(elementFound bool) WaitAndRetrieveCondition {
	return func(by, value string) SeleniumCondition {
		return func(drv selenium.WebDriver) (bool, error) {
			return elementFound, nil
		}
	}
}

// MockWaitAndRetrieveAll mocks the function MakeWaitAndRetrieveAll and the values returned by WaitAndRetrieveAll
func MockWaitAndRetrieveAll(element selenium.WebElement, err error) WaitAndRetrieve {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		return element, err
	}
}

// MockWaitAndRetrieveAllCondition mocks the function MakeWaitAndRetrieveAllCondition and the values returned by WaitAndRetrieveAllCondition
func MockWaitAndRetrieveAllCondition(elementFound bool) WaitAndRetrieveAllCondition {
	return func(by, value string) SeleniumCondition {
		return func(drv selenium.WebDriver) (bool, error) {
			return elementFound, nil
		}
	}
}

// MockRetrieveAndFillInput mocks the function MakeRetrieveAndFillInput and the values returned by RetrieveAndFillInput
func MockRetrieveAndFillInput(err error, elementID string) RetrieveAndFillInput {
	return func(by, value, element, inputText string, timeout time.Duration, newError ErrorCreator) error {
		if elementID == element || elementID == "" {
			return err
		}

		return nil
	}
}

// MockRetrieveAndClickButton mocks the function MakeRetrieveAndClickButton and the values returned by RetrieveAndClickButton
func MockRetrieveAndClickButton(err error, elementID string) RetrieveAndClickButton {
	return func(by, value, element string, timeout time.Duration, newError ErrorCreator) error {
		if elementID == element || elementID == "" {
			return err
		}

		return nil
	}
}
