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
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/scrapper"
	"goxcrap/cmd/api/search"
	"goxcrap/internal/driver"
)

func TestExecuteHandlerV1_success(t *testing.T) {
	mockGoXCrapWebDriver := new(driver.MockGoXCrapWebDriver)
	mockSeleniumService := &selenium.Service{}
	mockWebDriver := new(driver.MockWebDriver)
	mockNewWebDriver := driver.MockNew(mockGoXCrapWebDriver, mockSeleniumService, mockWebDriver)
	mockNewScrapper := scrapper.MockNew(nil)
	mockCriteria := search.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/execute-scrapper/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriver, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenHandlerThrowsInvalidBody(t *testing.T) {
	mockGoXCrapWebDriver := new(driver.MockGoXCrapWebDriver)
	mockSeleniumService := &selenium.Service{}
	mockWebDriver := new(driver.MockWebDriver)
	mockNewWebDriver := driver.MockNew(mockGoXCrapWebDriver, mockSeleniumService, mockWebDriver)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/execute-scrapper/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriver, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenExecuteThrowsError(t *testing.T) {
	mockGoXCrapWebDriver := new(driver.MockGoXCrapWebDriver)
	mockSeleniumService := &selenium.Service{}
	mockWebDriver := new(driver.MockWebDriver)
	mockNewWebDriver := driver.MockNew(mockGoXCrapWebDriver, mockSeleniumService, mockWebDriver)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockCriteria := search.MockCriteria()
	mockBody, _ := json.Marshal(mockCriteria)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/execute-scrapper/v1", bytes.NewReader(mockBody))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriver, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}
