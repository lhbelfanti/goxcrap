package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	"goxcrap/cmd/api/ping"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	_http "goxcrap/internal/http"
	"goxcrap/internal/log"
	"goxcrap/internal/setup"
	"goxcrap/internal/webdriver"
)

// Program arguments
var (
	localMode bool
	prodEnv   bool
)

func init() {
	flag.BoolVar(&localMode, "local", false, "Run locally instead of in a container")
	flag.BoolVar(&prodEnv, "prod", false, "Run in production environment")
	flag.Parse()
}

func main() {
	if localMode {
		runLocal()
	} else {
		runDockerized()
	}
}

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

	newScrapper := scrapper.MakeNew(httpClient)
	executeScrapper := newScrapper(webDriverManager.WebDriver())

	/* --- Run --- */
	log.Info(ctx, "Executing scrapper...")
	setup.Must(executeScrapper(ctx, criteria.MockExampleCriteria(), 10))
	log.Info(ctx, "Scrapper executed!")
	time.Sleep(10 * time.Minute) // Wait time to visually understand what happened
}

func runDockerized() {
	/* --- Dependencies --- */
	ctx := context.Background()

	logLevel := zerolog.DebugLevel
	if prodEnv {
		logLevel = zerolog.InfoLevel
	}

	log.NewCustomLogger(os.Stdout, logLevel)

	httpClient := _http.NewClient()
	messageBroker := setup.Init(broker.NewMessageBroker(ctx, httpClient))
	go messageBroker.InitMessageConsumer(ctx, 2, "/scrapper/execute/v1")

	newWebDriverManager := webdriver.MakeNewManager(localMode)
	newScrapper := scrapper.MakeNew(httpClient)

	/* --- Router --- */
	log.Info(ctx, "Initializing router...")
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /criteria/enqueue/v1", criteria.EnqueueHandlerV1(messageBroker))
	router.HandleFunc("POST /scrapper/execute/v1", scrapper.ExecuteHandlerV1(newWebDriverManager, newScrapper, messageBroker))
	log.Info(ctx, "Router initialized!")

	/* --- Server --- */
	log.Info(ctx, "GoXCrap server is ready to receive request on port :8091")
	setup.Must(http.ListenAndServe(":8091", router))
}
