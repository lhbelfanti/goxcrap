package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tebeka/selenium"
)

const (
	tweetTextElementCSSSelector string = "div > div > div:nth-child(2) > div:nth-child(2) > div:nth-child(2)"
)

// GetTweetInformation retrieves the tweet information from the given tweet element
type GetTweetInformation func(tweetArticleElement selenium.WebElement) (Tweet, error)

// MakeGetTweetInformation creates a new GetTweetInformation
func MakeGetTweetInformation() GetTweetInformation {
	return func(tweetArticleElement selenium.WebElement) (Tweet, error) {
		// Find the tweet element
		tweetElement, err := tweetArticleElement.FindElement(selenium.ByCSSSelector, tweetTextElementCSSSelector)
		if err != nil {
			fmt.Println("Error finding tweet element:", err)
			return Tweet{}, NewTweetsError(FailedToObtainTweetElement, err)
		}

		// Extract tweet text
		tweetTextElement, err := tweetElement.FindElement(selenium.ByCSSSelector, "div:nth-child(2) > div")
		if err != nil {
			fmt.Println("Error finding tweet text element:", err)
			return Tweet{}, NewTweetsError(FailedToObtainTweetTextElement, err)
		}
		tweetText, err := tweetTextElement.Text()
		if err != nil {
			fmt.Println("Error getting tweet text:", err)
			return Tweet{}, NewTweetsError(FailedToObtainTweetText, err)
		}
		fmt.Println("Tweet text:", tweetText)

		// Extract tweet timestamp
		tweetTimestampElement, err := tweetElement.FindElement(selenium.ByCSSSelector, "div[data-testid='tweet'] > div > div > div > div > div > div > div > div > div:nth-child(1) > div > div > div > div:nth-child(1) > div > div:nth-child(2) > div > div > a > time")
		if err != nil {
			fmt.Println("Error finding tweet timestamp element:", err)
			return Tweet{}, NewTweetsError(FailedToObtainTweetTimestampElement, err)
		}
		tweetTimestamp, err := tweetTimestampElement.Text()
		if err != nil {
			fmt.Println("Error getting tweet timestamp:", err)
			return Tweet{}, NewTweetsError(FailedToObtainTweetTimestamp, err)
		}

		fmt.Println("Tweet timestamp:", tweetTimestamp)

		// Check if there are images
		var tweetImages []string
		imagesExist, err := tweetElement.FindElement(selenium.ByCSSSelector, "div[data-testid='tweet'] > div > div > div > div > div > div > div > div:nth-child(2) > div > div > div > div > div > div > div > div:nth-child(2) > div > div")
		if err == nil {
			// Extract image URLs
			imageElements, err := imagesExist.FindElements(selenium.ByCSSSelector, "div > div > div > div:nth-child(1) > div > div > div > div > div > div > div > div")
			if err != nil {
				fmt.Println("Error finding image elements:", err)
			} else {
				for _, imageElement := range imageElements {
					imageURL, err := imageElement.GetAttribute("style")
					if err != nil {
						fmt.Println("Error getting image URL:", err)
					} else {
						urlStart := strings.Index(imageURL, "('") + 2
						urlEnd := strings.LastIndex(imageURL, "')")
						imageURL = imageURL[urlStart:urlEnd]
						tweetImages = append(tweetImages, imageURL)
					}
				}
			}
		} else {
			fmt.Println("No images found in the tweet.")
		}

		tweetTextHash := md5.Sum([]byte(tweetText))
		tweetTimestampHash := md5.Sum([]byte(tweetTimestamp))
		tweetID := hex.EncodeToString(tweetTextHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		return Tweet{
			ID:     tweetID,
			Text:   tweetText,
			Images: tweetImages,
		}, nil
	}
}
