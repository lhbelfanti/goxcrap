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
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(0)

	assert.Nil(t, got)
}

func TestExecute_successSkippingCriteriaDueAnErrorInParseDates(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockCriteria := search.MockCriteria()
	mockCriteria[0].Since = "error"
	mockGetSearchCriteria := search.MockGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(0)

	assert.Nil(t, got)
}

func TestExecute_successSkippingCriteriaDueAnErrorInExecuteAdvanceSearch(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockGetAdvanceSearchCriteria(mockCriteria)
	err := errors.New("error while executing ExecuteAdvanceSearch")
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(err)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(0)

	assert.Nil(t, got)
}

func TestExecute_successWhenRetrieveAllThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockRetrieveAllTweets := tweets.MockRetrieveAll(nil, errors.New("error while executing RetrieveAll"))

	executeScrapper := scrapper.MakeExecute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(0)

	assert.Nil(t, got)
}

func TestExecute_failsWhenLoginThrowsError(t *testing.T) {
	want := errors.New("error while executing login")
	mockLogin := auth.MockLogin(want)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	got := executeScrapper(0)

	assert.Equal(t, want, got)
}
