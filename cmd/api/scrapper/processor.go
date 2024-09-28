package scrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"

	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	"goxcrap/internal/log"
	"goxcrap/internal/webdriver"
)

// MakeSearchCriteriaMessageProcessor creates a new broker.ProcessorFunction
func MakeSearchCriteriaMessageProcessor(newWebDriverManager webdriver.NewManager, newScrapper New, messageBroker broker.MessageBroker) broker.ProcessorFunction {
	return func(ctx context.Context, body []byte) error {
		bodyBuffer := new(bytes.Buffer)
		teeReader := io.TeeReader(bytes.NewReader(body), bodyBuffer)

		var dto criteria.MessageDTO
		err := json.NewDecoder(teeReader).Decode(&dto)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToDecodeBodyIntoCriteria
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
					return FailedToReEnqueueFailedMessage
				}
			}

			log.Error(ctx, err.Error())
			return FailedToRunScrapperProcess
		}

		log.Info(ctx, "Scrapper successfully executed")
		return nil
	}
}
