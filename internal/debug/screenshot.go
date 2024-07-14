package debug

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log/slog"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

// TakeScreenshot takes a screenshot of the current state
type TakeScreenshot func() error

// MakeTakeScreenshot creates a new TakeScreenshot
func MakeTakeScreenshot(wd selenium.WebDriver) TakeScreenshot {
	return func() error {
		screenshotBytes, err := wd.Screenshot()
		if err != nil {
			slog.Error(err.Error())
			return FailedToTakeScreenshot
		}

		img, _, _ := image.Decode(bytes.NewReader(screenshotBytes))
		screenshotName := fmt.Sprintf("./twitter-%v.png", time.Now().Format(time.RFC822))
		out, err := os.Create(screenshotName)
		if err != nil {
			slog.Error(err.Error())
			return FailedToCreateScreenshotFile
		}

		err = png.Encode(out, img)
		if err != nil {
			slog.Error(err.Error())
			return FailedToSaveScreenshotFile
		}

		return err
	}
}
