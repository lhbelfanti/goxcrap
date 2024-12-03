package corpuscreator

import (
	"context"
	"encoding/json"
	"fmt"

	"goxcrap/internal/http"
	"goxcrap/internal/log"
)

type (
	// GetSearchCriteriaExecution calls the endpoint in charge of retrieving a search criteria execution seeking by its execution id
	GetSearchCriteriaExecution func(ctx context.Context, executionID int) (Execution, error)

	// UpdateSearchCriteriaExecution calls the endpoint in charge of updating a search criteria execution seeking by its execution id
	UpdateSearchCriteriaExecution func(ctx context.Context, executionID int, body UpdateSearchCriteriaExecutionBody) error

	// InsertSearchCriteriaExecutionDay calls the endpoint in charge of inserting a new search criteria execution day into the database
	InsertSearchCriteriaExecutionDay func(ctx context.Context, executionID int, body InsertSearchCriteriaExecutionDayBody) error
)

// MakeGetSearchCriteriaExecution creates a new GetSearchCriteriaExecution
func MakeGetSearchCriteriaExecution(httpClient http.Client, domain string, localMode bool) GetSearchCriteriaExecution {
	url := domain + "/criteria/executions/%d/v1"

	return func(ctx context.Context, executionID int) (Execution, error) {
		if localMode {
			return Execution{}, nil
		}

		finalURL := fmt.Sprintf(url, executionID)

		resp, err := httpClient.NewRequest(ctx, "GET", finalURL, nil)
		if err != nil {
			log.Error(ctx, err.Error())
			return Execution{}, FailedToExecuteRequest
		}

		var execution Execution
		err = json.Unmarshal([]byte(resp.Body), &execution)
		if err != nil {
			log.Error(ctx, err.Error())
			return Execution{}, FailedToUnmarshalResponse
		}

		log.Info(ctx, fmt.Sprintf("Get search criteria execution endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return execution, nil
	}
}

// MakeUpdateSearchCriteriaExecution creates a new UpdateSearchCriteriaExecution
func MakeUpdateSearchCriteriaExecution(httpClient http.Client, domain string, localMode bool) UpdateSearchCriteriaExecution {
	url := domain + "/criteria/executions/%d/v1"

	return func(ctx context.Context, executionID int, body UpdateSearchCriteriaExecutionBody) error {
		if localMode {
			return nil
		}

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
func MakeInsertSearchCriteriaExecutionDay(httpClient http.Client, domain string, localMode bool) InsertSearchCriteriaExecutionDay {
	url := domain + "/criteria/executions/%d/day/v1"

	return func(ctx context.Context, executionID int, body InsertSearchCriteriaExecutionDayBody) error {
		if localMode {
			return nil
		}

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
