package webdriver

import (
	"github.com/rs/zerolog/log"

	"goxcrap/internal/setup"
)

// NewManager initializes a new WebDriver with all its elements
type NewManager func() Manager

// MakeNewManager creates a new NewManager
func MakeNewManager(localMode bool) NewManager {
	return func() Manager {
		log.Info().Msg("Initializing WebDriver...")
		var manager Manager
		if localMode {
			manager = &LocalManager{}
		} else {
			manager = &DockerizedManager{}
		}
		setup.Must(manager.InitWebDriverService())
		setup.Must(manager.InitWebDriver())
		log.Info().Msg("WebDriver initialized!")

		return manager
	}
}
