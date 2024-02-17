package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/element"
	"goxcrap/cmd/env"
	"goxcrap/cmd/page"
	"goxcrap/cmd/scrapper"
	"goxcrap/cmd/search"
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

	waitAndRetrieveCondition := element.MakeWaitAndRetrieveCondition()
	waitAndRetrieveElement := element.MakeWaitAndRetrieve(driver, waitAndRetrieveCondition)
	retrieveAndFillInput := element.MakeRetrieveAndFillInput(waitAndRetrieveElement)
	retrieveAndClickButton := element.MakeRetrieveAndClickButton(waitAndRetrieveElement)

	// Functions
	login := auth.MakeLogin(variables, loadPage, waitAndRetrieveElement, retrieveAndFillInput, retrieveAndClickButton)
	getSearchCriteria := search.MakeGetAdvanceSearchCriteria()
	executeAdvanceSearch := search.MakeExecuteAdvanceSearch(loadPage)

	/* --- Scrapper --- */
	err := scrapper.Execute(login, getSearchCriteria, executeAdvanceSearch)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Minute)
}
