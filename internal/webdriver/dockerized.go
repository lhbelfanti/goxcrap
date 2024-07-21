package webdriver

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// DockerizedManager represents a Web Driver manager for the dockerized version of this application
type DockerizedManager struct {
	service   *selenium.Service
	webDriver selenium.WebDriver
}

// InitWebDriverService initializes a new Chrome *selenium.Service
func (dwd DockerizedManager) InitWebDriverService() error {
	driverPath := os.Getenv("DRIVER_PATH")
	if driverPath == "" {
		driverPath = chromeDriverPath
	}

	slog.Info(fmt.Sprintf(color.BlueString("Initializing Chrome Driver Service using driver from:\n%s"), color.GreenString(driverPath)))
	service, err := selenium.NewChromeDriverService(driverPath, chromeDriverServicePort)
	if err != nil {
		slog.Error(err.Error())
		return CannotInitializeWebDriverService
	}

	dwd.service = service

	return nil

}

// InitWebDriver initializes a new Chrome selenium.WebDriver
func (dwd DockerizedManager) InitWebDriver() error {
	browserPath := os.Getenv("BROWSER_PATH")

	capabilitiesArgs = append(capabilitiesArgs, "--headless")

	chromeCaps := chrome.Capabilities{
		Prefs: capabilitiesPreferences,
		Args:  capabilitiesArgs,
	}

	if browserPath != "" {
		chromeCaps.Path = browserPath
	}

	slog.Info(fmt.Sprintf(color.BlueString("Setting up Chrome Capacities using the following Args:\n%s\n"), color.GreenString(strings.Join(chromeCaps.Args, "\n"))))
	if chromeCaps.Path != "" {
		slog.Info(fmt.Sprintf("and the following Path:\n%s", color.GreenString(chromeCaps.Path)))
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)

	remotePath := fmt.Sprintf("http://localhost:%d/wd/hub", chromeDriverServicePort)
	slog.Info(color.BlueString(fmt.Sprintf("Creating Remote Client at: \n%s", color.GreenString(remotePath))))
	wd, err := selenium.NewRemote(caps, remotePath)
	if err != nil {
		slog.Error(err.Error())
		return CannotInitializeWebDriver
	}

	dwd.webDriver = wd

	return nil
}

// Quit stops the selenium.WebDriver and its *selenium.Service to avoid leaks if the app is terminated
func (dwd DockerizedManager) Quit() error {
	err := dwd.service.Stop()
	if err != nil {
		slog.Error(err.Error())
		return FailedToStopWebDriverService
	}

	err = dwd.webDriver.Quit()
	if err != nil {
		slog.Error(err.Error())
		return FailedToQuitWebDriver
	}

	return nil
}

// WebDriver returns the initialized selenium.WebDriver
func (dwd DockerizedManager) WebDriver() selenium.WebDriver {
	return dwd.webDriver
}
