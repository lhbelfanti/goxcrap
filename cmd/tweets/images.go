package tweets

import (
	"log/slog"

	"github.com/tebeka/selenium"
)

const (
	tweetImagesXPath      string = "div/div/div[2]/div[2]/div[3]/div[1]/div/div/div/div"
	replyTweetImagesXPath string = "div/div/div[2]/div[2]/div[4]/div[1]"
)

// GetImages retrieves the tweet images
type GetImages func(tweetArticleElement selenium.WebElement, isAReply bool) ([]string, error)

// MakeGetImages creates a new GetImages
func MakeGetImages() GetImages {
	return func(tweetArticleElement selenium.WebElement, isAReply bool) ([]string, error) {
		xPath := tweetImagesXPath
		if isAReply {
			xPath = replyTweetImagesXPath
		}

		tweetImagesElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			slog.Error(err.Error())
			return nil, FailedToObtainTweetImagesElement
		}

		tweetImagesElements, err := tweetImagesElement.FindElements(selenium.ByTagName, "img")
		if err != nil {
			slog.Error(err.Error())
			return nil, FailedToObtainTweetImages
		}

		tweetImages := make([]string, 0, len(tweetImagesElements))
		for _, tweetImage := range tweetImagesElements {
			tweetUrl, err := tweetImage.GetAttribute("src")
			if err != nil {
				continue
			}

			tweetImages = append(tweetImages, tweetUrl)
		}

		if len(tweetImagesElements) > 0 && len(tweetImages) == 0 {
			return nil, FailedToObtainTweetSrcFromImage
		}

		return tweetImages, nil
	}
}
