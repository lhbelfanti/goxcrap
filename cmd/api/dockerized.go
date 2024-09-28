package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

func runDockerized() {
	/* --- Dependencies --- */
	ctx := context.Background()

	logLevel := zerolog.DebugLevel
	if prodEnv {
		logLevel = zerolog.InfoLevel
	}

	log.NewCustomLogger(os.Stdout, logLevel)

	httpClient := _http.NewClient()

	newWebDriverManager := webdriver.MakeNewManager(localMode)
	newScrapper := scrapper.MakeNew(httpClient)

	messageBroker := setup.Init(broker.NewMessageBroker(ctx, httpClient))
	concurrentMessages := setup.Init(strconv.Atoi(os.Getenv("BROKER_CONCURRENT_MESSAGES")))
	searchCriteriaMessageProcessor := scrapper.MakeSearchCriteriaMessageProcessor(newWebDriverManager, newScrapper, messageBroker)
	go messageBroker.InitMessageConsumerWithFunction(concurrentMessages, searchCriteriaMessageProcessor)

	/* --- Router --- */
	log.Info(ctx, "Initializing router...")
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /criteria/enqueue/v1", criteria.EnqueueHandlerV1(messageBroker))
	router.HandleFunc("POST /scrapper/execute/v1", scrapper.ExecuteHandlerV1(newWebDriverManager, newScrapper, messageBroker))
	log.Info(ctx, "Router initialized!")

	/* --- Server --- */
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Info(ctx, fmt.Sprintf("GoXCrap server is ready to receive request on port %s", port))
	setup.Must(http.ListenAndServe(port, router))
}
