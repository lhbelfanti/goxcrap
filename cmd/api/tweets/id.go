package tweets

import (
	"context"
	"strings"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const (
	tweetIDXPath string = "div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[3]"

	tweetIDElementFromHeaderXPath string = "div[2]/div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[3]"
	tweetIDElementFromFooterXPath string = "div[3]/div[4]/div/div[1]/div/div[1]"
)

type (
	// GetID retrieves the ID of the tweet
	GetID func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)

	// GetIDFromTweetPage retrieves the ID of a tweet from the opened tweet page
	GetIDFromTweetPage func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)
)

// MakeGetID creates a new GetID
func MakeGetID() GetID {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetIDElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(tweetIDXPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetIDElement
		}

		return obtainTweetIDFromTweet(ctx, tweetIDElement)
	}
}

// MakeGetIDFromTweetPage creates a new GetIDFromTweetPage
func MakeGetIDFromTweetPage() GetIDFromTweetPage {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetIDElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(tweetIDElementFromHeaderXPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			log.Debug(ctx, FailedToObtainTweetIDElementFromTweetPageHeader.Error())

			tweetIDElement, err = tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(tweetIDElementFromFooterXPath))
			if err != nil {
				log.Warn(ctx, err.Error())
				log.Debug(ctx, FailedToObtainTweetIDElementFromTweetPageFooter.Error())
				return "", FailedToObtainTweetIDElementFromTweetPage
			}
		}

		return obtainTweetIDFromTweet(ctx, tweetIDElement)
	}
}

// obtainTweetIDFromTweet retrieves the tweet id from the given tweet element
func obtainTweetIDFromTweet(ctx context.Context, tweetIDElement selenium.WebElement) (string, error) {
	tweetIDATag, err := tweetIDElement.FindElement(selenium.ByTagName, "a")
	if err != nil {
		log.Warn(ctx, err.Error())
		return "", FailedToObtainTweetIDATag
	}

	tweetIDHref, err := tweetIDATag.GetAttribute("href")
	if err != nil {
		log.Warn(ctx, err.Error())
		return "", FailedToObtainTweetIDATagHref
	}

	hrefSplit := strings.Split(tweetIDHref, "/")
	tweetID := hrefSplit[len(hrefSplit)-1]

	return tweetID, nil
}
