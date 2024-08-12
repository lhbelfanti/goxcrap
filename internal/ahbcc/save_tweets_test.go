package ahbcc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goxcrap/internal/ahbcc"
	"goxcrap/internal/http"
)

func TestSaveTweets_success(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	resp := http.Response{
		Status: "200 OK",
		Body:   `{"test": "body"}`,
	}
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	mockSaveTweetsBody := ahbcc.MockSaveTweetsBody()
	mockCtx := context.Background()
	saveTweets := ahbcc.MakeSaveTweets(mockHTTPClient, "http://example.com")

	got := saveTweets(mockCtx, mockSaveTweetsBody)

	assert.Nil(t, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestSaveTweets_failsWhenNewRequestThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(http.Response{}, errors.New("failed to execute NewRequest"))
	mockSaveTweetsBody := ahbcc.MockSaveTweetsBody()
	mockCtx := context.Background()
	saveTweets := ahbcc.MakeSaveTweets(mockHTTPClient, "http://example.com")

	want := ahbcc.FailedToExecuteRequest
	got := saveTweets(mockCtx, mockSaveTweetsBody)

	assert.Equal(t, want, got)
	mockHTTPClient.AssertExpectations(t)
}
