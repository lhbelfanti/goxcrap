package scrapper

import (
	"fmt"
	"log/slog"
	"time"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/cmd/api/tweets"
)

// Execute starts the X (formerly Twitter) scrapper
type Execute func(searchCriteria criteria.Type, waitTimeAfterLogin time.Duration) error

// MakeExecute creates a new Execute
func MakeExecute(login auth.Login, executeAdvanceSearch search.ExecuteAdvanceSearch, retrieveTweets tweets.RetrieveAll) Execute {
	return func(searchCriteria criteria.Type, waitTimeAfterLogin time.Duration) error {
		err := login()
		if err != nil {
			return FailedToLogin
		}

		slog.Info(fmt.Sprintf("Waiting %d seconds after login", waitTimeAfterLogin))
		time.Sleep(waitTimeAfterLogin * time.Second)

		slog.Info(fmt.Sprintf("Criteria ID: %s", searchCriteria.ID))
		since, until, err := searchCriteria.ParseDates()
		if err != nil {
			slog.Error(err.Error())
			return FailedToParseDatesFromTheGivenCriteria
		}

		currentCriteria := searchCriteria
		for current := since; !current.After(until); current = current.AddDays(1) {
			currentCriteria.Since = current.String()
			currentCriteria.Until = current.AddDays(1).String()
			err := executeAdvanceSearch(currentCriteria)
			if err != nil {
				slog.Error(err.Error())
				continue
			}

			obtainedTweets, err := retrieveTweets()
			if err != nil {
				slog.Error(err.Error())
				continue
			}

			// TODO: save tweets
			slog.Info(fmt.Sprintf("%v", obtainedTweets))
		}

		slog.Info(fmt.Sprintf("All the tweets of the criteria '%s' were retrieved", searchCriteria.ID))

		return nil
	}
}
