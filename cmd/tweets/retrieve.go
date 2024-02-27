package tweets

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
)

const (
	articlesXPath   string        = "//article"
	articlesTimeout time.Duration = 10 * time.Second
)

// RetrieveAll retrieves all the tweets from the current page
type RetrieveAll func() ([]Tweet, error)

// MakeRetrieveAll creates a new RetrieveAll
func MakeRetrieveAll(waitAndRetrieveElements elements.WaitAndRetrieveAll, gatherTweetInformation GatherTweetInformation) RetrieveAll {
	return func() ([]Tweet, error) {
		articles, err := waitAndRetrieveElements(selenium.ByXPATH, articlesXPath, articlesTimeout)
		if err != nil {
			return nil, NewTweetsError(FailedToRetrieveArticles, err)
		}

		var tweets []Tweet
		for _, article := range articles {
			tweet, err := gatherTweetInformation(article)
			if err != nil {
				fmt.Println(err)
				continue
			}

			tweets = append(tweets, tweet)
		}

		return tweets, nil
	}
}
