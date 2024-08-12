package scrapper_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/broker"
	"goxcrap/internal/webdriver"
)

func TestExecuteHandlerV1_success(t *testing.T) {
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
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/scrapper/execute/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_successEvenWhenWebDriverManagerQuitThrowsErrorBecauseItJustLogsTheError(t *testing.T) {
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
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/scrapper/execute/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenHandlerThrowsInvalidBody(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit", mock.Anything).Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockMessageBroker := new(broker.MockMessageBroker)
	mockMessageBroker.On("EnqueueMessage", mock.Anything, mock.Anything).Return(nil)
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/scrapper/execute/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenExecuteThrowsError(t *testing.T) {
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
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/scrapper/execute/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenExecuteThrowsSpecificErrors(t *testing.T) {
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
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/scrapper/execute/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenExecuteThrowsSpecificErrorsAndTheEnqueueFails(t *testing.T) {
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
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/scrapper/execute/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper, mockMessageBroker)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
