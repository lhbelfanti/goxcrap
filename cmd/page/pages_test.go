package page_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goxcrap/cmd/page"
	"goxcrap/internal/chromedriver"
)

func TestLoad_success(t *testing.T) {
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(nil)
	mockWebDriver.On("Get", mock.Anything).Return(nil)
	load := page.MakeLoad(mockWebDriver)

	got := load("/test", 10*time.Minute)

	assert.Nil(t, got)
}

func TestLoad_failsWhenSetPageLoadTimeoutThrowsError(t *testing.T) {
	err := errors.New("error while executing driver.SetPageLoadTimeout")
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(err)
	load := page.MakeLoad(mockWebDriver)

	want := page.NewPageError(page.FailedToSetPageLoadTimeout, err)
	got := load("/test", 10*time.Minute)

	assert.Equal(t, got, want)
}

func TestLoad_failsWhenGetThrowsError(t *testing.T) {
	err := errors.New("error while executing driver.Get")
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(nil)
	mockWebDriver.On("Get", mock.Anything).Return(err)
	load := page.MakeLoad(mockWebDriver)

	want := page.NewPageError(page.FailedToRetrievePage, err)
	got := load("/test", 10*time.Minute)

	assert.Equal(t, got, want)
}
