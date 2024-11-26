package scrapper

import (
	"os"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/tweets"
	"goxcrap/internal/corpuscreator"
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
		domain := os.Getenv("CORPUS_CREATOR_API_URL")
		saveTweets := corpuscreator.MakeSaveTweets(httpClient, domain)
		updateSearchCriteriaExecution := corpuscreator.MakeUpdateSearchCriteriaExecution(httpClient, domain)
		insertSearchCriteriaExecutionDay := corpuscreator.MakeInsertSearchCriteriaExecutionDay(httpClient, domain)

		// Services
		login := auth.MakeLogin(loadPage, waitAndRetrieveElement, retrieveAndFillInput, retrieveAndClickButton)
		executeAdvanceSearch := search.MakeExecuteAdvanceSearch(loadPage)
		getTweetAuthor := tweets.MakeGetAuthor()
		getTweetTimestamp := tweets.MakeGetTimestamp()
		isAReply := tweets.MakeIsAReply()
		getTweetAvatar := tweets.MakeGetAvatar()
		getTweetText := tweets.MakeGetText()
		getTweetImages := tweets.MakeGetImages()
		hasQuote := tweets.MakeHasQuote()
		isQuoteAReply := tweets.MakeIsQuoteAReply()
		getQuoteAuthor := tweets.MakeGetQuoteAuthor()
		getQuoteAvatar := tweets.MakeGetQuoteAvatar()
		getQuoteTimestamp := tweets.MakeGetQuoteTimestamp()
		getQuoteText := tweets.MakeGetQuoteText()
		getQuoteImages := tweets.MakeGetQuoteImages()
		getTweetHash := tweets.MakeGetTweetHash(getTweetAuthor, getTweetTimestamp)
		getTweetInformation := tweets.MakeGetTweetInformation(isAReply, getTweetAvatar, getTweetText, getTweetImages, hasQuote, isQuoteAReply, getQuoteAuthor, getQuoteAvatar, getQuoteTimestamp, getQuoteText, getQuoteImages)
		retrieveAllTweets := tweets.MakeRetrieveAll(waitAndRetrieveElement, waitAndRetrieveElements, getTweetHash, getTweetInformation, scrollPage)

		executeScrapper := MakeExecute(login, updateSearchCriteriaExecution, insertSearchCriteriaExecutionDay, executeAdvanceSearch, retrieveAllTweets, saveTweets)

		return executeScrapper
	}
}
