package tweets

import (
	"log/slog"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
)

const (
	articlesTimeout time.Duration = 10 * time.Second

	articlesXPath string = "//article"
)

// RetrieveAll retrieves all the tweets from the current page
type RetrieveAll func() ([]Tweet, error)

// MakeRetrieveAll creates a new RetrieveAll
func MakeRetrieveAll(waitAndRetrieveElements elements.WaitAndRetrieveAll, gatherTweetInformation GatherTweetInformation) RetrieveAll {
	return func() ([]Tweet, error) {
		articles, err := waitAndRetrieveElements(selenium.ByXPATH, articlesXPath, articlesTimeout)
		if err != nil {
			slog.Error(err.Error())
			return nil, FailedToRetrieveArticles
		}

		var tweets []Tweet
		for _, article := range articles {
			tweet, err := gatherTweetInformation(article)
			if err != nil {
				slog.Error(err.Error())
				continue
			}

			tweets = append(tweets, tweet)
		}

		return tweets, nil
	}
}
