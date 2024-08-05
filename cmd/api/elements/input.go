package elements

import (
	"time"

	"github.com/rs/zerolog/log"
)

type (
	// RetrieveAndFillInput retrieves an input element, clicks on it, and fills it with the inputText param
	RetrieveAndFillInput func(by, value, element, inputText string, timeout time.Duration) error
)

// MakeRetrieveAndFillInput creates a new RetrieveAndFillInput
func MakeRetrieveAndFillInput(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndFillInput {
	return func(by, value, element, inputText string, timeout time.Duration) error {
		input, err := waitAndRetrieveElement(by, value, timeout)
		if err != nil {
			log.Error().Msgf("%s\nelement: %s", err.Error(), element)
			return FailedToRetrieveInput
		}

		err = input.Click()
		if err != nil {
			log.Error().Msgf("%s\nelement: %s", err.Error(), element)
			return FailedToClickInput
		}

		err = input.SendKeys(inputText)
		if err != nil {
			log.Error().Msgf("%s\nelement: %s", err.Error(), element)
			return FailedToFillInput
		}

		return nil
	}
}
