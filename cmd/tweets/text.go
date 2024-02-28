package tweets

import (
	"fmt"

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
			fmt.Println("Error finding tweet text element:", err)
			return "", NewTweetsError(FailedToObtainTweetTextElement, err)
		}

		textParts, err := tweetTextElement.FindElements(selenium.ByCSSSelector, "span, img")
		if err != nil {
			fmt.Println("Error finding text parts:", err)
			return "", NewTweetsError(FailedToObtainTweetTextParts, err)
		}

		var tweetText string
		for _, textPart := range textParts {
			tagName, err := textPart.TagName()
			if err != nil {
				fmt.Println("Error finding text part tag name:", err)
				return "", NewTweetsError(FailedToObtainTweetTextPartTagName, err)
			}

			switch tagName {
			case "span":
				spanText, err := textPart.Text()
				if err != nil {
					fmt.Println("Error getting tweet text from span:", err)
					return "", NewTweetsError(FailedToObtainTweetTextFromSpan, err)
				}
				tweetText += spanText
			case "img":
				alt, err := textPart.GetAttribute("alt")
				if err != nil {
					fmt.Println("Ignoring emoji. Error finding text part alt attribute", err)
					continue
				}

				tweetText += alt
			}
		}
		return tweetText, nil
	}
}
