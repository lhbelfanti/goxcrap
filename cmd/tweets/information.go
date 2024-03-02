package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log/slog"

	"github.com/tebeka/selenium"
)

const replyXPath string = "div/div/div[2]/div[2]/div[2]/div"

// GatherTweetInformation retrieves the tweet information from the given tweet element
type GatherTweetInformation func(tweetArticleElement selenium.WebElement) (Tweet, error)

// MakeGetTweetInformation creates a new GatherTweetInformation
func MakeGetTweetInformation(getAuthor GetAuthor, getTimestamp GetTimestamp, getText GetText, getImages GetImages) GatherTweetInformation {
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

		tweetText, err := getText(tweetArticleElement, isAReply)
		if err != nil {
			slog.Error(err.Error())
		}
		hasText := !errors.Is(err, FailedToObtainTweetTextElement)

		tweetImages, err := getImages(tweetArticleElement, isAReply)
		if err != nil {
			slog.Error(err.Error())
		}
		hasImages := !errors.Is(err, FailedToObtainTweetImagesElement)

		tweetAuthorHash := md5.Sum([]byte(tweetAuthor))
		tweetTimestampHash := md5.Sum([]byte(tweetTimestamp))
		tweetID := hex.EncodeToString(tweetAuthorHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		return Tweet{
			ID:        tweetID,
			Timestamp: tweetTimestamp,
			IsAReply:  isAReply,
			HasQuote:  true,
			Data: Data{
				HasText:   hasText,
				HasImages: hasImages,
				Text:      tweetText,
				Images:    tweetImages,
			},
			Quote: Quote{
				IsAReply: true,
				Data: Data{
					HasText:   true,
					HasImages: true,
					Text:      "Quote Description",
					Images:    []string{"https://url3.com", "https://url4.com"},
				},
			},
		}, nil
	}
}
