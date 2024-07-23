package scrapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	"goxcrap/internal/webdriver"
)

const waitTimeAfterLogin time.Duration = 10

// ExecuteHandlerV1 HTTP Handler of the endpoint /scrapper/execute/v1
func ExecuteHandlerV1(newWebDriverManager webdriver.NewManager, newScrapper New, messageBroker broker.MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto criteria.DTO

		bodyBuffer := new(bytes.Buffer)
		teeReader := io.TeeReader(r.Body, bodyBuffer)

		err := json.NewDecoder(teeReader).Decode(&dto)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, CantDecodeBodyIntoCriteria, http.StatusBadRequest)
			return
		}

		webDriverManager := newWebDriverManager()
		defer stop(webDriverManager)

		execute := newScrapper(webDriverManager.WebDriver())
		err = execute(dto.ToType(), waitTimeAfterLogin)
		if err != nil {
			if errors.Is(err, FailedToLogin) || errors.Is(err, search.FailedToLoadAdvanceSearchPage) {
				err = messageBroker.EnqueueMessage(string(bodyBuffer.Bytes()))
				if err != nil {
					http.Error(w, CantReEnqueueFailedMessage, http.StatusInternalServerError)
					return
				}
			}

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
