package elements

import (
	"log/slog"
	"time"
)

// RetrieveAndClickButton retrieves a button element and clicks on it
type RetrieveAndClickButton func(by, value, element string, timeout time.Duration) error

// MakeRetrieveAndClickButton creates a new RetrieveAndClickButton
func MakeRetrieveAndClickButton(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndClickButton {
	return func(by, value, element string, timeout time.Duration) error {
		button, err := waitAndRetrieveElement(by, value, timeout)
		if err != nil {
			slog.Error(err.Error(), slog.String("element", element))
			return FailedToRetrieveButton
		}

		err = button.Click()
		if err != nil {
			slog.Error(err.Error(), slog.String("element", element))
			return FailedToClickButton
		}

		return nil
	}
}
