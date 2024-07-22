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

	"goxcrap/cmd/api/ping"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	"goxcrap/internal/setup"
	"goxcrap/internal/webdriver"
)

var localMode bool

func init() {
	flag.BoolVar(&localMode, "local", false, "Run locally instead of in a container")
}

func main() {
	flag.Parse()
	slog.Info(fmt.Sprintf(color.BlueString("Starting GoXCrap with args:\n%s"), color.GreenString("local=%t", localMode)))

	if localMode {
		runLocal()
	} else {
		runDockerized()
	}
}

func runLocal() {
	/* --- Dependencies --- */
	setup.Must(godotenv.Load())
	newScrapper := scrapper.MakeNew()
	newWebDriverManager := webdriver.MakeNewManager(localMode)
	webDriverManager := newWebDriverManager()
	defer func(webDriverManager webdriver.Manager) {
		err := webDriverManager.Quit()
		if err != nil {
			slog.Error(err.Error())
		}
	}(webDriverManager)

	executeScrapper := newScrapper(webDriverManager.WebDriver())

	/* --- Run --- */
	slog.Info(color.BlueString("Executing scrapper..."))
	err := executeScrapper(criteria.MockExampleCriteria(), 10)
	if err != nil {
		log.Fatal(color.RedString(err.Error()))
	}
	slog.Info(color.GreenString("Scrapper executed!"))
	time.Sleep(10 * time.Minute)
}

func runDockerized() {
	/* --- Dependencies --- */
	messageBroker := setup.Init(broker.NewMessageBroker())
	go messageBroker.InitMessageConsumer(2, "/execute-scrapper/v1")

	newWebDriverManager := webdriver.MakeNewManager(localMode)
	newScrapper := scrapper.MakeNew()

	/* --- Router --- */
	slog.Info(color.BlueString("Initializing router..."))
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /enqueue-criteria/v1", criteria.EnqueueHandlerV1(messageBroker))
	router.HandleFunc("POST /execute-scrapper/v1", scrapper.ExecuteHandlerV1(newWebDriverManager, newScrapper, messageBroker))
	slog.Info(color.GreenString("Router initialized!"))

	/* --- Server --- */
	slog.Info(color.GreenString("GoXCrap server is ready to receive request on port :8091"))
	err := http.ListenAndServe(":8091", router)
	if err != nil {
		log.Fatalf(color.RedString("Could not start server: %s\n", err.Error()))
	}
}
