package tweets

import (
	"log/slog"

	"github.com/tebeka/selenium"
)

const (
	tweetTextXPath      string = "div[2]/div"
	replyTweetTextXPath string = "div[3]/div"

	tweetOnlyTextQuotedTweetTextXPath      string = "div[3]/div/div[2]/div/div[2]"
	tweetQuotedTweetTextXPath              string = "div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div"
	replyTweetOnlyTextQuotedTweetTextXPath string = "div[3]/div/div[2]/div/div[2]/div[2]"
	replyTweetQuotedTweetTextXPath         string = "div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div[2]"
)

type (
	// GetText retrieves the tweet text
	GetText func(tweetArticleElement selenium.WebElement, isAReply bool) (string, error)

	// GetQuoteText retrieves the quoted tweet text
	GetQuoteText func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) (string, error)
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
	return func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) (string, error) {
		var xPath string
		if isAReply {
			xPath = replyTweetQuotedTweetTextXPath
			if hasTweetOnlyText {
				xPath = replyTweetOnlyTextQuotedTweetTextXPath
			}
		} else {
			xPath = tweetQuotedTweetTextXPath
			if hasTweetOnlyText {
				xPath = tweetOnlyTextQuotedTweetTextXPath
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
