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
	mockLogin := auth.MockLogin(nil)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockGetAdvanceSearchCriteria(mockCriteria)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

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

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

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

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

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

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria, mockExecuteAdvanceSearch, mockRetrieveAllTweets)

	assert.Equal(t, want, got)
}
