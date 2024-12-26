package page_test

import (
	"context"
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

	got := load(context.Background(), "/test", 10*time.Minute)

	assert.Nil(t, got)
}

func TestLoad_failsWhenSetPageLoadTimeoutThrowsError(t *testing.T) {
	err := errors.New("error while executing driver.SetPageLoadTimeout")
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(err)

	load := page.MakeLoad(mockWebDriver)

	want := page.FailedToSetPageLoadTimeout
	got := load(context.Background(), "/test", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestLoad_failsWhenGetThrowsError(t *testing.T) {
	err := errors.New("error while executing driver.Get")
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("SetPageLoadTimeout", mock.Anything).Return(nil)
	mockWebDriver.On("Get", mock.Anything).Return(err)

	load := page.MakeLoad(mockWebDriver)

	want := page.FailedToRetrievePage
	got := load(context.Background(), "/test", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestScroll_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", mock.Anything, mock.Anything).Return(nil, nil)

	scroll := page.MakeScroll(mockWebDriver)

	got := scroll(context.Background())

	assert.Nil(t, got)
}

func TestScroll_failsWhenJSHeightCodeExecutionThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `return window.innerHeight;`, mock.Anything).Return(100, errors.New("error while executing first code"))

	scroll := page.MakeScroll(mockWebDriver)

	want := page.FailedToGetInnerHeight
	got := scroll(context.Background())

	assert.Equal(t, want, got)
}

func TestScroll_failsWhenScrollByCodeExecutionThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `return window.innerHeight;`, mock.Anything).Return(100, nil)
	mockWebDriver.On("ExecuteScript", `window.scrollBy(0, 100 * 2);`, mock.Anything).Return(nil, errors.New("error while executing second code"))

	scroll := page.MakeScroll(mockWebDriver)

	want := page.FailedToScroll
	got := scroll(context.Background())

	assert.Equal(t, want, got)
}

func TestOpenNewTab_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `window.open('', '_blank');`, mock.Anything).Return(nil, nil)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, nil)
	mockWebDriver.On("SwitchWindow", mock.Anything).Return(nil)
	mockLoadPage := page.MockLoad(nil)

	openNewTab := page.MakeOpenNewTab(mockWebDriver, mockLoadPage)

	got := openNewTab(context.Background(), "www.test.com", 0)

	assert.Nil(t, got)
}

func TestOpenNewTab_failsWhenWindowOpenScriptThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `window.open('', '_blank');`, mock.Anything).Return(nil, errors.New("error while executing code"))
	mockLoadPage := page.MockLoad(nil)

	openNewTab := page.MakeOpenNewTab(mockWebDriver, mockLoadPage)

	want := page.FailedToOpenNewTab
	got := openNewTab(context.Background(), "www.test.com", 0)

	assert.Equal(t, want, got)
}

func TestOpenNewTab_failsWhenWindowHandlesThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `window.open('', '_blank');`, mock.Anything).Return(nil, nil)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, errors.New("error while executing driver.WindowHandles"))
	mockLoadPage := page.MockLoad(nil)

	openNewTab := page.MakeOpenNewTab(mockWebDriver, mockLoadPage)

	want := page.FailedToObtainWindowHandles
	got := openNewTab(context.Background(), "www.test.com", 0)

	assert.Equal(t, want, got)
}

func TestOpenNewTab_failsWhenSwitchWindowThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `window.open('', '_blank');`, mock.Anything).Return(nil, nil)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, nil)
	mockWebDriver.On("SwitchWindow", mock.Anything).Return(errors.New("error while executing driver.SwitchWindow"))
	mockLoadPage := page.MockLoad(nil)

	openNewTab := page.MakeOpenNewTab(mockWebDriver, mockLoadPage)

	want := page.FailedToSwitchWindow
	got := openNewTab(context.Background(), "www.test.com", 0)

	assert.Equal(t, want, got)
}

func TestOpenNewTab_failsWhenLoadPageThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("ExecuteScript", `window.open('', '_blank');`, mock.Anything).Return(nil, nil)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, nil)
	mockWebDriver.On("SwitchWindow", mock.Anything).Return(nil)
	mockLoadPage := page.MockLoad(errors.New("error while executing page.Load"))

	openNewTab := page.MakeOpenNewTab(mockWebDriver, mockLoadPage)

	want := page.FailedToLoadPageOnTheNewTab
	got := openNewTab(context.Background(), "www.test.com", 0)

	assert.Equal(t, want, got)
}

func TestCloseOpenedTabs_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, nil)
	mockWebDriver.On("SwitchWindow", "b").Return(nil)
	mockWebDriver.On("Close").Return(nil)
	mockWebDriver.On("SwitchWindow", "a").Return(nil)

	closeOpenedTabs := page.MakeCloseOpenedTabs(mockWebDriver)

	got := closeOpenedTabs(context.Background())

	assert.Nil(t, got)
}

func TestCloseOpenedTabs_failsWhenWindowHandlesThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, errors.New("error while executing driver.WindowHandles"))

	closeOpenedTabs := page.MakeCloseOpenedTabs(mockWebDriver)

	want := page.FailedToObtainWindowHandles
	got := closeOpenedTabs(context.Background())

	assert.Equal(t, want, got)
}

func TestCloseOpenedTabs_failsWhenFirstCallToSwitchWindowThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, nil)
	mockWebDriver.On("SwitchWindow", "b").Return(errors.New("error while executing first driver.SwitchWindow"))

	closeOpenedTabs := page.MakeCloseOpenedTabs(mockWebDriver)

	want := page.FailedToSwitchWindow
	got := closeOpenedTabs(context.Background())

	assert.Equal(t, want, got)
}

func TestCloseOpenedTabs_failsWhenCloseThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, nil)
	mockWebDriver.On("SwitchWindow", "b").Return(nil)
	mockWebDriver.On("Close").Return(errors.New("error while executing driver.Close"))

	closeOpenedTabs := page.MakeCloseOpenedTabs(mockWebDriver)

	want := page.FailedToCloseWindow
	got := closeOpenedTabs(context.Background())

	assert.Equal(t, want, got)
}

func TestCloseOpenedTabs_failsWhenSecondSwitchWindowThrowsError(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockWebDriver.On("WindowHandles", mock.Anything).Return([]string{"a", "b"}, nil)
	mockWebDriver.On("SwitchWindow", "b").Return(nil)
	mockWebDriver.On("Close").Return(nil)
	mockWebDriver.On("SwitchWindow", "a").Return(errors.New("error while executing second driver.SwitchWindow"))

	closeOpenedTabs := page.MakeCloseOpenedTabs(mockWebDriver)

	want := page.FailedToSwitchToMainWindow
	got := closeOpenedTabs(context.Background())

	assert.Equal(t, want, got)
}
