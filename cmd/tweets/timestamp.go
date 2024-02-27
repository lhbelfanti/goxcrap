package tweets

import (
	"fmt"

	"github.com/tebeka/selenium"
)

const timestampXPath string = "div/div/div[2]/div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[3]/a/time"

// GetTimestamp retrieves the tweet timestamp from the datetime attribute of the time element.
// It will only be used to create a unique ID for the tweet
type GetTimestamp func(tweetArticleElement selenium.WebElement) (string, error)

// MakeGetTimestamp retrieves the tweet timestamp from the datetime attribute of the time element
func MakeGetTimestamp() GetTimestamp {
	return func(tweetArticleElement selenium.WebElement) (string, error) {
		tweetTimestampElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(timestampXPath))
		if err != nil {
			fmt.Println("Error finding tweet timestamp element:", err)
			return "", NewTweetsError(FailedToObtainTweetTimestampElement, err)
		}

		tweetTimestamp, err := tweetTimestampElement.GetAttribute("datetime")
		if err != nil {
			return "", NewTweetsError(FailedToObtainTweetTimestamp, err)
		}

		return tweetTimestamp, nil
	}
}
