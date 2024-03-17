package tweets

import (
	"log/slog"

	"github.com/tebeka/selenium"
)

const (
	tweetTextXPath      string = "div[2]/div"
	replyTweetTextXPath string = "div[3]/div"

	tweetIsReplyHasOnlyTextQuoteIsReplyXPath         string = "div[4]/div/div[2]/div/div[2]/div[2]"
	tweetIsReplyHasOnlyImagesQuoteIsReplyXPath       string = "div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsReplyHasTextAndImagesQuoteIsReplyXPath    string = "div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsReplyHasOnlyTextQuoteIsNotReplyXPath      string = "div[4]/div/div[2]/div/div[2]/div"
	tweetIsReplyHasOnlyImagesQuoteIsNotReplyXPath    string = "div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div"
	tweetIsReplyHasTextAndImagesQuoteIsNotReplyXPath string = "div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div"

	tweetIsNotReplyHasOnlyTextQuoteIsReplyXPath         string = "div[3]/div/div[2]/div/div[2]/div[2]"
	tweetIsNotReplyHasOnlyImagesQuoteIsReplyXPath       string = "div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsNotReplyHasTextAndImagesQuoteIsReplyXPath    string = "div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
	tweetIsNotReplyHasOnlyTextQuoteIsNotReplyXPath      string = "div[3]/div/div[2]/div/div[2]/div"
	tweetIsNotReplyHasOnlyImagesQuoteIsNotReplyXPath    string = "div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div"
	tweetIsNotReplyHasTextAndImagesQuoteIsNotReplyXPath string = "div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div"
)

type (
	// GetText retrieves the tweet text
	GetText func(tweetArticleElement selenium.WebElement, isAReply bool) (string, error)

	// GetQuoteText retrieves the quoted tweet text
	GetQuoteText func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText, hasTweetOnlyImages, isQuoteAReply bool) (string, error)
)

// MakeGetText creates a new GetText
func MakeGetText() GetText {
	return func(tweetArticleElement selenium.WebElement, isAReply bool) (string, error) {
		xPath := tweetTextXPath
		if isAReply {
			xPath = replyTweetTextXPath
		}

		tweetTextElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			slog.Error(err.Error())
			return "", FailedToObtainTweetTextElement
		}

		return obtainTextFromTweet(tweetTextElement, FailedToObtainTweetTextParts, FailedToObtainTweetTextPartTagName, FailedToObtainTweetTextFromSpan)
	}
}

// MakeGetQuoteText creates a new GetQuoteText
func MakeGetQuoteText() GetQuoteText {
	return func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText, hasTweetOnlyImages, isQuoteAReply bool) (string, error) {
		var xPath string
		if isAReply {
			if isQuoteAReply {
				if hasTweetOnlyText {
					xPath = tweetIsReplyHasOnlyTextQuoteIsReplyXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsReplyHasOnlyImagesQuoteIsReplyXPath
				} else {
					xPath = tweetIsReplyHasTextAndImagesQuoteIsReplyXPath
				}
			} else {
				if hasTweetOnlyText {
					xPath = tweetIsReplyHasOnlyTextQuoteIsNotReplyXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsReplyHasOnlyImagesQuoteIsNotReplyXPath
				} else {
					xPath = tweetIsReplyHasTextAndImagesQuoteIsNotReplyXPath
				}
			}
		} else {
			if isQuoteAReply {
				if hasTweetOnlyText {
					xPath = tweetIsNotReplyHasOnlyTextQuoteIsReplyXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsNotReplyHasOnlyImagesQuoteIsReplyXPath
				} else {
					xPath = tweetIsNotReplyHasTextAndImagesQuoteIsReplyXPath
				}
			} else {
				if hasTweetOnlyText {
					xPath = tweetIsNotReplyHasOnlyTextQuoteIsNotReplyXPath
				} else if hasTweetOnlyImages {
					xPath = tweetIsNotReplyHasOnlyImagesQuoteIsNotReplyXPath
				} else {
					xPath = tweetIsNotReplyHasTextAndImagesQuoteIsNotReplyXPath
				}
			}
		}

		tweetTextElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err != nil {
			slog.Error(err.Error())
			return "", FailedToObtainQuotedTweetTextElement
		}

		return obtainTextFromTweet(tweetTextElement, FailedToObtainQuotedTweetTextParts, FailedToObtainQuotedTweetTextPartTagName, FailedToObtainQuotedTweetTextFromSpan)
	}
}

// obtainTextFromTweet retrieves the text from the given tweet text element
func obtainTextFromTweet(tweetTextElement selenium.WebElement, failedToObtainTextParts, failedToObtainTextPartTagName, failedToObtainTextFromSpan error) (string, error) {
	textParts, err := tweetTextElement.FindElements(selenium.ByCSSSelector, "span, img")
	if err != nil {
		slog.Error(err.Error())
		return "", failedToObtainTextParts
	}

	var tweetText string
	for _, textPart := range textParts {
		tagName, err := textPart.TagName()
		if err != nil {
			slog.Error(err.Error())
			return "", failedToObtainTextPartTagName
		}

		switch tagName {
		case "span":
			spanText, err := textPart.Text()
			if err != nil {
				slog.Error(err.Error())
				return "", failedToObtainTextFromSpan
			}
			tweetText += spanText
		case "img":
			alt, err := textPart.GetAttribute("alt")
			if err != nil {
				slog.Error("Ignoring emoji: " + err.Error())
				continue
			}

			tweetText += alt
		}
	}
	return tweetText, nil
}
