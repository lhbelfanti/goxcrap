package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/tebeka/selenium"
)

const (
	tweetElementCSSSelector string = "div > div > div:nth-child(2) > div:nth-child(2)"
)

// GetTweetInformation retrieves the tweet information from the given tweet element
type GetTweetInformation func(tweetArticleElement selenium.WebElement) (Tweet, error)

// MakeGetTweetInformation creates a new GetTweetInformation
func MakeGetTweetInformation() GetTweetInformation {
	return func(tweetArticleElement selenium.WebElement) (Tweet, error) {
		// Find the tweet element
		tweetElement, err := tweetArticleElement.FindElement(selenium.ByCSSSelector, tweetElementCSSSelector)
		if err != nil {
			fmt.Println("Error finding tweet element:", err)
			return Tweet{}, NewTweetsError(FailedToObtainTweetElement, err)
		}

		tweetText, err := getTweetText(tweetElement)
		if err != nil {
			return Tweet{}, err
		}

		// TODO: Get Timestamp
		// TODO: Get Images

		tweetTextHash := md5.Sum([]byte(tweetText))
		tweetTimestampHash := md5.Sum([]byte("tweetTimestamp"))
		tweetID := hex.EncodeToString(tweetTextHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		return Tweet{
			ID:   tweetID,
			Text: tweetText,
			//Images: tweetImages,
		}, nil
	}
}

// getTweetText retrieves the tweet text from the different elements inside the div
func getTweetText(tweetElement selenium.WebElement) (string, error) {
	tweetTextElement, err := tweetElement.FindElement(selenium.ByXPATH, "div[position()=2]/div")
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
