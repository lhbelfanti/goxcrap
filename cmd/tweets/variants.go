package tweets

import (
	"strings"

	"github.com/tebeka/selenium"
)

const (
	replyXPath string = "div/div/div[2]/div[2]/div[2]/div"

	quoteXPath           string = "div/div/div[2]/div[2]/div[3]/article/div/div/div/div/div/div/div/span"
	replyTweetQuoteXPath string = "div/div/div[2]/div[2]/div[4]/article/div/div/div/div/div/div/div/span"
)

type (
	// IsAReply returns a bool indicating if the base tweet is replying to another tweet
	IsAReply func(tweetArticleElement selenium.WebElement) bool

	// HasQuote returns a bool indicating if the base tweet is quoting another tweet
	HasQuote func(tweetArticleElement selenium.WebElement, isAReply bool) bool
)

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

// MakeHasQuote creates a new HasQuote
func MakeHasQuote() HasQuote {
	return func(tweetArticleElement selenium.WebElement, isAReply bool) bool {
		xPath := quoteXPath
		if isAReply {
			xPath = replyTweetQuoteXPath
		}

		_, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))

		return err == nil
	}
}
