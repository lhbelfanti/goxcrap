package tweets

import (
	"log/slog"

	"github.com/tebeka/selenium"
)

const (
	normalTweetTextXPath string = "div/div/div[2]/div[2]/div[2]/div"
	replyTweetTextXPath  string = "div/div/div[2]/div[2]/div[3]/div"
)

// GetText retrieves the tweet text
type GetText func(tweetArticleElement selenium.WebElement, isAReply bool) (string, error)

// MakeGetText creates a new GetText
func MakeGetText() GetText {
	return func(tweetArticleElement selenium.WebElement, isAReply bool) (string, error) {
		xpath := normalTweetTextXPath
		if isAReply {
			xpath = replyTweetTextXPath
		}

		tweetTextElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xpath))
		if err != nil {
			slog.Error(err.Error())
			return "", FailedToObtainTweetTextElement
		}

		textParts, err := tweetTextElement.FindElements(selenium.ByCSSSelector, "span, img")
		if err != nil {
			slog.Error(err.Error())
			return "", FailedToObtainTweetTextParts
		}

		var tweetText string
		for _, textPart := range textParts {
			tagName, err := textPart.TagName()
			if err != nil {
				slog.Error(err.Error())
				return "", FailedToObtainTweetTextPartTagName
			}

			switch tagName {
			case "span":
				spanText, err := textPart.Text()
				if err != nil {
					slog.Error(err.Error())
					return "", FailedToObtainTweetTextFromSpan
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
}
