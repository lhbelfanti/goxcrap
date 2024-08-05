package elements

import (
	"time"

	"github.com/rs/zerolog/log"
)

// RetrieveAndClickButton retrieves a button element and clicks on it
type RetrieveAndClickButton func(by, value, element string, timeout time.Duration) error

// MakeRetrieveAndClickButton creates a new RetrieveAndClickButton
func MakeRetrieveAndClickButton(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndClickButton {
	return func(by, value, element string, timeout time.Duration) error {
		button, err := waitAndRetrieveElement(by, value, timeout)
		if err != nil {
			log.Error().Msgf("%s\nelement: %s", err.Error(), element)
			return FailedToRetrieveButton
		}

		err = button.Click()
		if err != nil {
			log.Error().Msgf("%s\nelement: %s", err.Error(), element)
			return FailedToClickButton
		}

		return nil
	}
}
