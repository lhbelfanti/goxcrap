package driver

import "github.com/tebeka/selenium"

// GoXCrapWebDriver interface adopted by the different implementations of a web driver used in this application
type GoXCrapWebDriver interface {
	// InitWebDriverService initializes a new *selenium.Service
	InitWebDriverService() (*selenium.Service, error)

	// StopWebDriverService stops a *selenium.Service
	StopWebDriverService(*selenium.Service)

	// InitWebDriver initializes a new selenium.WebDriver
	InitWebDriver() (selenium.WebDriver, error)

	// QuitWebDriver stops a selenium.WebDriver
	QuitWebDriver(selenium.WebDriver)
}

const (
	chromeDriverPath        string = "./internal/driver/chromedriver"
	chromeDriverServicePort int    = 9515
)
