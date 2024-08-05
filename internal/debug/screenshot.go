package debug

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tebeka/selenium"
)

// TakeScreenshot takes a screenshot of the current state
type TakeScreenshot func() error

// MakeTakeScreenshot creates a new TakeScreenshot
func MakeTakeScreenshot(wd selenium.WebDriver) TakeScreenshot {
	return func() error {
		screenshotBytes, err := wd.Screenshot()
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToTakeScreenshot
		}

		img, _, _ := image.Decode(bytes.NewReader(screenshotBytes))
		screenshotName := fmt.Sprintf("./goxcrap-%v.png", time.Now().Format(time.RFC822))
		out, err := os.Create(screenshotName)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToCreateScreenshotFile
		}

		err = png.Encode(out, img)
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToSaveScreenshotFile
		}

		return err
	}
}
