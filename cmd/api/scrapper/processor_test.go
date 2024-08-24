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

func TestMessageProcessor_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)

	messageProcessor := scrapper.MakeMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	got := messageProcessor(context.Background(), mockBody)

	assert.Nil(t, got)
}

func TestMessageProcessor_successEvenWhenWebDriverManagerQuitThrowsErrorBecauseItJustLogsTheError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(errors.New("error while executing WebDriverManager.Quit"))
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)

	messageProcessor := scrapper.MakeMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	got := messageProcessor(context.Background(), mockBody)

	assert.Nil(t, got)
}

func TestMessageProcessor_failsWhenBodyCantBeDecoded(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)

	messageProcessor := scrapper.MakeMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToDecodeBodyIntoCriteria
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}

func TestMessageProcessor_failsWhenExecuteThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)

	messageProcessor := scrapper.MakeMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToRunScrapperProcess
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}

func TestMessageProcessor_failsWhenExecuteThrowsSpecificErrors(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(scrapper.FailedToLogin)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)

	messageProcessor := scrapper.MakeMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToRunScrapperProcess
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}

func TestMessageProcessor_failsWhenExecuteThrowsSpecificErrorsAndTheEnqueueFails(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(scrapper.FailedToLogin)
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(errors.New("error while re enqueuing message"))
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)

	messageProcessor := scrapper.MakeMessageProcessor(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	want := scrapper.FailedToReEnqueueFailedMessage
	got := messageProcessor(context.Background(), mockBody)

	assert.Equal(t, want, got)
}
