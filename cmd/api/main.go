package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/env"
	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/ping"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/tweets"
	"goxcrap/internal/driver"
	"goxcrap/internal/setup"
)

func main() {
	/* --- Dependencies --- */
	service := setup.Init(driver.InitWebDriverService())
	defer driver.StopWebDriverService(service)
	webDriver := setup.Init(driver.InitWebDriver())

	//takeScreenshot := debug.MakeTakeScreenshot(webDriver) // Debug tool

	setup.Init(0, godotenv.Load())
	variables := env.LoadVariables()

	// Functions
	loadPage := page.MakeLoad(webDriver)
	waitAndRetrieveCondition := elements.MakeWaitAndRetrieveCondition()
	waitAndRetrieveAllCondition := elements.MakeWaitAndRetrieveAllCondition()
	waitAndRetrieveElement := elements.MakeWaitAndRetrieve(webDriver, waitAndRetrieveCondition)
	waitAndRetrieveElements := elements.MakeWaitAndRetrieveAll(webDriver, waitAndRetrieveAllCondition)
	retrieveAndFillInput := elements.MakeRetrieveAndFillInput(waitAndRetrieveElement)
	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(waitAndRetrieveElement)

	// Services
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
	scrollPage := page.MakeScroll(webDriver)
	retrieveAllTweets := tweets.MakeRetrieveAll(waitAndRetrieveElements, gatherTweetInformation, scrollPage)
	executeScrapper := scrapper.MakeExecute(login, getSearchCriteria, executeAdvanceSearch, retrieveAllTweets)

	/* --- Router --- */
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /execute-scrapper/v1", scrapper.ExecuteHandlerV1(executeScrapper))

	/* --- Server --- */
	log.Println("Starting GoXCrap server on :8091")
	err := http.ListenAndServe(":8091", router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
