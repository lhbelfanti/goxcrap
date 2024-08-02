package ahbcc_test

import (
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
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	mockSaveTweetsBody := ahbcc.MockSaveTweetsBody()
	saveTweets := ahbcc.MakeSaveTweets(mockHTTPClient, "http://example.com")

	got := saveTweets(mockSaveTweetsBody)

	assert.Nil(t, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestSaveTweets_failsWhenNewRequestThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything).Return(http.Response{}, errors.New("failed to execute NewRequest"))
	mockSaveTweetsBody := ahbcc.MockSaveTweetsBody()
	saveTweets := ahbcc.MakeSaveTweets(mockHTTPClient, "http://example.com")

	want := ahbcc.FailedToExecuteRequest
	got := saveTweets(mockSaveTweetsBody)

	assert.Equal(t, want, got)
	mockHTTPClient.AssertExpectations(t)
}
