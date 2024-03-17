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
func MakeGetTweetInformation(getAuthor GetAuthor, getTimestamp GetTimestamp, isAReply IsAReply, getText GetText, getImages GetImages, hasQuote HasQuote, isQuoteAReply IsQuoteAReply, getQuoteText GetQuoteText, getQuoteImages GetQuoteImages) GatherTweetInformation {
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

		isTweetAReply := isAReply(tweetArticleElement)

		tweetText, err := getText(tweetArticleElement, isTweetAReply)
		if err != nil {
			slog.Error(err.Error())
		}
		hasText := !errors.Is(err, FailedToObtainTweetTextElement)

		tweetImages, err := getImages(tweetArticleElement, isTweetAReply)
		if err != nil {
			slog.Error(err.Error())
		}
		hasImages := !errors.Is(err, FailedToObtainTweetImagesElement)

		tweetOnlyHasText := hasText && !hasImages
		tweetOnlyHasImages := !hasText && hasImages

		hasAQuote := hasQuote(tweetArticleElement, isTweetAReply, tweetOnlyHasText)

		fmt.Printf("\n|------- \n"+
			"Author: %s \n"+
			"Timestamp: %s \n"+
			"Text: %s \n"+
			"Images: %v \n"+
			"IsAReply: %t \n"+
			"HasTheTweetOnlyText: %t \n"+
			"HasTheTweetOnlyImages: %t \n"+
			"HasQuote: %t \n",
			tweetAuthor,
			tweetTimestamp,
			tweetText,
			tweetImages,
			isTweetAReply,
			tweetOnlyHasText,
			tweetOnlyHasImages,
			hasAQuote)

		var quote Quote
		if hasAQuote {
			isQuotedTweetAReply := isQuoteAReply(tweetArticleElement, isTweetAReply, tweetOnlyHasText)

			quoteText, err := getQuoteText(tweetArticleElement, isTweetAReply, tweetOnlyHasText, tweetOnlyHasImages, isQuotedTweetAReply)
			if err != nil {
				slog.Error(err.Error())
			}
			hasQuotedTweetText := !errors.Is(err, FailedToObtainQuotedTweetTextElement)

			quoteImages, err := getQuoteImages(tweetArticleElement, isTweetAReply, tweetOnlyHasText)
			if err != nil {
				slog.Error(err.Error())
			}
			hasQuotedTweetImages := !errors.Is(err, FailedToObtainQuotedTweetImagesElement)

			fmt.Printf("IsQuoteAReply: %t \nQuoteText: %s \nQuoteText: %v \n-------|\n\n", isQuotedTweetAReply, quoteText, quoteImages)

			quote = Quote{
				IsAReply: isQuotedTweetAReply,
				Data: Data{
					HasText:   hasQuotedTweetText,
					HasImages: hasQuotedTweetImages,
					Text:      quoteText,
					Images:    quoteImages,
				},
			}
		}

		return Tweet{
			ID:        tweetID,
			Timestamp: tweetTimestamp,
			IsAReply:  isTweetAReply,
			HasQuote:  hasAQuote,
			Data: Data{
				HasText:   hasText,
				HasImages: hasImages,
				Text:      tweetText,
				Images:    tweetImages,
			},
			Quote: quote,
		}, nil
	}
}
