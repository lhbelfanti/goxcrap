package element_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/element"
	"goxcrap/internal/chromedriver"
)

func TestWaitAndRetrieve_success(t *testing.T) {
	want := new(element.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(nil)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(want), nil)
	mockWaitAndRetrieveCondition := element.MockMakeWaitAndRetrieveCondition(true)

	waitAndRetrieve := element.MakeWaitAndRetrieve(mockWebDriver, mockWaitAndRetrieveCondition)

	got, err := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestWaitAndRetrieve_failsWhenWaitWithTimeoutThrowsError(t *testing.T) {
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(errors.New("error"))
	mockWaitAndRetrieveCondition := element.MockMakeWaitAndRetrieveCondition(true)

	waitAndRetrieve := element.MakeWaitAndRetrieve(mockWebDriver, mockWaitAndRetrieveCondition)

	_, got := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.NotNil(t, got)
}

func TestWaitAndRetrieve_failsFindElementThrowsError(t *testing.T) {
	mockWebElement := new(element.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(nil)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWebElement), errors.New("error"))
	mockWaitAndRetrieveCondition := element.MockMakeWaitAndRetrieveCondition(true)

	waitAndRetrieve := element.MakeWaitAndRetrieve(mockWebDriver, mockWaitAndRetrieveCondition)

	_, got := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.NotNil(t, got)
}

func TestWaitAndRetrieveCondition_successWithReturnValueTrue(t *testing.T) {
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("IsDisplayed").Return(true, nil)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWebElement), nil)

	waitAndRetrieveCondition := element.MakeWaitAndRetrieveCondition()
	seleniumCondition := waitAndRetrieveCondition(selenium.ByName, "value")

	got, err := seleniumCondition(mockWebDriver)

	assert.True(t, got)
	assert.Nil(t, err)
}

func TestWaitAndRetrieveCondition_successWithReturnValueFalse(t *testing.T) {
	mockWebElement := new(element.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWebElement), errors.New("error"))

	waitAndRetrieveCondition := element.MakeWaitAndRetrieveCondition()
	seleniumCondition := waitAndRetrieveCondition(selenium.ByName, "value")

	got, err := seleniumCondition(mockWebDriver)

	assert.False(t, got)
	assert.Nil(t, err)
}
