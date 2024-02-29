package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log/slog"

	"github.com/tebeka/selenium"
)

const replyXPath string = "div/div/div[2]/div[2]/div[2]/div"

// GatherTweetInformation retrieves the tweet information from the given tweet element
type GatherTweetInformation func(tweetArticleElement selenium.WebElement) (Tweet, error)

// MakeGetTweetInformation creates a new GatherTweetInformation
func MakeGetTweetInformation(getAuthor GetAuthor, getTimestamp GetTimestamp, getText GetText) GatherTweetInformation {
	return func(tweetArticleElement selenium.WebElement) (Tweet, error) {
		tweetAuthor, err := getAuthor(tweetArticleElement)
		if err != nil {
			slog.Error(err.Error())
			return Tweet{}, FailedToObtainTweetAuthorInformation
		}

		tweetTimestamp, err := getTimestamp(tweetArticleElement)
		if err != nil {
			slog.Error(err.Error())
			return Tweet{}, FailedToObtainTweetTimestampInformation
		}

		_, err = tweetArticleElement.FindElement(selenium.ByXPATH, globalToLocalXPath(replyXPath))
		isAReply := err == nil

		// TODO: validate error to know if the element exists or not
		/*		tweetText, err := getText(tweetArticleElement, isAReply)
				if err != nil {
					return Tweet{}, err
				}*/

		tweetAuthorHash := md5.Sum([]byte(tweetAuthor))
		tweetTimestampHash := md5.Sum([]byte(tweetTimestamp))
		tweetID := hex.EncodeToString(tweetAuthorHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		return Tweet{
			ID:        tweetID,
			Timestamp: tweetTimestamp,
			IsAReply:  isAReply,
			HasQuote:  true,
			Data: Data{
				HasText:   true,
				HasImages: true,
				Text:      "Tweet Description",
				Images:    []string{"Img 1", "Img 2"},
			},
			Quote: Quote{
				IsAReply: true,
				Data: Data{
					HasText:   true,
					HasImages: true,
					Text:      "Quote Description",
					Images:    []string{"Img 3", "Img 4"},
				},
			},
		}, nil
	}
}

func getTweetImages(tweetElement selenium.WebElement) ([]string, error) {
	tweetImagesElement, err := tweetElement.FindElement(selenium.ByXPATH, "div[position()=3]/div/div/div/div/div/div")
	if err != nil {
		fmt.Println("Error finding tweet images element:", err)
		return nil, FailedToObtainTweetImagesElement
	}

	tweetImagesElements, err := tweetImagesElement.FindElements(selenium.ByTagName, "img")
	if err != nil {
		fmt.Println("Error finding tweet images:", err)
		return nil, FailedToObtainTweetImages
	}

	tweetImages := make([]string, 0, len(tweetImagesElements))
	for _, tweetImage := range tweetImagesElements {
		tweetUrl, err := tweetImage.GetAttribute("src")
		if err != nil {
			return nil, FailedToObtainTweetSrcFromImage
		}

		tweetImages = append(tweetImages, tweetUrl)
	}

	return tweetImages, nil
}
