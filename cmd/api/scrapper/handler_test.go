package scrapper_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/scrapper"
	"goxcrap/internal/driver"
)

func TestExecuteHandlerV1_success(t *testing.T) {
	mockGoXCrapWebDriver := new(driver.MockGoXCrapWebDriver)
	mockSeleniumService := &selenium.Service{}
	mockWebDriver := new(driver.MockWebDriver)
	mockNewWebDriver := driver.MockNew(mockGoXCrapWebDriver, mockSeleniumService, mockWebDriver)
	mockNewScrapper := scrapper.MockNew(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/execute-scrapper/v1", strings.NewReader(""))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriver, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestExecuteHandlerV1_failsWhenExecuteThrowsError(t *testing.T) {
	mockGoXCrapWebDriver := new(driver.MockGoXCrapWebDriver)
	mockSeleniumService := &selenium.Service{}
	mockWebDriver := new(driver.MockWebDriver)
	mockNewWebDriver := driver.MockNew(mockGoXCrapWebDriver, mockSeleniumService, mockWebDriver)
	mockNewScrapper := scrapper.MockNew(errors.New("execute scrapper failed"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/execute-scrapper/v1", strings.NewReader(""))

	handlerV1 := scrapper.ExecuteHandlerV1(mockNewWebDriver, mockNewScrapper)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}
