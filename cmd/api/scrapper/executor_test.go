package scrapper_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/cmd/api/tweets"
	"goxcrap/internal/corpuscreator"
)

func TestExecute_success(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockUpdateSearchCriteriaExecution := corpuscreator.MockUpdateSearchCriteriaExecution(nil)
	mockInsertSearchCriteriaExecutionDay := corpuscreator.MockInsertSearchCriteriaExecutionDay(nil)
	mockTweet := tweets.MockTweet()
	mockTweet.Quote = tweets.MockQuote(true, true, true, "test", []string{"test"})
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()
	mockSaveTweets := corpuscreator.MockSaveTweets(nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockUpdateSearchCriteriaExecution, mockInsertSearchCriteriaExecutionDay, mockExecuteAdvanceSearch, mockRetrieveAllTweets, mockSaveTweets)

	got := executeScrapper(context.Background(), mockCriteria, 1)

	assert.Nil(t, got)
}

func TestExecute_successSkippingCriteriaDueAnErrorInExecuteAdvanceSearch(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(errors.New("error while executing ExecuteAdvanceSearch"))
	mockUpdateSearchCriteriaExecution := corpuscreator.MockUpdateSearchCriteriaExecution(nil)
	mockInsertSearchCriteriaExecutionDay := corpuscreator.MockInsertSearchCriteriaExecutionDay(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()
	mockSaveTweets := corpuscreator.MockSaveTweets(nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockUpdateSearchCriteriaExecution, mockInsertSearchCriteriaExecutionDay, mockExecuteAdvanceSearch, mockRetrieveAllTweets, mockSaveTweets)

	got := executeScrapper(context.Background(), mockCriteria, 1)

	assert.Nil(t, got)
}

func TestExecute_successWhenRetrieveAllThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockUpdateSearchCriteriaExecution := corpuscreator.MockUpdateSearchCriteriaExecution(nil)
	mockInsertSearchCriteriaExecutionDay := corpuscreator.MockInsertSearchCriteriaExecutionDay(nil)
	mockRetrieveAllTweets := tweets.MockRetrieveAll(nil, errors.New("error while executing RetrieveAll"))
	mockCriteria := criteria.MockCriteria()
	mockSaveTweets := corpuscreator.MockSaveTweets(nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockUpdateSearchCriteriaExecution, mockInsertSearchCriteriaExecutionDay, mockExecuteAdvanceSearch, mockRetrieveAllTweets, mockSaveTweets)

	got := executeScrapper(context.Background(), mockCriteria, 1)

	assert.Nil(t, got)
}

func TestExecute_successWhenSaveTweetsThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockUpdateSearchCriteriaExecution := corpuscreator.MockUpdateSearchCriteriaExecution(nil)
	mockInsertSearchCriteriaExecutionDay := corpuscreator.MockInsertSearchCriteriaExecutionDay(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()
	mockSaveTweets := corpuscreator.MockSaveTweets(errors.New("error while executing RetrieveAll"))

	executeScrapper := scrapper.MakeExecute(mockLogin, mockUpdateSearchCriteriaExecution, mockInsertSearchCriteriaExecutionDay, mockExecuteAdvanceSearch, mockRetrieveAllTweets, mockSaveTweets)

	got := executeScrapper(context.Background(), mockCriteria, 1)

	assert.Nil(t, got)
}

func TestExecute_failsWhenLoginThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(errors.New("error while executing login"))
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockUpdateSearchCriteriaExecution := corpuscreator.MockUpdateSearchCriteriaExecution(nil)
	mockInsertSearchCriteriaExecutionDay := corpuscreator.MockInsertSearchCriteriaExecutionDay(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()
	mockSaveTweets := corpuscreator.MockSaveTweets(nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockUpdateSearchCriteriaExecution, mockInsertSearchCriteriaExecutionDay, mockExecuteAdvanceSearch, mockRetrieveAllTweets, mockSaveTweets)

	want := scrapper.FailedToLogin
	got := executeScrapper(context.Background(), mockCriteria, 1)

	assert.Equal(t, want, got)
}

func TestExecute_failsWhenUpdateSearchCriteriaExecutionThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockUpdateSearchCriteriaExecution := corpuscreator.MockUpdateSearchCriteriaExecution(errors.New("error while executing UpdateSearchCriteriaExecution"))
	mockInsertSearchCriteriaExecutionDay := corpuscreator.MockInsertSearchCriteriaExecutionDay(nil)
	mockTweet := tweets.MockTweet()
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()
	mockSaveTweets := corpuscreator.MockSaveTweets(nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockUpdateSearchCriteriaExecution, mockInsertSearchCriteriaExecutionDay, mockExecuteAdvanceSearch, mockRetrieveAllTweets, mockSaveTweets)

	want := scrapper.FailedToUpdateSearchCriteriaExecution
	got := executeScrapper(context.Background(), mockCriteria, 1)

	assert.Equal(t, want, got)
}

func TestExecute_failsWhileTryingToParseDatesFromTheGivenCriteriaThrowsError(t *testing.T) {
	mockLogin := auth.MockLogin(nil)
	mockExecuteAdvanceSearch := search.MockExecuteAdvanceSearch(nil)
	mockTweet := tweets.MockTweet()
	mockUpdateSearchCriteriaExecution := corpuscreator.MockUpdateSearchCriteriaExecution(nil)
	mockInsertSearchCriteriaExecutionDay := corpuscreator.MockInsertSearchCriteriaExecutionDay(nil)
	mockRetrieveAllTweets := tweets.MockRetrieveAll([]tweets.Tweet{mockTweet, mockTweet}, nil)
	mockCriteria := criteria.MockCriteria()
	mockCriteria.Since = "error"
	mockSaveTweets := corpuscreator.MockSaveTweets(nil)

	executeScrapper := scrapper.MakeExecute(mockLogin, mockUpdateSearchCriteriaExecution, mockInsertSearchCriteriaExecutionDay, mockExecuteAdvanceSearch, mockRetrieveAllTweets, mockSaveTweets)

	want := scrapper.FailedToParseDatesFromTheGivenCriteria
	got := executeScrapper(context.Background(), mockCriteria, 1)

	assert.Equal(t, want, got)
}
