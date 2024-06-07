package elements

import (
	"log/slog"
	"time"
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
			slog.Error(err.Error(), slog.String("element", element))
			return FailedToRetrieveInput
		}

		err = input.Click()
		if err != nil {
			slog.Error(err.Error(), slog.String("element", element))
			return FailedToClickInput
		}

		err = input.SendKeys(inputText)
		if err != nil {
			slog.Error(err.Error(), slog.String("element", element))
			return FailedToFillInput
		}

		return nil
	}
}
