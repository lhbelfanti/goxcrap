package scrapper

import (
	"fmt"
	"log/slog"
	"time"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/tweets"
)

// Execute starts the X (formerly Twitter) scrapper
type Execute func(criteria search.Criteria, waitTimeAfterLogin time.Duration) error

// MakeExecute creates a new Execute
func MakeExecute(login auth.Login, executeAdvanceSearch search.ExecuteAdvanceSearch, retrieveTweets tweets.RetrieveAll) Execute {
	return func(criteria search.Criteria, waitTimeAfterLogin time.Duration) error {
		err := login()
		if err != nil {
			return err
		}
		slog.Info("Log In completed")

		slog.Info(fmt.Sprintf("Waiting %d seconds after login", waitTimeAfterLogin))
		time.Sleep(waitTimeAfterLogin * time.Second)

		for _, criterion := range criteria {
			slog.Info(fmt.Sprintf("Criterion: %s", criterion.ID))
			since, until, err := criterion.ParseDates()
			if err != nil {
				slog.Error(err.Error())
				continue
			}

			currentCriterion := criterion
			for current := since; !current.After(until); current = current.AddDays(1) {
				currentCriterion.Since = current.String()
				currentCriterion.Until = current.AddDays(1).String()
				err := executeAdvanceSearch(currentCriterion)
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

			slog.Info(fmt.Sprintf("All the tweets of the criterion '%s' were retrieved", criterion.ID))
		}

		return nil
	}
}
