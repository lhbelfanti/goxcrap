package elements

import (
	"context"
	"fmt"
	"time"

	"goxcrap/internal/log"
)

// RetrieveAndClickButton retrieves a button element and clicks on it
type RetrieveAndClickButton func(ctx context.Context, by, value, element string, timeout time.Duration) error

// MakeRetrieveAndClickButton creates a new RetrieveAndClickButton
func MakeRetrieveAndClickButton(waitAndRetrieveElement WaitAndRetrieve) RetrieveAndClickButton {
	return func(ctx context.Context, by, value, element string, timeout time.Duration) error {
		button, err := waitAndRetrieveElement(ctx, by, value, timeout)
		if err != nil {
			log.Err(ctx, err, fmt.Sprintf("element: %s", element))
			return FailedToRetrieveButton
		}

		err = button.Click()
		if err != nil {
			log.Err(ctx, err, fmt.Sprintf("element: %s", element))
			return FailedToClickButton
		}

		return nil
	}
}
