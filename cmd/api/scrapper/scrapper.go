package scrapper

import (
	"os"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/tweets"
	"goxcrap/internal/ahbcc"
	"goxcrap/internal/http"
)

// New initializes all the functions of a scrapper (only Execute for now)
// It also initializes all the dependencies for its functions
type New func(webDriver selenium.WebDriver) Execute

// MakeNew creates a new New
func MakeNew(httpClient http.Client) New {
	return func(webDriver selenium.WebDriver) Execute {
		// Functions
		loadPage := page.MakeLoad(webDriver)
		scrollPage := page.MakeScroll(webDriver)
		waitAndRetrieveCondition := elements.MakeWaitAndRetrieveCondition()
		waitAndRetrieveAllCondition := elements.MakeWaitAndRetrieveAllCondition()
		waitAndRetrieveElement := elements.MakeWaitAndRetrieve(webDriver, waitAndRetrieveCondition)
		waitAndRetrieveElements := elements.MakeWaitAndRetrieveAll(webDriver, waitAndRetrieveAllCondition)
		retrieveAndFillInput := elements.MakeRetrieveAndFillInput(waitAndRetrieveElement)
		retrieveAndClickButton := elements.MakeRetrieveAndClickButton(waitAndRetrieveElement)

		// Calls to external services
		saveTweets := ahbcc.MakeSaveTweets(httpClient, os.Getenv("SAVE_TWEETS_API_URL"))

		// Services
		login := auth.MakeLogin(loadPage, waitAndRetrieveElement, retrieveAndFillInput, retrieveAndClickButton)
		executeAdvanceSearch := search.MakeExecuteAdvanceSearch(loadPage)
		getTweetAuthor := tweets.MakeGetAuthor()
		getTweetTimestamp := tweets.MakeGetTimestamp()
		isAReply := tweets.MakeIsAReply()
		getTweetText := tweets.MakeGetText()
		getTweetImages := tweets.MakeGetImages()
		hasQuote := tweets.MakeHasQuote()
		isQuoteAReply := tweets.MakeIsQuoteAReply()
		getQuoteText := tweets.MakeGetQuoteText()
		getQuoteImages := tweets.MakeGetQuoteImages()
		getTweetHash := tweets.MakeGetTweetHash(getTweetAuthor, getTweetTimestamp)
		getTweetInformation := tweets.MakeGetTweetInformation(isAReply, getTweetText, getTweetImages, hasQuote, isQuoteAReply, getQuoteText, getQuoteImages)
		retrieveAllTweets := tweets.MakeRetrieveAll(waitAndRetrieveElements, getTweetHash, getTweetInformation, scrollPage)

		executeScrapper := MakeExecute(login, executeAdvanceSearch, retrieveAllTweets, saveTweets)

		return executeScrapper
	}
}
