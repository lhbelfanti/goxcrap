package tweets

import (
	"context"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/page"
	"goxcrap/internal/log"
)

const articlesXPath string = "//article/div/div/div[2]/div[2]"

type (
	// RetrieveAll retrieves all the tweets from the current page
	RetrieveAll func(ctx context.Context) ([]Tweet, error)

	// compareTweets compares two tweets by Tweet.ID
	compareTweets func(Tweet) bool
)

// MakeRetrieveAll creates a new RetrieveAll
func MakeRetrieveAll(waitAndRetrieveElements elements.WaitAndRetrieveAll, getTweetHash GetTweetHash, getTweetInformation GetTweetInformation, scrollPage page.Scroll) RetrieveAll {
	articlesTimeoutValue, _ := strconv.Atoi(os.Getenv("ARTICLES_TIMEOUT"))
	articlesTimeout := time.Duration(articlesTimeoutValue) * time.Second

	return func(ctx context.Context) ([]Tweet, error) {
		var tweets []Tweet
		for {
			previousTweetsQuantity := len(tweets)

			articles, err := waitAndRetrieveElements(ctx, selenium.ByXPATH, globalToLocalXPath(articlesXPath), articlesTimeout)
			if err != nil {
				log.Error(ctx, err.Error())
				return nil, FailedToRetrieveArticles
			}

			for _, article := range articles {
				tweetHash, err := getTweetHash(ctx, article)
				if err != nil {
					continue
				}

				if !slices.ContainsFunc(tweets, compareTweetsByID(tweetHash.ID)) {
					tweet, err := getTweetInformation(ctx, article, tweetHash.ID, tweetHash.Timestamp)
					if err != nil {
						continue
					}
					tweets = append(tweets, tweet)
				}
			}

			if len(tweets) > previousTweetsQuantity {
				err = scrollPage(ctx)
				if err != nil {
					break
				}

				// As the scrollPage function was called there could be more tweets to obtain
				continue
			}

			// All the tweets of the current page were obtained
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
