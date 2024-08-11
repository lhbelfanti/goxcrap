package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"goxcrap/cmd/api/ping"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	_http "goxcrap/internal/http"
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
	log.Info().Msgf("Starting GoXCrap with args:\nlocal=%t\nprod=%t", localMode, prodEnv)

	if localMode {
		runLocal()
	} else {
		runDockerized()
	}
}

func runLocal() {
	/* --- Dependencies --- */
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Logger = logger

	httpClient := _http.NewClient()

	setup.Must(godotenv.Load())

	newWebDriverManager := webdriver.MakeNewManager(localMode)
	webDriverManager := newWebDriverManager()
	defer func(webDriverManager webdriver.Manager) {
		err := webDriverManager.Quit()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}(webDriverManager)

	newScrapper := scrapper.MakeNew(httpClient)
	executeScrapper := newScrapper(webDriverManager.WebDriver())

	/* --- Run --- */
	log.Info().Msg("Executing scrapper...")
	setup.Must(executeScrapper(criteria.MockExampleCriteria(), 10))
	log.Info().Msg("Scrapper executed!")
	time.Sleep(10 * time.Minute) // Wait time to visually understand what happened
}

func runDockerized() {
	/* --- Dependencies --- */
	// TODO: connect to ELK to save logs
	//conn := setup.Init(net.Dial("tcp", "logstash:5000"))
	//logger := zerolog.New(conn).With().Timestamp().Logger()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Logger = logger

	if prodEnv {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	httpClient := _http.NewClient()

	messageBroker := setup.Init(broker.NewMessageBroker(httpClient))
	go messageBroker.InitMessageConsumer(2, "/scrapper/execute/v1")

	newWebDriverManager := webdriver.MakeNewManager(localMode)
	newScrapper := scrapper.MakeNew(httpClient)

	/* --- Router --- */
	log.Info().Msg("Initializing router...")
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /criteria/enqueue/v1", criteria.EnqueueHandlerV1(messageBroker))
	router.HandleFunc("POST /scrapper/execute/v1", scrapper.ExecuteHandlerV1(newWebDriverManager, newScrapper, messageBroker))
	log.Info().Msg("Router initialized!")

	/* --- Server --- */
	log.Info().Msg("GoXCrap server is ready to receive request on port :8091")
	setup.Must(http.ListenAndServe(":8091", router))
}
