package scrapper

import (
	"context"
	"fmt"
	"time"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/cmd/api/tweets"
	"goxcrap/internal/ahbcc"
	"goxcrap/internal/log"
)

// Execute starts the X (formerly Twitter) scrapper
type Execute func(ctx context.Context, searchCriteria criteria.Type, waitTimeAfterLogin time.Duration) error

// MakeExecute creates a new Execute
func MakeExecute(login auth.Login, executeAdvanceSearch search.ExecuteAdvanceSearch, retrieveTweets tweets.RetrieveAll, saveTweets ahbcc.SaveTweets) Execute {
	return func(ctx context.Context, searchCriteria criteria.Type, waitTimeAfterLogin time.Duration) error {
		err := login(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToLogin
		}

		log.Debug(ctx, fmt.Sprintf("Waiting %d seconds after login", waitTimeAfterLogin))
		time.Sleep(waitTimeAfterLogin * time.Second)

		log.Debug(ctx, fmt.Sprintf("Criteria ID: %d", searchCriteria.ID))
		ctx = log.With(ctx, log.Param("criteria_id", searchCriteria.ID))

		since, until, err := searchCriteria.ParseDates()
		ctx = log.With(ctx, log.Param("criteria_since", searchCriteria.Since), log.Param("criteria_until", searchCriteria.Until))
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToParseDatesFromTheGivenCriteria
		}

		currentCriteria := searchCriteria
		for current := since; !current.After(until); current = current.AddDays(1) {
			currentCriteria.Since = current.String()
			currentCriteria.Until = current.AddDays(1).String()
			err = executeAdvanceSearch(ctx, currentCriteria)
			if err != nil {
				continue
			}

			obtainedTweets, err := retrieveTweets(ctx)
			if err != nil {
				continue
			}

			if len(obtainedTweets) > 0 {
				requestBody := createSaveTweetsBody(obtainedTweets, currentCriteria.ID)
				err = saveTweets(ctx, requestBody)
				if err != nil {
					continue
				}
			}

		}

		log.Info(ctx, "All the tweets of the criteria were retrieved")

		return nil
	}
}

// createSaveTweetsBody creates the SaveTweets Body with the obtained []tweets.Tweet
func createSaveTweetsBody(obtainedTweets []tweets.Tweet, searchCriteria int) ahbcc.SaveTweetsBody {
	saveTweetsBody := make(ahbcc.SaveTweetsBody, 0, len(obtainedTweets))
	for _, tweet := range obtainedTweets {
		requestTweet := ahbcc.TweetDTO{
			Hash:             &tweet.ID,
			IsAReply:         tweet.IsAReply,
			SearchCriteriaID: &searchCriteria,
		}

		if tweet.HasText {
			requestTweet.TextContent = &tweet.Text
		}

		if tweet.HasImages {
			requestTweet.Images = tweet.Images
		}

		if tweet.HasQuote {
			requestTweet.Quote = &ahbcc.QuoteDTO{IsAReply: tweet.Quote.IsAReply}

			if tweet.Quote.HasText {
				requestTweet.Quote.TextContent = &tweet.Quote.Text
			}

			if tweet.Quote.HasImages {
				requestTweet.Quote.Images = tweet.Quote.Images
			}
		}

		saveTweetsBody = append(saveTweetsBody, requestTweet)
	}

	return saveTweetsBody
}
