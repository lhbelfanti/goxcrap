package tweets

import (
	"strings"

	"github.com/tebeka/selenium"
)

const (
	replyXPath string = "div[2]/div[2]/div"

	quoteXPath                   string = "div[2]/div[3]/div[2]/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[3]"
	quoteOnlyTextXPath           string = "div[2]/div[3]/div/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[3]"
	replyTweetQuoteXPath         string = "div[2]/div[4]/div[2]/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[3]/div/div"
	replyTweetQuoteOnlyTextXPath string = "div[2]/div[4]/div/div[2]/div/div[1]/div/div/div/div[2]/div/div[2]/div/div[3]/div/div"

	replyTweetReplyQuoteXPath         string = "div[2]/div[4]/div[2]/div[2]/div/div[2]/div[2]/div/div[1]"
	replyTweetOnlyTextReplyQuoteXPath string = "div[2]/div[4]/div/div[2]/div/div[2]/div[1]"
	replyQuoteXPath                   string = "div[2]/div[3]/div[2]/div[2]/div/div[2]/div[2]/div/div[1]"
	tweetOnlyTextReplyQuoteXPath      string = "div[2]/div[3]/div/div[2]/div/div[2]/div[1]"
)

type (
	// IsAReply returns a bool indicating if the base tweet is replying to another tweet
	IsAReply func(tweetArticleElement selenium.WebElement) bool

	// HasQuote returns a bool indicating if the base tweet is quoting another tweet
	HasQuote func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) bool

	// IsQuoteAReply returns a bool indicating if the quoted tweet is replying to another tweet
	IsQuoteAReply func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) bool
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
	return func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) bool {
		var xPath string
		if isAReply {
			xPath = replyTweetQuoteXPath
			if hasTweetOnlyText {
				xPath = replyTweetQuoteOnlyTextXPath
			}
		} else {
			xPath = quoteXPath
			if hasTweetOnlyText {
				xPath = quoteOnlyTextXPath
			}
		}

		_, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		return err == nil
	}
}

// MakeIsQuoteAReply creates a new IsQuoteAReply
func MakeIsQuoteAReply() IsQuoteAReply {
	return func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) bool {
		var xPath string
		if isAReply {
			xPath = replyTweetReplyQuoteXPath
			if hasTweetOnlyText {
				xPath = replyTweetOnlyTextReplyQuoteXPath
			}
		} else {
			xPath = replyQuoteXPath
			if hasTweetOnlyText {
				xPath = tweetOnlyTextReplyQuoteXPath
			}
		}

		quoteReplyingToElement, err := tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(xPath))
		if err == nil {
			replyingToText, err := quoteReplyingToElement.Text()
			if err == nil && strings.Contains(replyingToText, "Replying to") {
				return true
			}
		}

		return false
	}
}
