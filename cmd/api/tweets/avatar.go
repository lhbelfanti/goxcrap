package tweets

import (
	"context"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const (
	tweetAvatarXPath string = "div[1]/div"

	quoteAvatarTweetHasOnlyTextXPath     string = "/div[2]/div[3]/div/div[2]/div/div[1]/div/div/div/div[1]"
	quoteAvatarTweetHasTextAndImageXPath string = "/div[2]/div[3]/div[2]/div[2]/div/div[1]/div/div/div/div[1]"
)

type (
	// GetAvatar retrieves the tweet author's avatar
	GetAvatar func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)

	// GetQuoteAvatar retrieves the quote author's avatar
	GetQuoteAvatar func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error)
)

// MakeGetAvatar creates a new GetAvatar
func MakeGetAvatar() GetAvatar {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetAvatarElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(tweetAvatarXPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetAvatarElement
		}

		tweetAvatarImage, err := tweetAvatarElement.FindElement(selenium.ByTagName, "img")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetAvatarImage
		}

		tweetAvatarURL, err := tweetAvatarImage.GetAttribute("src")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetAvatarSrcFromImage
		}

		return tweetAvatarURL, nil
	}
}

// MakeGetQuoteAvatar creates a new GetQuoteAvatar
func MakeGetQuoteAvatar() GetQuoteAvatar {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error) {
		xPath := quoteAvatarTweetHasTextAndImageXPath
		if hasTweetOnlyText {
			xPath = quoteAvatarTweetHasOnlyTextXPath
		}

		quoteAvatarElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainQuotedTweetAvatarElement
		}

		quoteAvatarImage, err := quoteAvatarElement.FindElement(selenium.ByTagName, "img")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainQuotedTweetAvatarImage
		}

		quoteAvatarURL, err := quoteAvatarImage.GetAttribute("src")
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainQuotedTweetAvatarSrcFromImage
		}

		return quoteAvatarURL, nil
	}
}
