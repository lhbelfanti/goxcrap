package webdriver

import (
	"log/slog"

	"goxcrap/internal/setup"
)

// NewManager initializes a new WebDriver with all its elements
type NewManager func() Manager

// MakeNewManager creates a new NewManager
func MakeNewManager(localMode bool) NewManager {
	return func() Manager {
		slog.Info("Initializing WebDriver...")
		var manager Manager
		if localMode {
			manager = &LocalManager{}
		} else {
			manager = &DockerizedManager{}
		}
		setup.Must(manager.InitWebDriverService())
		setup.Must(manager.InitWebDriver())
		slog.Info("WebDriver initialized!")

		return manager
	}
}
