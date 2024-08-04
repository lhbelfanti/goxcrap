package webdriver

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// LocalManager represents a Web Driver manager for the local version of this application
type LocalManager struct {
	service   *selenium.Service
	webDriver selenium.WebDriver
}

// InitWebDriverService initializes a new Chrome *selenium.Service
func (lwd *LocalManager) InitWebDriverService() error {
	slog.Info(fmt.Sprintf("Initializing Chrome Driver Service using driver from:\n%s", chromeDriverPath))
	service, err := selenium.NewChromeDriverService(chromeDriverPath, chromeDriverServicePort)
	if err != nil {
		slog.Error(err.Error())
		return CannotInitializeWebDriverService
	}

	lwd.service = service

	return nil
}

// InitWebDriver initializes a new Chrome selenium.WebDriver
func (lwd *LocalManager) InitWebDriver() error {
	chromeCaps := chrome.Capabilities{
		Prefs: capabilitiesPreferences,
		Args:  capabilitiesArgs,
	}

	slog.Info(fmt.Sprintf("Setting up Chrome Capacities using the following Args:\n%s\n", strings.Join(chromeCaps.Args, "\n")))
	if chromeCaps.Path != "" {
		slog.Info(fmt.Sprintf("and the following Path:\n%s", chromeCaps.Path))
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)

	remotePath := fmt.Sprintf("http://localhost:%d/wd/hub", chromeDriverServicePort)
	slog.Info(fmt.Sprintf("Creating Remote Client at: \n%s", remotePath))
	wd, err := selenium.NewRemote(caps, remotePath)
	if err != nil {
		slog.Error(err.Error())
		return CannotInitializeWebDriver
	}

	// maximize the current window to avoid responsive rendering
	err = wd.MaximizeWindow("")
	if err != nil {
		slog.Error(err.Error())
		return CannotMaximizeWindow
	}

	lwd.webDriver = wd

	return nil
}

// Quit stops the selenium.WebDriver and its *selenium.Service to avoid leaks if the app is terminated
func (lwd *LocalManager) Quit() error {
	err := lwd.service.Stop()
	if err != nil {
		slog.Error(err.Error())
		return FailedToStopWebDriverService
	}

	err = lwd.webDriver.Quit()
	if err != nil {
		slog.Error(err.Error())
		return FailedToQuitWebDriver
	}

	return nil
}

// WebDriver returns the initialized selenium.WebDriver
func (lwd *LocalManager) WebDriver() selenium.WebDriver {
	return lwd.webDriver
}
