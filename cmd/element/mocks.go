package element

import (
	"time"

	"github.com/tebeka/selenium"
)

// MockMakeWaitAndRetrieve mocks the function MakeWaitAndRetrieve and the values returned by WaitAndRetrieve
func MockMakeWaitAndRetrieve(element selenium.WebElement, err error) WaitAndRetrieve {
	return func(by, value string, timeout time.Duration) (selenium.WebElement, error) {
		return element, err
	}
}

// MockMakeWaitAndRetrieveCondition mocks the function MakeWaitAndRetrieveCondition and the values returned by WaitAndRetrieve
func MockMakeWaitAndRetrieveCondition(elementFound bool) WaitAndRetrieveCondition {
	return func(by, value string) SeleniumCondition {
		return func(drv selenium.WebDriver) (bool, error) {
			return elementFound, nil
		}
	}
}

// MockMakeRetrieveAndFillInput mocks the function MakeRetrieveAndFillInput and the values returned by RetrieveAndFillInput
func MockMakeRetrieveAndFillInput(err error, elementID string) RetrieveAndFillInput {
	return func(by, value, element, inputText string, timeout time.Duration, newError ErrorCreator) error {
		if elementID == element || elementID == "" {
			return err
		}

		return nil
	}
}

// MockMakeRetrieveAndClickButton mocks the function MakeRetrieveAndClickButton and the values returned by RetrieveAndClickButton
func MockMakeRetrieveAndClickButton(err error, elementID string) RetrieveAndClickButton {
	return func(by, value, element string, timeout time.Duration, newError ErrorCreator) error {
		if elementID == element || elementID == "" {
			return err
		}

		return nil
	}
}
