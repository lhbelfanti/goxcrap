package tweets

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/page"
	"goxcrap/internal/log"
)

// GetTweetInformation retrieves the tweet information from the given tweet element
type GetTweetInformation func(ctx context.Context, tweetArticleElement selenium.WebElement, tweetID string) (Tweet, error)

// MakeGetTweetInformation creates a new GetTweetInformation
func MakeGetTweetInformation(isAReply IsAReply, getAuthor GetAuthor, getTimestamp GetTimestamp, getAvatar GetAvatar, getText GetText, getImages GetImages, hasQuote HasQuote, isQuoteAReply IsQuoteAReply, getQuoteAuthor GetQuoteAuthor, getQuoteAvatar GetQuoteAvatar, getQuoteTimestamp GetQuoteTimestamp, getQuoteText GetQuoteText, getQuoteImages GetQuoteImages, loadPage page.Load, getLongText GetLongText, goBack page.GoBack) GetTweetInformation {
	pageLoaderTimeoutValue, _ := strconv.Atoi(os.Getenv("TWEET_PAGE_TIMEOUT"))
	pageLoaderTimeout := time.Duration(pageLoaderTimeoutValue) * time.Second

	return func(ctx context.Context, tweetArticleElement selenium.WebElement, tweetID string) (Tweet, error) {
		isTweetAReply := isAReply(tweetArticleElement)

		tweetAuthor, err := getAuthor(ctx, tweetArticleElement)
		if err != nil {
			log.Warn(ctx, err.Error())
			return Tweet{}, FailedToObtainTweetAuthorInformation
		}

		tweetTimestamp, err := getTimestamp(ctx, tweetArticleElement)
		if err != nil {
			log.Warn(ctx, err.Error())
			return Tweet{}, FailedToObtainTweetTimestampInformation
		}

		tweetAvatar, err := getAvatar(ctx, tweetArticleElement)
		if err != nil {
			log.Debug(ctx, err.Error())
		}

		tweetText, hasLongText, err := getText(ctx, tweetArticleElement, isTweetAReply)
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

			quoteAuthor, err := getQuoteAuthor(ctx, tweetArticleElement, tweetOnlyHasText)
			if err != nil {
				log.Debug(ctx, err.Error())
			}

			quoteAvatar, err := getQuoteAvatar(ctx, tweetArticleElement, tweetOnlyHasText)
			if err != nil {
				log.Debug(ctx, err.Error())
			}

			quoteTimestamp, err := getQuoteTimestamp(ctx, tweetArticleElement, tweetOnlyHasText)
			if err != nil {
				log.Debug(ctx, err.Error())
			}

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
					Author:    quoteAuthor,
					Avatar:    quoteAvatar,
					Timestamp: quoteTimestamp,
					IsAReply:  isQuotedTweetAReply,
					HasText:   hasQuotedTweetText,
					HasImages: hasQuotedTweetImages,
					Text:      quoteText,
					Images:    quoteImages,
				},
			}
		}

		tweet := Tweet{
			ID:       tweetID,
			HasQuote: hasAQuote,
			Data: Data{
				Author:    tweetAuthor,
				Avatar:    tweetAvatar,
				Timestamp: tweetTimestamp,
				IsAReply:  isTweetAReply,
				HasText:   hasText,
				HasImages: hasImages,
				Text:      tweetText,
				Images:    tweetImages,
			},
			Quote: quote,
		}

		// As obtaining the long text requires to load a new page and then go back to the previous one, it is called last.
		// Worst case scenario, the long text can't be obtained but the rest of the information was already retrieved.
		if hasLongText {
			url := "/" + tweetAuthor + "/status/" + tweetID

			err = loadPage(ctx, url, pageLoaderTimeout)
			if err != nil {
				log.Error(ctx, err.Error())
				return tweet, FailedToLoadTweetLongTextPage
			}

			longText, err := getLongText(ctx, isTweetAReply)
			if err != nil {
				log.Debug(ctx, err.Error())
			} else {
				tweet.Text = longText
			}

			err = goBack(ctx)
			if err != nil {
				log.Error(ctx, err.Error())
				return tweet, FailedToGoBackAfterRetrievingTweetLongText
			}
		}

		return tweet, nil
	}
}
