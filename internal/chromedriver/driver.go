package chromedriver

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	chromeDriverPath        string = "./internal/chromedriver/driver"
	chromeDriverServicePort int    = 4444
)

// InitWebDriverService creates a new Chrome Driver Service
func InitWebDriverService() (*selenium.Service, error) {
	return selenium.NewChromeDriverService(chromeDriverPath, chromeDriverServicePort)
}

func StopWebDriverService(service *selenium.Service) {
	err := service.Stop()
	if err != nil {
		panic(err)
	}
}

// InitWebDriver creates a new Chrome WebDriver
func InitWebDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless-new", // comment out this line for testing
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		return nil, err
	}

	return driver, nil
}
