package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"

	"github.com/tebeka/selenium"
)

// GatherTweetInformation retrieves the tweet information from the given tweet element
type GatherTweetInformation func(tweetArticleElement selenium.WebElement) (Tweet, error)

// MakeGetTweetInformation creates a new GatherTweetInformation
func MakeGetTweetInformation(getAuthor GetAuthor, getTimestamp GetTimestamp, isAReply IsAReply, getText GetText, getImages GetImages, hasQuote HasQuote, isQuoteAReply IsQuoteAReply) GatherTweetInformation {
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

		tweetAuthorHash := md5.Sum([]byte(tweetAuthor))
		tweetTimestampHash := md5.Sum([]byte(tweetTimestamp))
		tweetID := hex.EncodeToString(tweetAuthorHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		isTheTweetAReply := isAReply(tweetArticleElement)

		tweetText, err := getText(tweetArticleElement, isTheTweetAReply)
		if err != nil {
			slog.Error(err.Error())
		}
		hasTheTweetText := !errors.Is(err, FailedToObtainTweetTextElement)

		tweetImages, err := getImages(tweetArticleElement, isTheTweetAReply)
		if err != nil {
			slog.Error(err.Error())
		}
		hasTheTweetImages := !errors.Is(err, FailedToObtainTweetImagesElement)

		fmt.Printf("\n |------- \nAuthor: %s \nTimestamp: %s \nText: %s \nImages: %v \nIsAReply: %t \n", tweetAuthor, tweetTimestamp, tweetText, tweetImages, isTheTweetAReply)

		hasTheTweetOnlyText := hasTheTweetText && !hasTheTweetImages

		fmt.Printf("HasTheTweetOnlyText: %t\n", hasTheTweetOnlyText)

		hasTheTweetAQuote := hasQuote(tweetArticleElement, isTheTweetAReply, hasTheTweetOnlyText)

		fmt.Printf("HasQuote: %t\n", hasTheTweetAQuote)

		var quote Quote
		if hasTheTweetAQuote {
			//hasTheTweetOnlyImages := !hasTheTweetText && hasTheTweetImages
			//hasTheTweetTextAndImages := hasTheTweetText && hasTheTweetImages

			isQuotedTweetAReply := isQuoteAReply(tweetArticleElement, isTheTweetAReply, hasTheTweetOnlyText)

			fmt.Printf("IsQuoteAReply: %t -------|\n\n", isQuotedTweetAReply)

			// Gather Text, images with

			quote = Quote{
				IsAReply: isQuotedTweetAReply,
				Data:     Data{}, // TODO: complete this object
			}
		}

		return Tweet{
			ID:        tweetID,
			Timestamp: tweetTimestamp,
			IsAReply:  isTheTweetAReply,
			HasQuote:  hasTheTweetAQuote,
			Data: Data{
				HasText:   hasTheTweetText,
				HasImages: hasTheTweetImages,
				Text:      tweetText,
				Images:    tweetImages,
			},
			Quote: quote,
		}, nil
	}
}
