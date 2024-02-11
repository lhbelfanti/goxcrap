package element

import (
	"fmt"
	"time"
)

// RetrieveAndClickButton retrieves a button element and clicks on it
type RetrieveAndClickButton func(by, value, element string, timeout time.Duration, newError ErrorCreator) error

// MakeRetrieveAndClickButton creates a new RetrieveAndClickButton
func MakeRetrieveAndClickButton(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndClickButton {
	return func(by, value, element string, timeout time.Duration, newError ErrorCreator) error {
		button, err := waitAndRetrieveElement(by, value, timeout)
		if err != nil {
			return newError(fmt.Sprintf(FailedToRetrieveButton, element), err)
		}

		err = button.Click()
		if err != nil {
			return newError(fmt.Sprintf(FailedToClickButton, element), err)
		}

		return nil
	}
}
