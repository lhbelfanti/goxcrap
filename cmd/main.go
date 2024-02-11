package main

import (
	"log"
	"time"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/element"
	"goxcrap/cmd/env"
	"goxcrap/cmd/page"
	"goxcrap/cmd/scrapper"
	"goxcrap/internal/chromedriver"
	"goxcrap/internal/setup"
)

func main() {
	/* --- Dependencies --- */
	service := setup.Init(chromedriver.InitWebDriverService())
	defer chromedriver.StopWebDriverService(service)
	driver := setup.Init(chromedriver.InitWebDriver())

	//takeScreenshot := debug.MakeTakeScreenshot(driver) // Debug tool

	variables := setup.Init(env.LoadVariables())

	loadPage := page.MakeLoad(driver)

	waitAndRetrieveElement := element.MakeWaitAndRetrieve(driver)
	retrieveAndFillInput := element.MakeRetrieveAndFillInput(waitAndRetrieveElement)
	retrieveAndClickButton := element.MakeRetrieveAndClickButton(waitAndRetrieveElement)

	// Functions
	login := auth.MakeLogin(variables, loadPage, retrieveAndFillInput, retrieveAndClickButton)

	/* --- Scrapper --- */
	err := scrapper.Init(login)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Minute)
}
