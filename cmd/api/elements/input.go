package elements

import (
	"context"
	"fmt"
	"time"

	"goxcrap/internal/log"
)

// RetrieveAndFillInput retrieves an input element, clicks on it, and fills it with the inputText param
type RetrieveAndFillInput func(ctx context.Context, by, value, element, inputText string, timeout time.Duration) error

// MakeRetrieveAndFillInput creates a new RetrieveAndFillInput
func MakeRetrieveAndFillInput(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndFillInput {
	return func(ctx context.Context, by, value, element, inputText string, timeout time.Duration) error {
		input, err := waitAndRetrieveElement(ctx, by, value, timeout)
		if err != nil {
			log.Err(ctx, err, fmt.Sprintf("element: %s", element))
			return FailedToRetrieveInput
		}

		err = input.Click()
		if err != nil {
			log.Err(ctx, err, fmt.Sprintf("element: %s", element))
			return FailedToClickInput
		}

		err = input.SendKeys(inputText)
		if err != nil {
			log.Err(ctx, err, fmt.Sprintf("element: %s", element))
			return FailedToFillInput
		}

		return nil
	}
}
