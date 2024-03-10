package tweets

import (
	"log/slog"

	"github.com/tebeka/selenium"
)

const authorXPath string = "div[1]/div/div[1]/div/div/div[2]/div/div[1]/a/div/span"

// GetAuthor retrieves the tweet author
// It will only be used to create a unique ID for the tweet
type GetAuthor func(tweetArticleElement selenium.WebElement) (string, error)

// MakeGetAuthor creates a new GetAuthor
func MakeGetAuthor() GetAuthor {
	return func(tweetArticleElement selenium.WebElement) (string, error) {
		tweetAuthorElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(authorXPath))
		if err != nil {
			slog.Error(err.Error())
			return "", FailedToObtainTweetAuthorElement
		}

		tweetAuthor, err := tweetAuthorElement.Text()
		if err != nil {
			slog.Error(err.Error())
			return "", FailedToObtainTweetAuthor
		}

		return tweetAuthor, nil
	}
}
