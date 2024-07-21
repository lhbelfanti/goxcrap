package page_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"goxcrap/cmd/api/page"
	"goxcrap/internal/webdriver"
)

func TestLoad_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(nil)
	mockWebDriver.On("Get", mock.Anything).Return(nil)

	load := page.MakeLoad(mockWebDriver)

	got := load("/test", 10*time.Minute)

	assert.Nil(t, got)
}

func TestLoad_failsWhenSetPageLoadTimeoutThrowsError(t *testing.T) {
	err := errors.New("error while executing driver.SetPageLoadTimeout")
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(err)

	load := page.MakeLoad(mockWebDriver)

	want := page.FailedToSetPageLoadTimeout
	got := load("/test", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestLoad_failsWhenGetThrowsError(t *testing.T) {
	err := errors.New("error while executing driver.Get")
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(nil)
	mockWebDriver.On("Get", mock.Anything).Return(err)

	load := page.MakeLoad(mockWebDriver)

	want := page.FailedToRetrievePage
	got := load("/test", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestScroll_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", mock.Anything, mock.Anything).Return(nil, nil)

	scroll := page.MakeScroll(mockWebDriver)

	got := scroll()

	assert.Nil(t, got)
}

func TestScroll_failsWhenJSHeightCodeExecutionThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `return window.innerHeight;`, mock.Anything).Return(100, errors.New("error while executing first code"))

	scroll := page.MakeScroll(mockWebDriver)

	want := page.FailedToGetInnerHeight
	got := scroll()

	assert.Equal(t, want, got)
}

func TestScroll_failsWhenScrollByCodeExecutionThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `return window.innerHeight;`, mock.Anything).Return(100, nil)
	mockWebDriver.On("ExecuteScript", `window.scrollBy(0, 100 * 2);`, mock.Anything).Return(nil, errors.New("error while executing second code"))

	scroll := page.MakeScroll(mockWebDriver)

	want := page.FailedToScroll
	got := scroll()

	assert.Equal(t, want, got)
}
