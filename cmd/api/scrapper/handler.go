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
	"goxcrap/internal/http/response"
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
			response.Send(ctx, w, http.StatusBadRequest, CantDecodeBodyIntoCriteria, nil, err)
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
					response.Send(ctx, w, http.StatusInternalServerError, CantReEnqueueFailedMessage, nil, err)
					return
				}
			}

			response.Send(ctx, w, http.StatusInternalServerError, FailedToRunScrapper, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Scrapper successfully executed", nil, nil)
	}
}

// stop deferred function that handles the webDriver.Quit method
func stop(ctx context.Context, webDriverManager webdriver.Manager) {
	err := webDriverManager.Quit(ctx)
	if err != nil {
		log.Warn(ctx, err.Error())
	}
}
