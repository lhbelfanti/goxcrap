package tweets

import (
	"log/slog"
	"slices"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/page"
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
func MakeRetrieveAll(waitAndRetrieveElements elements.WaitAndRetrieveAll, getTweetHash GetTweetHash, getTweetInformation GetTweetInformation, scrollPage page.Scroll) RetrieveAll {
	return func() ([]Tweet, error) {
		var tweets []Tweet
		for {
			previousTweetsQuantity := len(tweets)

			articles, err := waitAndRetrieveElements(selenium.ByXPATH, globalToLocalXPath(articlesXPath), articlesTimeout)
			if err != nil {
				slog.Error(err.Error())
				return nil, FailedToRetrieveArticles
			}

			for _, article := range articles {
				tweetHash, err := getTweetHash(article)
				if err != nil {
					slog.Error(err.Error())
					continue
				}

				if !slices.ContainsFunc(tweets, compareTweetsByID(tweetHash.ID)) {
					tweet, err := getTweetInformation(article, tweetHash.ID, tweetHash.Timestamp)
					if err != nil {
						slog.Error(err.Error())
						continue
					}
					tweets = append(tweets, tweet)
				}
			}

			if len(tweets) > previousTweetsQuantity {
				err = scrollPage()
				if err != nil {
					slog.Error(err.Error())
					break
				}
				continue
			}

			break
		}

		return tweets, nil
	}
}

// compareTweetsByID returns a function to compare two tweets by ID
func compareTweetsByID(ID string) compareTweets {
	return func(t Tweet) bool {
		return t.ID == ID
	}
}
