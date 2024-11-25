package tweets

import (
	"context"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const (
	tweetOnlyTextXPath      string = "div[2]/div[3]/div[1]/div[1]/span"
	replyTweetOnlyTextXPath string = "div[2]/div[4]/div[1]/div[1]/span"

	tweetImagesXPath      string = "div[2]/div[3]/div[1]/div/div/div/div"
	replyTweetImagesXPath string = "div[2]/div[4]/div[1]/div/div/div/div"

	tweetIsReplyHasOnlyTextQuoteIsReplyImagesXPath      string = "div[2]/div[4]/div/div[2]/div/div[3]"
	tweetIsReplyHasTextAndImagesQuoteIsReplyImagesXPath string = "div[2]/div[4]/div[2]/div[2]/div/div[2]/div[1]"

	tweetIsNotReplyHasOnlyTextQuoteIsReplyImagesXPath      string = "div[2]/div[3]/div/div[2]/div/div[3]"
	tweetIsNotReplyHasTextAndImagesQuoteIsReplyImagesXPath string = "div[2]/div[3]/div[2]/div[2]/div/div[2]/div[1]"
)

type (
	// GetImages retrieves the tweet images
	GetImages func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply bool) ([]string, error)

	// GetQuoteImages retrieves the quoted tweet images
	GetQuoteImages func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) ([]string, error)
)

// MakeGetImages creates a new GetImages
func MakeGetImages() GetImages {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply bool) ([]string, error) {
		xPath := tweetOnlyTextXPath
		if isAReply {
			xPath = replyTweetOnlyTextXPath
		}

		// Pre-check, before accessing to the images
		_, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err == nil {
			// This tweet does not contain images
			return nil, FailedToObtainTweetImagesElement
		}

		xPath = tweetImagesXPath
		if isAReply {
			xPath = replyTweetImagesXPath
		}

		tweetImagesElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return nil, FailedToObtainTweetImagesElement
		}

		return obtainImagesFromTweet(ctx, tweetImagesElement, FailedToObtainTweetImages, FailedToObtainTweetSrcFromImage)
	}
}

// MakeGetQuoteImages creates a new GetQuoteImages
func MakeGetQuoteImages() GetQuoteImages {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) ([]string, error) {
		var xPath string
		if isAReply {
			if hasTweetOnlyText {
				xPath = tweetIsReplyHasOnlyTextQuoteIsReplyImagesXPath
			} else {
				xPath = tweetIsReplyHasTextAndImagesQuoteIsReplyImagesXPath
			}
		} else {
			if hasTweetOnlyText {
				xPath = tweetIsNotReplyHasOnlyTextQuoteIsReplyImagesXPath
			} else {
				xPath = tweetIsNotReplyHasTextAndImagesQuoteIsReplyImagesXPath
			}
		}

		tweetImagesElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			log.Warn(ctx, err.Error())
			return nil, FailedToObtainQuotedTweetImagesElement
		}

		return obtainImagesFromTweet(ctx, tweetImagesElement, FailedToObtainQuotedTweetImages, FailedToObtainQuotedTweetSrcFromImage)
	}
}

// obtainImagesFromTweet retrieves the images from the given tweet images element
func obtainImagesFromTweet(ctx context.Context, tweetImagesElement selenium.WebElement, failedToObtainTweetImages, failedToObtainTweetSrcFromImage error) ([]string, error) {
	tweetImagesElements, err := tweetImagesElement.FindElements(selenium.ByTagName, "img")
	if err != nil {
		log.Warn(ctx, err.Error())
		return nil, failedToObtainTweetImages
	}

	tweetImages := make([]string, 0, len(tweetImagesElements))
	for _, tweetImage := range tweetImagesElements {
		tweetUrl, err := tweetImage.GetAttribute("src")
		if err != nil {
			continue
		}

		tweetImages = append(tweetImages, tweetUrl)
	}

	if len(tweetImagesElements) > 0 && len(tweetImages) == 0 {
		return nil, failedToObtainTweetSrcFromImage
	}

	return tweetImages, nil
}
