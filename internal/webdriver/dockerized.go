package webdriver

import (
	"context"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"os"
	"strings"

	"goxcrap/internal/log"
)

// DockerizedManager represents a Web Driver manager for the dockerized version of this application
type DockerizedManager struct {
	service   *selenium.Service
	webDriver selenium.WebDriver
}

// InitWebDriverService initializes a new Chrome *selenium.Service
func (dwd *DockerizedManager) InitWebDriverService(ctx context.Context) error {
	driverPath := os.Getenv("DRIVER_PATH")
	if driverPath == "" {
		driverPath = chromeDriverPath
	}

	log.Info(ctx, fmt.Sprintf("Initializing Chrome Driver Service using driver from: %s", driverPath))
	service, err := selenium.NewChromeDriverService(driverPath, chromeDriverServicePort)
	if err != nil {
		log.Error(ctx, err.Error())
		return CannotInitializeWebDriverService
	}

	dwd.service = service

	return nil

}

// InitWebDriver initializes a new Chrome selenium.WebDriver
func (dwd *DockerizedManager) InitWebDriver(ctx context.Context) error {
	browserPath := os.Getenv("BROWSER_PATH")

	capabilitiesArgs = append(capabilitiesArgs, "--headless")

	chromeCaps := chrome.Capabilities{
		Prefs: capabilitiesPreferences,
		Args:  capabilitiesArgs,
	}

	if browserPath != "" {
		chromeCaps.Path = browserPath
	}

	if chromeCaps.Path != "" {
		log.Info(ctx, fmt.Sprintf("Setting up Chrome Capacities using the following Args: ( %s ) and the following Path: %s", strings.Join(chromeCaps.Args, " | "), chromeCaps.Path))
	} else {
		log.Info(ctx, fmt.Sprintf("Setting up Chrome Capacities using the following Args: ( %s )", strings.Join(chromeCaps.Args, " | ")))
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)

	remotePath := fmt.Sprintf("http://localhost:%d/wd/hub", chromeDriverServicePort)
	log.Info(ctx, fmt.Sprintf("Creating Remote Client at: \n%s", remotePath))
	wd, err := selenium.NewRemote(caps, remotePath)
	if err != nil {
		log.Error(ctx, err.Error())
		return CannotInitializeWebDriver
	}

	dwd.webDriver = wd

	return nil
}

// Quit stops the selenium.WebDriver and its *selenium.Service to avoid leaks if the app is terminated
func (dwd *DockerizedManager) Quit(ctx context.Context) error {
	err := dwd.service.Stop()
	if err != nil {
		log.Error(ctx, err.Error())
		return FailedToStopWebDriverService
	}

	err = dwd.webDriver.Quit()
	if err != nil {
		log.Error(ctx, err.Error())
		return FailedToQuitWebDriver
	}

	return nil
}

// WebDriver returns the initialized selenium.WebDriver
func (dwd *DockerizedManager) WebDriver() selenium.WebDriver {
	return dwd.webDriver
}
