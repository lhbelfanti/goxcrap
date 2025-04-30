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

const (
	articlesXPath            string = "//article/div/div/div[2]"
	emptyStateXPath          string = "//*[@data-testid='emptyState']"
	openedTweetArticlesXPath string = "//article/div/div"
)

type (
	// RetrieveAll retrieves all the tweets from the current page
	RetrieveAll func(ctx context.Context) ([]Tweet, error)

	// OpenAndRetrieveArticleByID navigates into a specific tweet and retrieves the <article> HTML element associated with it.
	// This allows us to extract the full text of tweets exceeding 280 characters (the limit for free accounts).
	// The function ensures the retrieved <article> corresponds to the provided tweet ID, even if the tweet is part of a thread
	// or accompanied by replies
	OpenAndRetrieveArticleByID func(ctx context.Context, author, id string) (selenium.WebElement, error)

	// compareTweets compares two tweets by Tweet.ID
	compareTweets func(tweet Tweet) bool
)

// MakeRetrieveAll creates a new RetrieveAll
func MakeRetrieveAll(waitAndRetrieveElement elements.WaitAndRetrieve, waitAndRetrieveElements elements.WaitAndRetrieveAll, getTweetID GetID, getTweetInformation GetTweetInformation, scrollPage page.Scroll) RetrieveAll {
	articlesTimeoutValue, _ := strconv.Atoi(os.Getenv("ARTICLES_TIMEOUT"))
	articlesTimeout := time.Duration(articlesTimeoutValue) * time.Second

	return func(ctx context.Context) ([]Tweet, error) {
		var tweets []Tweet
		for {
			previousTweetsQuantity := len(tweets)

			// If the empty state is present, is not necessary to keep waiting for the articles to appear on screen
			_, err := waitAndRetrieveElement(ctx, selenium.ByXPATH, emptyStateXPath, articlesTimeout)
			if err == nil {
				log.Info(ctx, EmptyStateNoArticlesToRetrieve.Error())
				return nil, EmptyStateNoArticlesToRetrieve
			}

			articles, err := waitAndRetrieveElements(ctx, selenium.ByXPATH, globalToLocalXPath(articlesXPath), articlesTimeout)
			if err != nil {
				log.Error(ctx, err.Error())
				return nil, FailedToRetrieveArticles
			}

			pageURL := log.RetrieveParam[string](ctx, "page_url")
			for _, article := range articles {
				tweetID, err := getTweetID(ctx, article)
				if err != nil {
					continue
				}

				if !slices.ContainsFunc(tweets, compareTweetsByID(tweetID)) {
					tweet, err := getTweetInformation(ctx, article, tweetID)
					ctx = log.With(ctx, log.Param("page_url", pageURL))
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

				// As the scrollPage function was called, there could be more tweets to obtain
				continue
			}

			// All the tweets of the current page were obtained
			break
		}

		return tweets, nil
	}
}

// MakeOpenAndRetrieveArticleByID creates a new OpenAndRetrieveArticleByID
func MakeOpenAndRetrieveArticleByID(openNewTab page.OpenNewTab, waitAndRetrieveElements elements.WaitAndRetrieveAll, getTweetIDFromTweetPage GetIDFromTweetPage) OpenAndRetrieveArticleByID {
	pageLoaderTimeoutValue, _ := strconv.Atoi(os.Getenv("TWEET_PAGE_TIMEOUT"))
	pageLoaderTimeout := time.Duration(pageLoaderTimeoutValue) * time.Second

	articlesTimeoutValue, _ := strconv.Atoi(os.Getenv("ARTICLES_TIMEOUT"))
	articlesTimeout := time.Duration(articlesTimeoutValue) * time.Second

	return func(ctx context.Context, author, id string) (selenium.WebElement, error) {
		url := "/" + author + "/status/" + id

		err := openNewTab(ctx, url, pageLoaderTimeout)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToLoadTweetPage
		}

		articles, err := waitAndRetrieveElements(ctx, selenium.ByXPATH, globalToLocalXPath(openedTweetArticlesXPath), articlesTimeout)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToRetrieveArticles
		}

		var tweetArticle selenium.WebElement
		for _, article := range articles {
			tweetID, err := getTweetIDFromTweetPage(ctx, article)
			if err != nil {
				continue
			}

			if tweetID == id {
				tweetArticle = article
				break
			}
		}

		if tweetArticle == nil {
			log.Debug(ctx, FailedToRetrieveArticle.Error())
			return nil, FailedToRetrieveArticle
		}

		return tweetArticle, nil
	}
}

// compareTweetsByID returns a function to compare two tweets by ID
func compareTweetsByID(ID string) compareTweets {
	return func(t Tweet) bool {
		return t.ID == ID
	}
}
