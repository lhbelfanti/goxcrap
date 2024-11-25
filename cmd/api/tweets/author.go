package tweets

import (
	"context"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const (
	tweetAuthorXPath string = "div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[1]/a/div/span"

	quoteAuthorTweetHasOnlyTextXPath     = "div[2]/div[3]/div/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[1]/div/div/span"
	quoteAuthorTweetHasTextAndImageXPath = "div[2]/div[3]/div[2]/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[1]/div/div/span"
)

type (
	// GetAuthor retrieves the tweet author
	GetAuthor func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)

	// GetQuoteAuthor retrieves the quoted tweet author
	GetQuoteAuthor func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error)
)

// MakeGetAuthor creates a new GetAuthor
func MakeGetAuthor() GetAuthor {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetAuthorElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(tweetAuthorXPath))
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

// MakeGetQuoteAuthor creates a new GetQuoteAuthor
func MakeGetQuoteAuthor() GetQuoteAuthor {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error) {
		xPath := quoteAuthorTweetHasTextAndImageXPath
		if hasTweetOnlyText {
			xPath = quoteAuthorTweetHasOnlyTextXPath
		}

		quoteAuthorElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			log.Error(ctx, err.Error())
			return "", FailedToObtainQuotedTweetAuthorElement
		}

		quoteAuthor, err := quoteAuthorElement.Text()
		if err != nil {
			log.Error(ctx, err.Error())
			return "", FailedToObtainQuotedTweetAuthor
		}

		return quoteAuthor, nil
	}
}
