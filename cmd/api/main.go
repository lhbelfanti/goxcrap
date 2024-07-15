package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/fatih/color"

	"goxcrap/cmd/api/ping"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/internal/driver"
)

var localMode bool

func init() {
	flag.BoolVar(&localMode, "local", false, "Run locally instead of in a container")
}

func main() {
	flag.Parse()
	slog.Info(fmt.Sprintf(color.BlueString("Starting GoXCrap with args:\n%s"), color.GreenString("local=%t", localMode)))

	/* --- Program run --- */
	if localMode {
		/* --- Dependencies --- */
		newWebDriver := driver.MakeNew(localMode)
		goXCrapWebDriver, service, webDriver := newWebDriver()
		defer goXCrapWebDriver.StopWebDriverService(service)
		defer goXCrapWebDriver.QuitWebDriver(webDriver)

		newScrapper := scrapper.MakeNew(localMode)
		executeScrapper := newScrapper(webDriver)

		slog.Info(color.BlueString("Executing scrapper..."))
		err := executeScrapper(10)
		if err != nil {
			log.Fatal(color.RedString(err.Error()))
		}
		slog.Info(color.GreenString("Scrapper executed!"))
		time.Sleep(10 * time.Minute)
	} else {
		/* --- Dependencies --- */
		newWebDriver := driver.MakeNew(localMode)
		newScrapper := scrapper.MakeNew(localMode)

		/* --- Router --- */
		slog.Info(color.BlueString("Initializing router..."))
		router := http.NewServeMux()
		router.HandleFunc("GET /ping/v1", ping.HandlerV1())
		router.HandleFunc("POST /execute-scrapper/v1", scrapper.ExecuteHandlerV1(newWebDriver, newScrapper))
		slog.Info(color.GreenString("Router initialized!"))

		/* --- Server --- */
		slog.Info(color.GreenString("GoXCrap server is ready to receive request on port :8091"))
		err := http.ListenAndServe(":8091", router)
		if err != nil {
			log.Fatalf(color.RedString("Could not start server: %s\n", err.Error()))
		}
	}
}
