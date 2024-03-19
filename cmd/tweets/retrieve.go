package tweets

import (
	"log/slog"
	"slices"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
	"goxcrap/cmd/page"
)

const (
	articlesTimeout time.Duration = 10 * time.Second

	articlesXPath string = "//article/div/div/div[2]/div[2]"
)

type (
	// RetrieveAll retrieves all the tweets from the current page
	RetrieveAll func() ([]Tweet, error)

	// compareTweets compares two tweets by Tweet.ID
	compareTweets func(Tweet) bool
)

// MakeRetrieveAll creates a new RetrieveAll
func MakeRetrieveAll(waitAndRetrieveElements elements.WaitAndRetrieveAll, gatherTweetInformation GatherTweetInformation, scrollPage page.Scroll) RetrieveAll {
	return func() ([]Tweet, error) {
		var tweets []Tweet

		for {
			tweetsIDsLenBefore := len(tweets)

			articles, err := waitAndRetrieveElements(selenium.ByXPATH, globalToLocalXPath(articlesXPath), articlesTimeout)
			if err != nil {
				slog.Error(err.Error())
				return nil, FailedToRetrieveArticles
			}

			for _, article := range articles {
				tweet, err := gatherTweetInformation(article)
				if err != nil {
					slog.Error(err.Error())
					continue
				}

				if !slices.ContainsFunc(tweets, tweetsComparator(tweet)) {
					tweets = append(tweets, tweet)
				}
			}

			tweetsIDsLenAfter := len(tweets)

			if tweetsIDsLenAfter > tweetsIDsLenBefore {
				err = scrollPage()
				if err != nil {
					slog.Error(err.Error())
					break
				}
			}
		}

		return tweets, nil
	}
}

// tweetsComparator returns a function to compare two tweets by ID
func tweetsComparator(oldTweet Tweet) compareTweets {
	return func(newTweet Tweet) bool {
		return oldTweet.ID == newTweet.ID
	}
}
