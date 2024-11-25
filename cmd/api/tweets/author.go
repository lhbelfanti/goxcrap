package tweets

import (
	"context"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const authorXPath string = "div[1]/div/div[1]/div/div/div[2]/div/div[1]/a/div/span"

// GetAuthor retrieves the tweet author
type GetAuthor func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)

// MakeGetAuthor creates a new GetAuthor
func MakeGetAuthor() GetAuthor {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetAuthorElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(authorXPath))
		if err != nil {
			log.Error(ctx, err.Error())
			return "", FailedToObtainTweetAuthorElement
		}

		tweetAuthor, err := tweetAuthorElement.Text()
		if err != nil {
			log.Error(ctx, err.Error())
			return "", FailedToObtainTweetAuthor
		}

		return tweetAuthor, nil
	}
}
