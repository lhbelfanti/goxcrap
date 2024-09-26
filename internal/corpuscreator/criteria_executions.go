package corpuscreator

import (
	"context"
	"fmt"

	"goxcrap/internal/http"
	"goxcrap/internal/log"
)

type (
	// InsertSearchCriteriaExecution calls the endpoint in charge of inserting a new search criteria execution into the database
	InsertSearchCriteriaExecution func(ctx context.Context, searchCriteriaID int) error

	// UpdateSearchCriteriaExecution calls the endpoint in charge of updating a search criteria execution seeking by its execution id
	UpdateSearchCriteriaExecution func(ctx context.Context, executionID int, body UpdateSearchCriteriaExecutionBody) error

	// InsertSearchCriteriaExecutionDay calls the endpoint in charge of inserting a new search criteria execution day into the database
	InsertSearchCriteriaExecutionDay func(ctx context.Context, executionID int, body InsertSearchCriteriaExecutionDayBody) error
)

// MakeInsertSearchCriteriaExecution creates a new InsertSearchCriteriaExecution
func MakeInsertSearchCriteriaExecution(httpClient http.Client, domain string) InsertSearchCriteriaExecution {
	url := domain + "/criteria/%d/executions/v1"

	return func(ctx context.Context, searchCriteriaID int) error {
		url = fmt.Sprintf(url, searchCriteriaID)

		resp, err := httpClient.NewRequest(ctx, "POST", url, nil)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteRequest
		}

		log.Info(ctx, fmt.Sprintf("Insert search criteria execution endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}

// MakeUpdateSearchCriteriaExecution creates a new UpdateSearchCriteriaExecution
func MakeUpdateSearchCriteriaExecution(httpClient http.Client, domain string) UpdateSearchCriteriaExecution {
	url := domain + "/criteria/executions/%d/v1"

	return func(ctx context.Context, executionID int, body UpdateSearchCriteriaExecutionBody) error {
		url = fmt.Sprintf(url, executionID)

		resp, err := httpClient.NewRequest(ctx, "PUT", url, body)
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
		url = fmt.Sprintf(url, executionID)

		resp, err := httpClient.NewRequest(ctx, "POST", url, body)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteRequest
		}

		log.Info(ctx, fmt.Sprintf("Insert search criteria execution day endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}
