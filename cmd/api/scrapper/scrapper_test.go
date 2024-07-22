package scrapper_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/cmd/api/tweets"
)

func TestExecute_success(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()

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
	mockCriteria := criteria.MockCriteria()

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(mockCriteria, 0)

	assert.Nil(t, got)
}

func TestExecute_successWhenRetrieveAllThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockRetrieveAllTweets := tweets.MockRetrieveAll(nil, errors.New("error while executing RetrieveAll"))
	mockCriteria := criteria.MockCriteria()

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(mockCriteria, 0)

	assert.Nil(t, got)
}

func TestExecute_failsWhenLoginThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(errors.New("error while executing login"))
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	want := scrapper.FailedToLogin
	got := executeScrapper(mockCriteria, 0)

	assert.Equal(t, want, got)
}

func TestExecute_failsWhileTryingToParseDatesFromTheGivenCriteriaThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()
	mockCriteria.Since = "error"

	executeScrapper := scrapper.MakeExecute(mockLogin, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	want := scrapper.FailedToParseDatesFromTheGivenCriteria
	got := executeScrapper(mockCriteria, 0)

	assert.Equal(t, want, got)
}
