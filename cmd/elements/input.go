package elements

import (
	"fmt"
	"time"
)

type (
	// RetrieveAndFillInput retrieves an input element, clicks on it, and fills it with the inputText param
	RetrieveAndFillInput func(by, value, element, inputText string, timeout time.Duration, newError ErrorCreator) error
)

// MakeRetrieveAndFillInput creates a new RetrieveAndFillInput
func MakeRetrieveAndFillInput(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndFillInput {
	return func(by, value, element, inputText string, timeout time.Duration, newError ErrorCreator) error {
		input, err := waitAndRetrieveElement(by, value, timeout)
		if err != nil {
			return newError(fmt.Sprintf(FailedToRetrieveInput, element), err)
		}

		err = input.Click()
		if err != nil {
			return newError(fmt.Sprintf(FailedToClickInput, element), err)
		}

		err = input.SendKeys(inputText)
		if err != nil {
			return newError(fmt.Sprintf(FailedToFillInput, element), err)
		}

		return nil
	}
}
