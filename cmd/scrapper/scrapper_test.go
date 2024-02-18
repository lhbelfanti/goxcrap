package scrapper_test

import (
	"errors"
	"goxcrap/cmd/tweets"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/scrapper"
	"goxcrap/cmd/search"
)

func TestExecute_success(t *testing.T) {
	mockLogin := auth.MockMakeLogin(nil)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockMakeGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockMakeExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockMakeRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	assert.Nil(t, got)
}

func TestExecute_successSkippingCriteriaDueAnErrorInParseDates(t *testing.T) {
	mockLogin := auth.MockMakeLogin(nil)
	mockCriteria := search.MockCriteria()
	mockCriteria[0].Since = "error"
	mockGetSearchCriteria := search.MockMakeGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockMakeExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockMakeRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	assert.Nil(t, got)
}

func TestExecute_successSkippingCriteriaDueAnErrorInExecuteAdvanceSearch(t *testing.T) {
	mockLogin := auth.MockMakeLogin(nil)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockMakeGetAdvanceSearchCriteria(mockCriteria)
	err := errors.New("error while executing ExecuteAdvanceSearch")
	mockExecuteAdvanceSearch := search.MockMakeExecuteAdvanceSearch(err)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockMakeRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	assert.Nil(t, got)
}

func TestExecute_failsWhenLoginThrowsError(t *testing.T) {
	want := errors.New("error while executing login")
	mockLogin := auth.MockMakeLogin(want)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockMakeGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockMakeExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockMakeRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	assert.Equal(t, want, got)
}
