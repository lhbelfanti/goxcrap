package scrapper_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	"goxcrap/internal/webdriver"
)

func TestSearchCriteriaMessageProcessor_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockMessage := criteria.MockMessageDTO()
	mockBody, _ := json.Marshal(mockMessage)

	messageProcessor := scrapper.MakeSearchCriteriaMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	got := messageProcessor(context.Background(), mockBody)

	assert.Nil(t, got)
}

func TestSearchCriteriaMessageProcessor_successEvenWhenWebDriverManagerQuitThrowsErrorBecauseItJustLogsTheError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(errors.New("error while executing WebDriverManager.Quit"))
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockMessage := criteria.MockMessageDTO()
	mockBody, _ := json.Marshal(mockMessage)

	messageProcessor := scrapper.MakeSearchCriteriaMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	got := messageProcessor(context.Background(), mockBody)

	assert.Nil(t, got)
}

func TestSearchCriteriaMessageProcessor_failsWhenBodyCantBeDecoded(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)

	messageProcessor := scrapper.MakeSearchCriteriaMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToDecodeBodyIntoCriteria
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}

func TestSearchCriteriaMessageProcessor_failsWhenExecuteThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockMessage := criteria.MockMessageDTO()
	mockBody, _ := json.Marshal(mockMessage)

	messageProcessor := scrapper.MakeSearchCriteriaMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToRunScrapperProcess
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}

func TestSearchCriteriaMessageProcessor_failsWhenExecuteThrowsSpecificErrors(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(scrapper.FailedToLogin)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockMessage := criteria.MockMessageDTO()
	mockBody, _ := json.Marshal(mockMessage)

	messageProcessor := scrapper.MakeSearchCriteriaMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToRunScrapperProcess
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}

func TestSearchCriteriaMessageProcessor_failsWhenExecuteThrowsSpecificErrorsAndTheEnqueueFails(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(scrapper.FailedToLogin)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(errors.New("error while re enqueuing message"))
	mockMessage := criteria.MockMessageDTO()
	mockBody, _ := json.Marshal(mockMessage)

	messageProcessor := scrapper.MakeSearchCriteriaMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToReEnqueueFailedMessage
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}