package tweets

import (
	"strings"

	"github.com/tebeka/selenium"
)

const replyXPath string = "div/div/div[2]/div[2]/div[2]/div"

// IsAReply returns a bool indicating if the base tweet is replying to another tweet
type IsAReply func(tweetArticleElement selenium.WebElement) bool

// MakeIsAReply creates a new GetIsAReply
func MakeIsAReply() IsAReply {
	return func(tweetArticleElement selenium.WebElement) bool {
		replyingToElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(replyXPath))
		if err == nil {
			replyingToText, err := replyingToElement.Text()
			if err == nil && strings.Contains(replyingToText, "Replying to") {
				return true
			}
		}

		return false
	}
}
