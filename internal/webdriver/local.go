package webdriver

import (
	"context"
	"fmt"
	"strings"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

	"goxcrap/internal/log"
)

// LocalManager represents a Web Driver manager for the local version of this application
type LocalManager struct {
	service   *selenium.Service
	webDriver selenium.WebDriver
}

// InitWebDriverService initializes a new Chrome *selenium.Service
func (lwd *LocalManager) InitWebDriverService(ctx context.Context) error {
	log.Info(ctx, fmt.Sprintf("Initializing Chrome Driver Service using driver from: %s", chromeDriverPath))
	service, err := selenium.NewChromeDriverService(chromeDriverPath, chromeDriverServicePort)
	if err != nil {
		log.Error(ctx, err.Error())
		return CannotInitializeWebDriverService
	}

	lwd.service = service

	return nil
}

// InitWebDriver initializes a new Chrome selenium.WebDriver
func (lwd *LocalManager) InitWebDriver(ctx context.Context) error {
	chromeCaps := chrome.Capabilities{
		Prefs: capabilitiesPreferences,
		Args:  capabilitiesArgs,
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

	// maximize the current window to avoid responsive rendering
	err = wd.MaximizeWindow("")
	if err != nil {
		log.Error(ctx, err.Error())
		return CannotMaximizeWindow
	}

	lwd.webDriver = wd

	return nil
}

// Quit stops the selenium.WebDriver and its *selenium.Service to avoid leaks if the app is terminated
func (lwd *LocalManager) Quit(ctx context.Context) error {
	err := lwd.service.Stop()
	if err != nil {
		log.Error(ctx, err.Error())
		return FailedToStopWebDriverService
	}

	err = lwd.webDriver.Quit()
	if err != nil {
		log.Error(ctx, err.Error())
		return FailedToQuitWebDriver
	}

	return nil
}

// WebDriver returns the initialized selenium.WebDriver
func (lwd *LocalManager) WebDriver() selenium.WebDriver {
	return lwd.webDriver
}
