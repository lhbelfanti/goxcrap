package driver

import "github.com/tebeka/selenium"

// GoXCrapWebDriver interface adopted by the different implementations of a web driver used in this application
type GoXCrapWebDriver interface {
	InitWebDriverService() (*selenium.Service, error)
	StopWebDriverService(*selenium.Service)
	QuitWebDriver(selenium.WebDriver)
	InitWebDriver() (selenium.WebDriver, error)
}

const (
	chromeDriverPath        string = "./internal/driver/chromedriver"
	chromeDriverServicePort int    = 9515
)
