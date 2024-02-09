package main

import (
	"goxcrap/cmd/app"
	"log"
	"os"

	"github.com/joho/godotenv"

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
	pageLoader := scrap.MakePageLoader(driver)

	scrapper := scrap.New(driver, email, password, pageLoader)

	/* Program */
	err := app.Init(scrapper)
	if err != nil {
		log.Fatal(err)
	}
}
