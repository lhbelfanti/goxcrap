package corpuscreator

import (
	"context"
	"fmt"

	"goxcrap/internal/http"
	"goxcrap/internal/log"
)

type (
	// UpdateSearchCriteriaExecution calls the endpoint in charge of updating a search criteria execution seeking by its execution id
	UpdateSearchCriteriaExecution func(ctx context.Context, executionID int, body UpdateSearchCriteriaExecutionBody) error

	// InsertSearchCriteriaExecutionDay calls the endpoint in charge of inserting a new search criteria execution day into the database
	InsertSearchCriteriaExecutionDay func(ctx context.Context, executionID int, body InsertSearchCriteriaExecutionDayBody) error
)

// MakeUpdateSearchCriteriaExecution creates a new UpdateSearchCriteriaExecution
func MakeUpdateSearchCriteriaExecution(httpClient http.Client, domain string) UpdateSearchCriteriaExecution {
	url := domain + "/criteria/executions/%d/v1"

	return func(ctx context.Context, executionID int, body UpdateSearchCriteriaExecutionBody) error {
		finalURL := fmt.Sprintf(url, executionID)

		resp, err := httpClient.NewRequest(ctx, "PUT", finalURL, body)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteRequest
		}

		log.Info(ctx, fmt.Sprintf("Update search criteria execution endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}

// MakeInsertSearchCriteriaExecutionDay creates a new InsertSearchCriteriaExecutionDay
func MakeInsertSearchCriteriaExecutionDay(httpClient http.Client, domain string) InsertSearchCriteriaExecutionDay {
	url := domain + "/criteria/executions/%d/day/v1"

	return func(ctx context.Context, executionID int, body InsertSearchCriteriaExecutionDayBody) error {
		finalURL := fmt.Sprintf(url, executionID)

		resp, err := httpClient.NewRequest(ctx, "POST", finalURL, body)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteRequest
		}

		log.Info(ctx, fmt.Sprintf("Insert search criteria execution day endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}
