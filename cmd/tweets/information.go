package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/tebeka/selenium"
)

const replyXPath string = "div/div/div[2]/div[2]/div[2]/div"

// GatherTweetInformation retrieves the tweet information from the given tweet element
type GatherTweetInformation func(tweetArticleElement selenium.WebElement) (Tweet, error)

// MakeGetTweetInformation creates a new GatherTweetInformation
func MakeGetTweetInformation(getTimestamp GetTimestamp, getAuthor GetAuthor) GatherTweetInformation {
	return func(tweetArticleElement selenium.WebElement) (Tweet, error) {
		tweetAuthor, err := getAuthor(tweetArticleElement)
		if err != nil {
			return Tweet{}, err
		}

		tweetTimestamp, err := getTimestamp(tweetArticleElement)
		if err != nil {
			return Tweet{}, err
		}

		isAReply := true
		_, err = tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(replyXPath))
		if err != nil {
			isAReply = false
		}

		tweetAuthorHash := md5.Sum([]byte(tweetAuthor))
		tweetTimestampHash := md5.Sum([]byte(tweetTimestamp))
		tweetID := hex.EncodeToString(tweetAuthorHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		return Tweet{
			ID:        tweetID,
			Timestamp: "",
			IsAReply:  isAReply,
			HasQuote:  false,
			Data: Data{
				HasText:   false,
				HasImages: false,
				Text:      "",
				Images:    nil,
			},
			Quote: Quote{
				IsAReply: false,
				Data: Data{
					HasText:   false,
					HasImages: false,
					Text:      "",
					Images:    nil,
				},
			},
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

func getTweetImages(tweetElement selenium.WebElement) ([]string, error) {
	tweetImagesElement, err := tweetElement.FindElement(selenium.ByXPATH, "div[position()=3]/div/div/div/div/div/div")
	if err != nil {
		fmt.Println("Error finding tweet images element:", err)
		return nil, NewTweetsError(FailedToObtainTweetImagesElement, err)
	}

	tweetImagesElements, err := tweetImagesElement.FindElements(selenium.ByTagName, "img")
	if err != nil {
		fmt.Println("Error finding tweet images:", err)
		return nil, NewTweetsError(FailedToObtainTweetImages, err)
	}

	tweetImages := make([]string, 0, len(tweetImagesElements))
	for _, tweetImage := range tweetImagesElements {
		tweetUrl, err := tweetImage.GetAttribute("src")
		if err != nil {
			return nil, NewTweetsError(FailedToObtainTweetSrcFromImage, err)
		}

		tweetImages = append(tweetImages, tweetUrl)
	}

	return tweetImages, nil
}
