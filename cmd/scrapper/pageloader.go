package scrapper

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

// PageLoader waits until the page is fully loaded
type PageLoader func(page string, timeout time.Duration) error

// MakePageLoader creates a new PageLoader
func MakePageLoader(driver selenium.WebDriver) PageLoader {
	return func(page string, timeout time.Duration) error {
		err := driver.Get(page)
		if err != nil {
			return errors.New(fmt.Sprintf("%s -- %s", CantAccessRequestedPage, err.Error()))
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				readyState, err := driver.ExecuteScript("return document.readyState", nil)
				if err != nil {
					return err
				}
				if readyState == "complete" {
					return nil
				}
				time.Sleep(500 * time.Millisecond) // Check every 500 milliseconds
			}
		}
	}
}
