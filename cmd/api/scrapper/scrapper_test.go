package scrapper_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/tweets"
)

func TestExecute_success(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := search.MockCriteria()

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(mockCriteria, 0)

	assert.Nil(t, got)
}

func TestExecute_successSkippingCriteriaDueAnErrorInParseDates(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := search.MockCriteria()
	mockCriteria[0].Since = "error"

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(mockCriteria, 0)

	assert.Nil(t, got)
}

func TestExecute_successSkippingCriteriaDueAnErrorInExecuteAdvanceSearch(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	err := errors.New("error while executing ExecuteAdvanceSearch")
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(err)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := search.MockCriteria()

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(mockCriteria, 0)

	assert.Nil(t, got)
}

func TestExecute_successWhenRetrieveAllThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockRetrieveAllTweets := tweets.MockRetrieveAll(nil, errors.New("error while executing RetrieveAll"))
	mockCriteria := search.MockCriteria()

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(mockCriteria, 0)

	assert.Nil(t, got)
}

func TestExecute_failsWhenLoginThrowsError(t *testing.T) {
	want := errors.New("error while executing login")
	mockLogin := auth.MockLogin(want)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := search.MockCriteria()

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(mockCriteria, 0)

	assert.Equal(t, want, got)
}
