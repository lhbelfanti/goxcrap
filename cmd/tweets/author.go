package tweets

import (
	"fmt"

	"github.com/tebeka/selenium"
)

const authorXPath string = "div/div/div[2]/div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[1]/a/div/span"

// GetAuthor retrieves the tweet timestamp from the datetime attribute of the time element
type GetAuthor func(tweetArticleElement selenium.WebElement) (string, error)

// MakeGetAuthor retrieves the tweet author. It will only be used to create a unique ID for the tweet
func MakeGetAuthor() GetAuthor {
	return func(tweetArticleElement selenium.WebElement) (string, error) {
		tweetAuthorElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(authorXPath))
		if err != nil {
			fmt.Println("Error finding tweet author element:", err)
			return "", NewTweetsError(FailedToObtainTweetAuthorElement, err)
		}

		tweetAuthor, err := tweetAuthorElement.Text()
		if err != nil {
			return "", NewTweetsError(FailedToObtainTweetAuthor, err)
		}

		return tweetAuthor, nil
	}
}
