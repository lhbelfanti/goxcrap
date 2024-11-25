package tweets

import (
	"context"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const avatarXPath string = "div[1]/div/div[1]/div/div/div[2]/div/div[1]/a/div/span"

// GetAvatar retrieves the tweet author's avatar
type GetAvatar func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error)

// MakeGetAvatar creates a new GetAvatar
func MakeGetAvatar() GetAuthor {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		tweetAvatarElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(avatarXPath))
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
