package tweets

import (
	"slices"
	"time"

	"github.com/rs/zerolog/log"
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
				log.Info().Msg(err.Error())
				return nil, FailedToRetrieveArticles
			}

			for _, article := range articles {
				tweetHash, err := getTweetHash(article)
				if err != nil {
					log.Info().Msg(err.Error())
					continue
				}

				if !slices.ContainsFunc(tweets, compareTweetsByID(tweetHash.ID)) {
					tweet, err := getTweetInformation(article, tweetHash.ID, tweetHash.Timestamp)
					if err != nil {
						log.Info().Msg(err.Error())
						continue
					}
					tweets = append(tweets, tweet)
				}
			}

			if len(tweets) > previousTweetsQuantity {
				err = scrollPage()
				if err != nil {
					log.Error().Msg(err.Error())
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
