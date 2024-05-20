package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
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
		hasText := !errors.Is(err, FailedToObtainTweetTextElement)
		if err != nil && hasText {
			slog.Error(err.Error())
		}

		tweetImages, err := getImages(tweetArticleElement, isTweetAReply)
		hasImages := !errors.Is(err, FailedToObtainTweetImagesElement)
		if err != nil && hasImages {
			slog.Error(err.Error())
		}

		tweetOnlyHasText := hasText && !hasImages
		tweetOnlyHasImages := !hasText && hasImages

		hasAQuote := hasQuote(tweetArticleElement, isTweetAReply, tweetOnlyHasText)

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
