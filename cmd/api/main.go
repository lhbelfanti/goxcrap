package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
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

var localMode bool

func init() {
	flag.BoolVar(&localMode, "local", false, "Run locally instead of in a container")
}

func main() {
	flag.Parse()
	slog.Info(fmt.Sprintf(color.BlueString("Starting GoXCrap with args:\n%s"), color.GreenString("local=%t", localMode)))

	/* --- Dependencies --- */
	slog.Info(color.BlueString("Initializing WebDriver..."))
	var goXCrapWebDriver driver.GoXCrapWebDriver
	if localMode {
		goXCrapWebDriver = driver.NewLocalWebDriver()
	} else {
		goXCrapWebDriver = driver.NewDockerizedWebDriver()
	}
	service := setup.Init(goXCrapWebDriver.InitWebDriverService())
	defer goXCrapWebDriver.StopWebDriverService(service)
	webDriver := setup.Init(goXCrapWebDriver.InitWebDriver())
	defer goXCrapWebDriver.QuitWebDriver(webDriver)
	slog.Info(color.GreenString("WebDriver initialized!"))

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
	executeScrapper := scrapper.MakeExecute(login, getSearchCriteria, executeAdvanceSearch, retrieveAllTweets)
	slog.Info(color.GreenString("Services initialized!"))

	/* --- Program run --- */
	if localMode {
		slog.Info(color.BlueString("Executing scrapper..."))
		err := executeScrapper(10)
		if err != nil {
			log.Fatal(color.RedString(err.Error()))
		}
		slog.Info(color.GreenString("Scrapper executed!"))
		time.Sleep(10 * time.Minute)
	} else {
		/* --- Router --- */
		slog.Info(color.BlueString("Initializing router..."))
		router := http.NewServeMux()
		router.HandleFunc("GET /ping/v1", ping.HandlerV1())
		router.HandleFunc("POST /execute-scrapper/v1", scrapper.ExecuteHandlerV1(executeScrapper))
		slog.Info(color.GreenString("Router initialized!"))

		/* --- Server --- */
		slog.Info(color.GreenString("GoXCrap server is ready to receive request on port :8091"))
		err := http.ListenAndServe(":8091", router)
		if err != nil {
			log.Fatalf(color.RedString("Could not start server: %s\n", err.Error()))
		}
	}
}
