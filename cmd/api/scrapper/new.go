package scrapper

import (
	"log/slog"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/env"
	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/tweets"
	"goxcrap/internal/setup"
)

// New initializes all the functions of a scrapper (only Execute for now)
// It also initializes all the dependencies for its functions
type New func(webDriver selenium.WebDriver) Execute

// MakeNew creates a new New
func MakeNew(localMode bool) New {
	return func(webDriver selenium.WebDriver) Execute {
		slog.Info(color.BlueString("Loading env variables..."))
		if localMode {
			setup.Init(0, godotenv.Load())
		}
		variables := env.LoadVariables()
		slog.Info(color.GreenString("Env variables initialized!"))

		// Functions
		slog.Info(color.BlueString("Initializing functions..."))
		loadPage := page.MakeLoad(webDriver)
		scrollPage := page.MakeScroll(webDriver)
		waitAndRetrieveCondition := elements.MakeWaitAndRetrieveCondition()
		waitAndRetrieveAllCondition := elements.MakeWaitAndRetrieveAllCondition()
		waitAndRetrieveElement := elements.MakeWaitAndRetrieve(webDriver, waitAndRetrieveCondition)
		waitAndRetrieveElements := elements.MakeWaitAndRetrieveAll(webDriver, waitAndRetrieveAllCondition)
		retrieveAndFillInput := elements.MakeRetrieveAndFillInput(waitAndRetrieveElement)
		retrieveAndClickButton := elements.MakeRetrieveAndClickButton(waitAndRetrieveElement)
		slog.Info(color.GreenString("Functions initialized!"))

		// Services
		slog.Info(color.BlueString("Initializing services dependencies..."))
		login := auth.MakeLogin(variables, loadPage, waitAndRetrieveElement, retrieveAndFillInput, retrieveAndClickButton)
		getSearchCriteria := search.MakeGetAdvanceSearchCriteria()
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
		slog.Info(color.GreenString("Services dependencies initialized!"))

		slog.Info(color.BlueString("Initializing services..."))
		executeScrapper := MakeExecute(login, getSearchCriteria, executeAdvanceSearch, retrieveAllTweets)
		slog.Info(color.GreenString("Services initialized!"))

		return executeScrapper
	}
}