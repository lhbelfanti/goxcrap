package tweets

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/tebeka/selenium"

	"goxcrap/internal/log"
)

type (
	// GetTweetHash retrieves the tweet timestamp and hash from the given tweet element
	GetTweetHash func(ctx context.Context, tweetArticleElement selenium.WebElement) (TweetHash, error)

	// GetTweetInformation retrieves the tweet information from the given tweet element
	GetTweetInformation func(ctx context.Context, tweetArticleElement selenium.WebElement, tweetHash TweetHash) (Tweet, error)
)

// MakeGetTweetHash creates a new GetTweetHash
func MakeGetTweetHash(getAuthor GetAuthor, getTimestamp GetTimestamp) GetTweetHash {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (TweetHash, error) {
		tweetAuthor, err := getAuthor(ctx, tweetArticleElement)
		if err != nil {
			log.Warn(ctx, err.Error())
			return TweetHash{}, FailedToObtainTweetAuthorInformation
		}

		tweetTimestamp, err := getTimestamp(ctx, tweetArticleElement)
		if err != nil {
			log.Warn(ctx, err.Error())
			return TweetHash{}, FailedToObtainTweetTimestampInformation
		}

		tweetAuthorHash := md5.Sum([]byte(tweetAuthor))
		tweetTimestampHash := md5.Sum([]byte(tweetTimestamp))
		tweetID := hex.EncodeToString(tweetAuthorHash[:]) + hex.EncodeToString(tweetTimestampHash[:])

		return TweetHash{
			ID:        tweetID,
			Author:    tweetAuthor,
			Timestamp: tweetTimestamp,
		}, nil
	}
}

// MakeGetTweetInformation creates a new GetTweetInformation
func MakeGetTweetInformation(isAReply IsAReply, getText GetText, getImages GetImages, hasQuote HasQuote, isQuoteAReply IsQuoteAReply, getQuoteText GetQuoteText, getQuoteImages GetQuoteImages) GetTweetInformation {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, tweetHash TweetHash) (Tweet, error) {
		isTweetAReply := isAReply(tweetArticleElement)

		tweetText, err := getText(ctx, tweetArticleElement, isTweetAReply)
		hasText := !errors.Is(err, FailedToObtainTweetTextElement)
		if err != nil && hasText {
			log.Debug(ctx, err.Error())
		}

		tweetImages, err := getImages(ctx, tweetArticleElement, isTweetAReply)
		hasImages := !errors.Is(err, FailedToObtainTweetImagesElement)
		if err != nil && hasImages {
			log.Debug(ctx, err.Error())
		}

		tweetOnlyHasText := hasText && !hasImages
		tweetOnlyHasImages := !hasText && hasImages

		hasAQuote := hasQuote(tweetArticleElement, isTweetAReply, tweetOnlyHasText)

		var quote Quote
		if hasAQuote {
			isQuotedTweetAReply := isQuoteAReply(tweetArticleElement, isTweetAReply, tweetOnlyHasText)

			quoteText, err := getQuoteText(ctx, tweetArticleElement, isTweetAReply, tweetOnlyHasText, tweetOnlyHasImages, isQuotedTweetAReply)
			if err != nil {
				log.Debug(ctx, err.Error())
			}
			hasQuotedTweetText := !errors.Is(err, FailedToObtainQuotedTweetTextElement)

			quoteImages, err := getQuoteImages(ctx, tweetArticleElement, isTweetAReply, tweetOnlyHasText)
			if err != nil {
				log.Debug(ctx, err.Error())
			}
			hasQuotedTweetImages := !errors.Is(err, FailedToObtainQuotedTweetImagesElement)

			quote = Quote{
				Data: Data{
					Author:    "", // TODO: Complete it
					Avatar:    "", // TODO: Complete it
					Timestamp: "", // TODO: Complete it
					IsAReply:  isQuotedTweetAReply,
					HasText:   hasQuotedTweetText,
					HasImages: hasQuotedTweetImages,
					Text:      quoteText,
					Images:    quoteImages,
				},
			}
		}

		return Tweet{
			ID:       tweetHash.ID,
			HasQuote: hasAQuote,
			Data: Data{
				Author:    tweetHash.Author,
				Avatar:    "", // TODO: Complete it
				Timestamp: tweetHash.Timestamp,
				IsAReply:  isTweetAReply,
				HasText:   hasText,
				HasImages: hasImages,
				Text:      tweetText,
				Images:    tweetImages,
			},
			Quote: quote,
		}, nil
	}
}
