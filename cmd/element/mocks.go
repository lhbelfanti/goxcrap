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
