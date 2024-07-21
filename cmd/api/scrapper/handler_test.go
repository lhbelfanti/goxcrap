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

	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search/criteria"
	"goxcrap/internal/webdriver"
)

func TestExecuteHandlerV1_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit").Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/execute-scrapper/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_successEvenWhenWebDriverManagerQuitThrowsErrorBecauseItJustLogsTheError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit").Return(errors.New("error while executing WebDriverManager.Quit"))
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(nil)
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/execute-scrapper/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenHandlerThrowsInvalidBody(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit").Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/execute-scrapper/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenExecuteThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockManager := new(webdriver.MockManager)
	mockManager.On("WebDriver").Return(mockWebDriver)
	mockManager.On("Quit").Return(nil)
	mockNewWebDriverManager := webdriver.MockNewManager(mockManager)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockCriteria := criteria.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/execute-scrapper/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriverManager, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}
