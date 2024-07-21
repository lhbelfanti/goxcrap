package webdriver

import "errors"

var (
	CannotInitializeWebDriverService = errors.New("cannot initialize webdriver service")
	FailedToStopWebDriverService     = errors.New("failed to stop webdriver service")
	CannotInitializeWebDriver        = errors.New("cannot initialize webdriver")
	FailedToQuitWebDriver            = errors.New("failed to quit webdriver")
	CannotMaximizeWindow             = errors.New("cannot maximize window")
)
