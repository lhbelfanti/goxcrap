package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"goxcrap/cmd/app"
	scrap "goxcrap/cmd/scrapper"
	"goxcrap/internal/chromedriver"
	"goxcrap/internal/setup"
)

func main() {
	/* Dependencies */
	service := setup.Must(chromedriver.InitWebDriverService())
	defer chromedriver.StopWebDriverService(service)
	driver := setup.Must(chromedriver.InitWebDriver())

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	username := os.Getenv("USERNAME")
	loadPage := scrap.MakeLoadPage(driver)
	waitAndRetrieveElement := scrap.MakeWaitAndRetrieveElement(driver)
	takeScreenshot := scrap.MakeTakeScreenshot(driver)
	scrapper := scrap.New(driver, email, password, username, takeScreenshot, loadPage, waitAndRetrieveElement)

	/* Program */
	err := app.Init(scrapper)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Minute)
}
