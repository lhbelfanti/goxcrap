package tweets

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log/slog"

	"github.com/tebeka/selenium"
)

type (
	// GetTweetHash retrieves the tweet timestamp and hash from the given tweet element
	GetTweetHash func(tweetArticleElement selenium.WebElement) (TweetHash, error)

	// GetTweetInformation retrieves the tweet information from the given tweet element
	GetTweetInformation func(tweetArticleElement selenium.WebElement, tweetID, tweetTimestamp string) (Tweet, error)
)

// MakeGetTweetHash creates a new GetTweetHash
func MakeGetTweetHash(getAuthor GetAuthor, getTimestamp GetTimestamp) GetTweetHash {
	return func(tweetArticleElement selenium.WebElement) (TweetHash, error) {
		tweetAuthor, err := getAuthor(tweetArticleElement)
		if err != nil {
			slog.Error(err.Error())
			return TweetHash{}, FailedToObtainTweetAuthorInformation
		}

		tweetTimestamp, err := getTimestamp(tweetArticleElement)
		if err != nil {
			slog.Error(err.Error())
			return TweetHash{}, FailedToObtainTweetTimestampInformation
		}

		tweetAuthorHash := md5.Sum([]byte(tweetAuthor))
		tweetTimestampHash := md5.Sum([]byte(tweetTimestamp))
		tweetID := hex.EncodeToString(tweetAuthorHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		return TweetHash{
			ID:        tweetID,
			Timestamp: tweetTimestamp,
		}, nil
	}
}

// MakeGetTweetInformation creates a new GetTweetInformation
func MakeGetTweetInformation(isAReply IsAReply, getText GetText, getImages GetImages, hasQuote HasQuote, isQuoteAReply IsQuoteAReply, getQuoteText GetQuoteText, getQuoteImages GetQuoteImages) GetTweetInformation {
	return func(tweetArticleElement selenium.WebElement, tweetID, tweetTimestamp string) (Tweet, error) {
		isTweetAReply := isAReply(tweetArticleElement)

		tweetText, err := getText(tweetArticleElement, isTweetAReply)
		hasText := !errors.Is(err, FailedToObtainTweetTextElement)
		if err != nil && hasText {
			slog.Info(err.Error())
		}

		tweetImages, err := getImages(tweetArticleElement, isTweetAReply)
		hasImages := !errors.Is(err, FailedToObtainTweetImagesElement)
		if err != nil && hasImages {
			slog.Info(err.Error())
		}

		tweetOnlyHasText := hasText && !hasImages
		tweetOnlyHasImages := !hasText && hasImages

		hasAQuote := hasQuote(tweetArticleElement, isTweetAReply, tweetOnlyHasText)

		var quote Quote
		if hasAQuote {
			isQuotedTweetAReply := isQuoteAReply(tweetArticleElement, isTweetAReply, tweetOnlyHasText)

			quoteText, err := getQuoteText(tweetArticleElement, isTweetAReply, tweetOnlyHasText, tweetOnlyHasImages, isQuotedTweetAReply)
			if err != nil {
				slog.Info(err.Error())
			}
			hasQuotedTweetText := !errors.Is(err, FailedToObtainQuotedTweetTextElement)

			quoteImages, err := getQuoteImages(tweetArticleElement, isTweetAReply, tweetOnlyHasText)
			if err != nil {
				slog.Info(err.Error())
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
