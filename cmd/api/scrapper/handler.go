package scrapper

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"goxcrap/cmd/api/search"
	"goxcrap/internal/driver"
)

const waitTimeAfterLogin time.Duration = 10

// ExecuteHandlerV1 HTTP Handler of the endpoint /execute-scrapper/v1
func ExecuteHandlerV1(newWebDriver driver.New, newScrapper New) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var criteria search.Criteria
		if json.NewDecoder(r.Body).Decode(&criteria) != nil {
			http.Error(w, InvalidBody, http.StatusBadRequest)
			return
		}

		goXCrapWebDriver, service, webDriver := newWebDriver()
		defer goXCrapWebDriver.StopWebDriverService(service)
		defer goXCrapWebDriver.QuitWebDriver(webDriver)

		execute := newScrapper(webDriver)
		err := execute(criteria, waitTimeAfterLogin)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, FailedToRunScrapper, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Scrapper successfully executed"))
	}
}
