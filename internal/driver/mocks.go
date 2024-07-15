package driver

import "github.com/tebeka/selenium"

// MockNew mocks New function
func MockNew(goXCrapWebDriver GoXCrapWebDriver, service *selenium.Service, webDriver selenium.WebDriver) New {
	return func() (GoXCrapWebDriver, *selenium.Service, selenium.WebDriver) {
		return goXCrapWebDriver, service, webDriver
	}
}
