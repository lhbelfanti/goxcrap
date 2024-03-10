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
func MakeGetTweetInformation(getAuthor GetAuthor, getTimestamp GetTimestamp, isAReply IsAReply, getText GetText, getImages GetImages, hasQuote HasQuote, isQuoteAReply IsQuoteAReply, getQuoteText GetQuoteText) GatherTweetInformation {
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

		hasTheTweetOnlyText := hasTheTweetText && !hasTheTweetImages

		hasTheTweetAQuote := hasQuote(tweetArticleElement, isTheTweetAReply, hasTheTweetOnlyText)

		fmt.Printf("\n|------- \nAuthor: %s \nTimestamp: %s \nText: %s \nImages: %v \nIsAReply: %t \nHasTheTweetOnlyText: %t \nHasQuote: %t \n", tweetAuthor, tweetTimestamp, tweetText, tweetImages, isTheTweetAReply, hasTheTweetOnlyText, hasTheTweetAQuote)

		var quote Quote
		if hasTheTweetAQuote {
			isQuotedTweetAReply := isQuoteAReply(tweetArticleElement, isTheTweetAReply, hasTheTweetOnlyText)

			quoteText, err := getQuoteText(tweetArticleElement, isTheTweetAReply, hasTheTweetOnlyText)
			if err != nil {
				slog.Error(err.Error())
			}
			hasTheQuotedTweetText := !errors.Is(err, FailedToObtainQuotedTweetTextElement)

			fmt.Printf("IsQuoteAReply: %t \n quoteText: %s \n-------|\n\n", isQuotedTweetAReply, quoteText)

			// Gather images

			quote = Quote{
				IsAReply: isQuotedTweetAReply,
				Data: Data{
					HasText:   hasTheQuotedTweetText,
					HasImages: false,
					Text:      quoteText,
					Images:    nil,
				},
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
