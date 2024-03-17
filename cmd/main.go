package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/elements"
	"goxcrap/cmd/env"
	"goxcrap/cmd/page"
	"goxcrap/cmd/scrapper"
	"goxcrap/cmd/search"
	"goxcrap/cmd/tweets"
	"goxcrap/internal/chromedriver"
	"goxcrap/internal/setup"
)

func main() {
	/* --- Dependencies --- */
	service := setup.Init(chromedriver.InitWebDriverService())
	defer chromedriver.StopWebDriverService(service)
	driver := setup.Init(chromedriver.InitWebDriver())

	//takeScreenshot := debug.MakeTakeScreenshot(driver) // Debug tool

	setup.Init(0, godotenv.Load())
	variables := env.LoadVariables()

	loadPage := page.MakeLoad(driver)

	waitAndRetrieveCondition := elements.MakeWaitAndRetrieveCondition()
	waitAndRetrieveAllCondition := elements.MakeWaitAndRetrieveAllCondition()
	waitAndRetrieveElement := elements.MakeWaitAndRetrieve(driver, waitAndRetrieveCondition)
	waitAndRetrieveElements := elements.MakeWaitAndRetrieveAll(driver, waitAndRetrieveAllCondition)
	retrieveAndFillInput := elements.MakeRetrieveAndFillInput(waitAndRetrieveElement)
	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(waitAndRetrieveElement)

	// Functions
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
	gatherTweetInformation := tweets.MakeGetTweetInformation(getTweetAuthor, getTweetTimestamp, isAReply, getTweetText, getTweetImages, hasQuote, isQuoteAReply, getQuoteText, getQuoteImages)
	retrieveAllTweets := tweets.MakeRetrieveAll(waitAndRetrieveElements, gatherTweetInformation)

	/* --- Scrapper --- */
	executeScrapper := scrapper.MakeExecute(login, getSearchCriteria, executeAdvanceSearch, retrieveAllTweets)
	err := executeScrapper(10)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Minute)
}
