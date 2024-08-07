package scrapper

import (
	"time"

	"github.com/rs/zerolog/log"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/cmd/api/tweets"
	"goxcrap/internal/ahbcc"
)

// Execute starts the X (formerly Twitter) scrapper
type Execute func(searchCriteria criteria.Type, waitTimeAfterLogin time.Duration) error

// MakeExecute creates a new Execute
func MakeExecute(login auth.Login, executeAdvanceSearch search.ExecuteAdvanceSearch, retrieveTweets tweets.RetrieveAll, saveTweets ahbcc.SaveTweets) Execute {
	return func(searchCriteria criteria.Type, waitTimeAfterLogin time.Duration) error {
		err := login()
		if err != nil {
			return FailedToLogin
		}

		log.Info().Msgf("Waiting %d seconds after login", waitTimeAfterLogin)
		time.Sleep(waitTimeAfterLogin * time.Second)

		log.Info().Msgf("Criteria ID: %d", searchCriteria.ID)
		since, until, err := searchCriteria.ParseDates()
		if err != nil {
			log.Error().Msg(err.Error())
			return FailedToParseDatesFromTheGivenCriteria
		}

		currentCriteria := searchCriteria
		for current := since; !current.After(until); current = current.AddDays(1) {
			currentCriteria.Since = current.String()
			currentCriteria.Until = current.AddDays(1).String()
			err := executeAdvanceSearch(currentCriteria)
			if err != nil {
				log.Info().Msg(err.Error())
				continue
			}

			obtainedTweets, err := retrieveTweets()
			if err != nil {
				log.Info().Msg(err.Error())
				continue
			}

			if len(obtainedTweets) > 0 {
				requestBody := createSaveTweetsBody(obtainedTweets, currentCriteria.ID)
				err = saveTweets(requestBody)
				if err != nil {
					log.Info().Msg(err.Error())
					continue
				}
			}

			log.Info().Msgf("%v", obtainedTweets)
		}

		log.Info().Msgf("All the tweets of the criteria '%d' were retrieved", searchCriteria.ID)

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
