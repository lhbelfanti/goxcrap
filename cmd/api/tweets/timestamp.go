package tweets

import (
	"github.com/rs/zerolog/log"
	"github.com/tebeka/selenium"
)

const timestampXPath string = "div[1]/div"

// GetTimestamp retrieves the tweet timestamp from the datetime attribute of the time element.
// It will only be used to create a unique ID for the tweet
type GetTimestamp func(tweetArticleElement selenium.WebElement) (string, error)

// MakeGetTimestamp creates a new GetTimestamp
func MakeGetTimestamp() GetTimestamp {
	return func(tweetArticleElement selenium.WebElement) (string, error) {
		tweetTimestampElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(timestampXPath))
		if err != nil {
			log.Info().Msg(err.Error())
			return "", FailedToObtainTweetTimestampElement
		}

		tweetTimestampTimeTag, err := tweetTimestampElement.FindElement(selenium.ByTagName, "time")
		if err != nil {
			log.Info().Msg(err.Error())
			return "", FailedToObtainTweetTimestampTimeTag
		}

		tweetTimestamp, err := tweetTimestampTimeTag.GetAttribute("datetime")
		if err != nil {
			log.Info().Msg(err.Error())
			return "", FailedToObtainTweetTimestamp
		}

		return tweetTimestamp, nil
	}
}
