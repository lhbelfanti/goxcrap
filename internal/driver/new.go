package driver

import (
	"log/slog"

	"github.com/fatih/color"
	"github.com/tebeka/selenium"

	"goxcrap/internal/setup"
)

// New initializes a new WebDriver with all its elements
type New func() (GoXCrapWebDriver, *selenium.Service, selenium.WebDriver)

// MakeNew creates a new New
func MakeNew(localMode bool) New {
	return func() (GoXCrapWebDriver, *selenium.Service, selenium.WebDriver) {
		slog.Info(color.BlueString("Initializing WebDriver..."))
		var goXCrapWebDriver GoXCrapWebDriver
		if localMode {
			goXCrapWebDriver = LocalWebDriver{}
		} else {
			goXCrapWebDriver = DockerizedWebDriver{}
		}
		service := setup.Init(goXCrapWebDriver.InitWebDriverService())
		webDriver := setup.Init(goXCrapWebDriver.InitWebDriver())
		slog.Info(color.GreenString("WebDriver initialized!"))

		return goXCrapWebDriver, service, webDriver
	}
}
