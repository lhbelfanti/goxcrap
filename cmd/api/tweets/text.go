package tweets

import (
	"context"
	"fmt"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

const (
	tweetTextXPath      string = "div[2]/div[2]/div"
	replyTweetTextXPath string = "div[2]/div[3]/div"

	tweetIsReplyHasOnlyTextQuoteIsReplyTextXPath         string = "div[2]/div[4]/div/div[2]/div/div[2]/div[2]"
	tweetIsReplyHasOnlyImagesQuoteIsReplyTextXPath       string = "div[2]/div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsReplyHasTextAndImagesQuoteIsReplyTextXPath    string = "div[2]/div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsReplyHasOnlyTextQuoteIsNotReplyTextXPath      string = "div[2]/div[4]/div/div[2]/div/div[2]/div"
	tweetIsReplyHasOnlyImagesQuoteIsNotReplyTextXPath    string = "div[2]/div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div"
	tweetIsReplyHasTextAndImagesQuoteIsNotReplyTextXPath string = "div[2]/div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div"

	tweetIsNotReplyHasOnlyTextQuoteIsReplyTextXPath         string = "div[2]/div[3]/div/div[2]/div/div[2]/div[2]"
	tweetIsNotReplyHasOnlyImagesQuoteIsReplyTextXPath       string = "div[2]/div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsNotReplyHasTextAndImagesQuoteIsReplyTextXPath    string = "div[2]/div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsNotReplyHasOnlyTextQuoteIsNotReplyTextXPath      string = "div[2]/div[3]/div/div[2]/div/div[2]/div"
	tweetIsNotReplyHasOnlyImagesQuoteIsNotReplyTextXPath    string = "div[2]/div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div"
	tweetIsNotReplyHasTextAndImagesQuoteIsNotReplyTextXPath string = "div[2]/div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div"
)

type (
	// GetText retrieves the tweet text
	GetText func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply bool) (string, error)

	// GetQuoteText retrieves the quoted tweet text
	GetQuoteText func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText, hasTweetOnlyImages, isQuoteAReply bool) (string, error)
)

// MakeGetText creates a new GetText
func MakeGetText() GetText {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply bool) (string, error) {
		xPath := tweetTextXPath
		if isAReply {
			xPath = replyTweetTextXPath
		}

		tweetTextElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			// This tweet does not contain text
			return "", FailedToObtainTweetTextElement
		}

		return obtainTextFromTweet(ctx, tweetTextElement, FailedToObtainTweetTextParts, FailedToObtainTweetTextPartTagName, FailedToObtainTweetTextFromSpan)
	}
}

// MakeGetQuoteText creates a new GetQuoteText
func MakeGetQuoteText() GetQuoteText {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText, hasTweetOnlyImages, isQuoteAReply bool) (string, error) {
		var xPath string
		if isAReply {
			if isQuoteAReply {
				if hasTweetOnlyText {
					xPath = tweetIsReplyHasOnlyTextQuoteIsReplyTextXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsReplyHasOnlyImagesQuoteIsReplyTextXPath
				} else {
					xPath = tweetIsReplyHasTextAndImagesQuoteIsReplyTextXPath
				}
			} else {
				if hasTweetOnlyText {
					xPath = tweetIsReplyHasOnlyTextQuoteIsNotReplyTextXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsReplyHasOnlyImagesQuoteIsNotReplyTextXPath
				} else {
					xPath = tweetIsReplyHasTextAndImagesQuoteIsNotReplyTextXPath
				}
			}
		} else {
			if isQuoteAReply {
				if hasTweetOnlyText {
					xPath = tweetIsNotReplyHasOnlyTextQuoteIsReplyTextXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsNotReplyHasOnlyImagesQuoteIsReplyTextXPath
				} else {
					xPath = tweetIsNotReplyHasTextAndImagesQuoteIsReplyTextXPath
				}
			} else {
				if hasTweetOnlyText {
					xPath = tweetIsNotReplyHasOnlyTextQuoteIsNotReplyTextXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsNotReplyHasOnlyImagesQuoteIsNotReplyTextXPath
				} else {
					xPath = tweetIsNotReplyHasTextAndImagesQuoteIsNotReplyTextXPath
				}
			}
		}

		tweetTextElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			// This quoted tweet does not contain text
			return "", FailedToObtainQuotedTweetTextElement
		}

		return obtainTextFromTweet(ctx, tweetTextElement, FailedToObtainQuotedTweetTextParts, FailedToObtainQuotedTweetTextPartTagName, FailedToObtainQuotedTweetTextFromSpan)
	}
}

// obtainTextFromTweet retrieves the text from the given tweet text element
func obtainTextFromTweet(ctx context.Context, tweetTextElement selenium.WebElement, failedToObtainTextParts, failedToObtainTextPartTagName, failedToObtainTextFromSpan error) (string, error) {
	textParts, err := tweetTextElement.FindElements(selenium.ByCSSSelector, "span, img")
	if err != nil {
		log.Debug(ctx, err.Error())
		return "", failedToObtainTextParts
	}

	var tweetText string
	for _, textPart := range textParts {
		tagName, err := textPart.TagName()
		if err != nil {
			log.Debug(ctx, err.Error())
			return "", failedToObtainTextPartTagName
		}

		switch tagName {
		case "span":
			spanText, err := textPart.Text()
			if err != nil {
				log.Debug(ctx, err.Error())
				return "", failedToObtainTextFromSpan
			}
			tweetText += spanText
		case "img":
			alt, err := textPart.GetAttribute("alt")
			if err != nil {
				log.Debug(ctx, fmt.Sprintf("Ignoring emoji: %v", err.Error()))
				continue
			}

			tweetText += alt
		}
	}
	return tweetText, nil
}
