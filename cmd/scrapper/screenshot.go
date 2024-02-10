package scrapper

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

// TakeScreenshot takes a screenshot of the current state
type TakeScreenshot func() error

// MakeTakeScreenshot creates a new TakeScreenshot
func MakeTakeScreenshot(driver selenium.WebDriver) TakeScreenshot {
	return func() error {
		screenshotBytes, err := driver.Screenshot()
		if err != nil {
			return NewScrapperError(FailedToTakeScreenshot, err)
		}

		img, _, _ := image.Decode(bytes.NewReader(screenshotBytes))
		screenshotName := fmt.Sprintf("./twitter-%v.png", time.Now().Format(time.RFC822))
		out, err := os.Create(screenshotName)
		if err != nil {
			return NewScrapperError(FailedToCreateScreenshotFile, err)
		}

		err = png.Encode(out, img)
		if err != nil {
			return NewScrapperError(FailedToSaveScreenshotFile, err)
		}

		return err
	}
}
