package scrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	"goxcrap/internal/log"
	"goxcrap/internal/webdriver"
)

// ExecuteHandlerV1 HTTP Handler of the endpoint /scrapper/execute/v1
func ExecuteHandlerV1(newWebDriverManager webdriver.NewManager, newScrapper New, messageBroker broker.MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bodyBuffer := new(bytes.Buffer)
		teeReader := io.TeeReader(r.Body, bodyBuffer)

		var dto criteria.MessageDTO
		err := json.NewDecoder(teeReader).Decode(&dto)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, CantDecodeBodyIntoCriteria, http.StatusBadRequest)
			return
		}

		webDriverManager := newWebDriverManager(ctx)
		defer stop(ctx, webDriverManager)

		execute := newScrapper(webDriverManager.WebDriver())
		err = execute(ctx, dto.Criteria.ToType(), dto.ExecutionID)
		if err != nil {
			if errors.Is(err, FailedToLogin) || errors.Is(err, search.FailedToLoadAdvanceSearchPage) {
				enqueueErr := messageBroker.EnqueueMessage(ctx, string(bodyBuffer.Bytes()))
				if enqueueErr != nil {
					log.Error(ctx, enqueueErr.Error())
					http.Error(w, CantReEnqueueFailedMessage, http.StatusInternalServerError)
					return
				}
			}

			log.Error(ctx, err.Error())
			http.Error(w, FailedToRunScrapper, http.StatusInternalServerError)
			return
		}

		log.Info(ctx, "Scrapper successfully executed")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Scrapper successfully executed"))
	}
}

// stop deferred function that handles the webDriver.Quit method
func stop(ctx context.Context, webDriverManager webdriver.Manager) {
	err := webDriverManager.Quit(ctx)
	if err != nil {
		log.Warn(ctx, err.Error())
	}
}
