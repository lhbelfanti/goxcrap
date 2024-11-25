package tweets

import (
	"context"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const (
	tweetTimestampXPath string = "div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[3]"

	quoteTimestampTweetHasOnlyTextXPath     = "div[2]/div[3]/div/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[3]/div/div"
	quoteTimestampTweetHasTextAndImageXPath = "div[2]/div[3]/div[2]/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[3]/div/div"
)

type (
	// GetTimestamp retrieves the tweet timestamp from the datetime attribute of the time element
	GetTimestamp func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)

	// GetQuoteTimestamp retrieves the quoted tweet timestamp from the datetime attribute of the time element
	GetQuoteTimestamp func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error)
)

// MakeGetTimestamp creates a new GetTimestamp
func MakeGetTimestamp() GetTimestamp {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetTimestampElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(tweetTimestampXPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetTimestampElement
		}

		tweetTimestampTimeTag, err := tweetTimestampElement.FindElement(selenium.ByTagName, "time")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetTimestampTimeTag
		}

		tweetTimestamp, err := tweetTimestampTimeTag.GetAttribute("datetime")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetTimestamp
		}

		return tweetTimestamp, nil
	}
}

// MakeGetQuoteTimestamp creates a new GetQuoteTimestamp
func MakeGetQuoteTimestamp() GetQuoteTimestamp {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error) {
		xPath := quoteTimestampTweetHasTextAndImageXPath
		if hasTweetOnlyText {
			xPath = quoteTimestampTweetHasOnlyTextXPath
		}

		quoteTimestampElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainQuotedTweetTimestampElement
		}

		quoteTimestampTimeTag, err := quoteTimestampElement.FindElement(selenium.ByTagName, "time")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainQuotedTweetTimestampTimeTag
		}

		quoteTimestamp, err := quoteTimestampTimeTag.GetAttribute("datetime")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainQuotedTweetTimestamp
		}

		return quoteTimestamp, nil
	}
}
