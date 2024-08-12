package webdriver

import (
	"context"

	"goxcrap/internal/log"
	"goxcrap/internal/setup"
)

// NewManager initializes a new WebDriver with all its elements
type NewManager func(ctx context.Context) Manager

// MakeNewManager creates a new NewManager
func MakeNewManager(localMode bool) NewManager {
	return func(ctx context.Context) Manager {
		log.Info(ctx, "Initializing WebDriver...")
		var manager Manager
		if localMode {
			manager = &LocalManager{}
		} else {
			manager = &DockerizedManager{}
		}
		setup.Must(manager.InitWebDriverService(ctx))
		setup.Must(manager.InitWebDriver(ctx))
		log.Info(ctx, "WebDriver initialized!")

		return manager
	}
}
