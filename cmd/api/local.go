package main

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"

	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search/criteria"
	_http "goxcrap/internal/http"
	"goxcrap/internal/log"
	"goxcrap/internal/setup"
	"goxcrap/internal/webdriver"
)

func runLocal() {
	/* --- Dependencies --- */
	ctx := context.Background()

	log.Info(ctx, fmt.Sprintf("Starting GoXCrap with args: local=%t prod=%t", localMode, prodEnv))

	httpClient := _http.NewClient()

	setup.Must(godotenv.Load())

	newWebDriverManager := webdriver.MakeNewManager(localMode)
	webDriverManager := newWebDriverManager(ctx)
	defer func(webDriverManager webdriver.Manager) {
		err := webDriverManager.Quit(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
		}
	}(webDriverManager)

	newScrapper := scrapper.MakeNew(httpClient, localMode)
	executeScrapper := newScrapper(webDriverManager.WebDriver())

	/* --- Run --- */
	log.Info(ctx, "Executing scrapper...")
	setup.Must(executeScrapper(ctx, criteria.MockExampleCriteria(), 1))
	log.Info(ctx, "Scrapper executed!")
	time.Sleep(10 * time.Minute) // Wait time to visually understand what happened
}
