package scrapper

import (
	"fmt"
	"log/slog"
	"time"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/search"
	"goxcrap/cmd/tweets"
)

// Execute starts the twitter scrapper
type Execute func(waitAfterLogin time.Duration) error

// MakeExecute creates a new Execute
func MakeExecute(login auth.Login, getAdvanceSearchCriteria search.GetAdvanceSearchCriteria, executeAdvanceSearch search.ExecuteAdvanceSearch, retrieveTweets tweets.RetrieveAll) Execute {
	return func(secondsAfterLogin time.Duration) error {
		err := login()
		if err != nil {
			return err
		}
		slog.Info("Log In completed")

		slog.Info(fmt.Sprintf("Waiting %d after login", secondsAfterLogin))
		time.Sleep(secondsAfterLogin * time.Second)

		searchCriteria := getAdvanceSearchCriteria()
		for _, criteria := range searchCriteria {
			slog.Info(fmt.Sprintf("Criteria: %s", criteria.ID))
			since, until, err := criteria.ParseDates()
			if err != nil {
				slog.Error(err.Error())
				continue
			}

			currentCriteria := criteria
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

				slog.Info("Obtained tweets", obtainedTweets)
			}

			slog.Info(fmt.Sprintf("All the tweets of the criteria '%s' were retrieved", criteria.ID))
		}

		return nil
	}
}
