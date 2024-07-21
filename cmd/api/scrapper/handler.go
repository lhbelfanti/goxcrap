package scrapper

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/webdriver"
)

const waitTimeAfterLogin time.Duration = 10

// ExecuteHandlerV1 HTTP Handler of the endpoint /execute-scrapper/v1
func ExecuteHandlerV1(newWebDriverManager webdriver.NewManager, newScrapper New) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto criteria.DTO
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, InvalidBody, http.StatusBadRequest)
			return
		}

		webDriverManager := newWebDriverManager()
		defer stop(webDriverManager)

		execute := newScrapper(webDriverManager.WebDriver())
		err = execute(dto.ToType(), waitTimeAfterLogin)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, FailedToRunScrapper, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Scrapper successfully executed"))
	}
}

// stop deferred function that handles the webDriver.Quit method
func stop(webDriverManager webdriver.Manager) {
	err := webDriverManager.Quit()
	if err != nil {
		slog.Error(err.Error())
	}
}
