package corpuscreator_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goxcrap/internal/corpuscreator"
	"goxcrap/internal/http"
)

func TestInsertSearchCriteriaExecution_success(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	resp := http.Response{
		Status: "200 OK",
		Body:   `{"test": "body"}`,
	}
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	insertSearchCriteriaExecution := corpuscreator.MakeInsertSearchCriteriaExecution(mockHTTPClient, "http://example.com")

	got := insertSearchCriteriaExecution(context.Background(), 1)

	assert.Nil(t, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestInsertSearchCriteriaExecution_failsWhenNewRequestThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(http.Response{}, errors.New("failed to execute NewRequest"))
	insertSearchCriteriaExecution := corpuscreator.MakeInsertSearchCriteriaExecution(mockHTTPClient, "http://example.com")

	want := corpuscreator.FailedToExecuteRequest
	got := insertSearchCriteriaExecution(context.Background(), 1)

	assert.Equal(t, want, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestUpdateSearchCriteriaExecution_success(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	resp := http.Response{
		Status: "200 OK",
		Body:   `{"test": "body"}`,
	}
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	mockUpdateSearchCriteriaExecutionBody := corpuscreator.MockUpdateSearchCriteriaExecutionBody()
	updateSearchCriteriaExecution := corpuscreator.MakeUpdateSearchCriteriaExecution(mockHTTPClient, "http://example.com")

	got := updateSearchCriteriaExecution(context.Background(), 1, mockUpdateSearchCriteriaExecutionBody)

	assert.Nil(t, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestUpdateSearchCriteriaExecution_failsWhenNewRequestThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(http.Response{}, errors.New("failed to execute NewRequest"))
	mockUpdateSearchCriteriaExecutionBody := corpuscreator.MockUpdateSearchCriteriaExecutionBody()
	updateSearchCriteriaExecution := corpuscreator.MakeUpdateSearchCriteriaExecution(mockHTTPClient, "http://example.com")

	want := corpuscreator.FailedToExecuteRequest
	got := updateSearchCriteriaExecution(context.Background(), 1, mockUpdateSearchCriteriaExecutionBody)

	assert.Equal(t, want, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestInsertSearchCriteriaExecutionDay_success(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	resp := http.Response{
		Status: "200 OK",
		Body:   `{"test": "body"}`,
	}
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	mockInsertSearchCriteriaExecutionDayBody := corpuscreator.MockInsertSearchCriteriaExecutionDayBody()
	insertSearchCriteriaExecutionDay := corpuscreator.MakeInsertSearchCriteriaExecutionDay(mockHTTPClient, "http://example.com")

	got := insertSearchCriteriaExecutionDay(context.Background(), 1, mockInsertSearchCriteriaExecutionDayBody)

	assert.Nil(t, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestInsertSearchCriteriaExecutionDay_failsWhenNewRequestThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(http.Response{}, errors.New("failed to execute NewRequest"))
	mockInsertSearchCriteriaExecutionDayBody := corpuscreator.MockInsertSearchCriteriaExecutionDayBody()
	insertSearchCriteriaExecutionDay := corpuscreator.MakeInsertSearchCriteriaExecutionDay(mockHTTPClient, "http://example.com")

	want := corpuscreator.FailedToExecuteRequest
	got := insertSearchCriteriaExecutionDay(context.Background(), 1, mockInsertSearchCriteriaExecutionDayBody)

	assert.Equal(t, want, got)
	mockHTTPClient.AssertExpectations(t)
}
