package tweets

import (
	"context"
	"strings"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const tweetIDXPath string = "div[2]/div[1]/div/div[1]/div/div/div[2]/div/div[3]"

type GetID func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)

// MakeGetID creates a new GetID
func MakeGetID() GetID {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetIDElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(tweetIDXPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return "", FailedToObtainTweetIDElement
		}

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
}
