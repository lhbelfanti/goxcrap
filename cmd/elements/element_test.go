package elements_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
	"goxcrap/internal/chromedriver"
)

func TestWaitAndRetrieve_success(t *testing.T) {
	want := new(elements.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(nil)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(want), nil)
	mockWaitAndRetrieveCondition := elements.MockMakeWaitAndRetrieveCondition(true)

	waitAndRetrieve := elements.MakeWaitAndRetrieve(mockWebDriver, mockWaitAndRetrieveCondition)

	got, err := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestWaitAndRetrieve_failsWhenWaitWithTimeoutThrowsError(t *testing.T) {
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(errors.New("error"))
	mockWaitAndRetrieveCondition := elements.MockMakeWaitAndRetrieveCondition(true)

	waitAndRetrieve := elements.MakeWaitAndRetrieve(mockWebDriver, mockWaitAndRetrieveCondition)

	_, got := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.NotNil(t, got)
}

func TestWaitAndRetrieve_failsFindElementThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(nil)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWebElement), errors.New("error"))
	mockWaitAndRetrieveCondition := elements.MockMakeWaitAndRetrieveCondition(true)

	waitAndRetrieve := elements.MakeWaitAndRetrieve(mockWebDriver, mockWaitAndRetrieveCondition)

	_, got := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.NotNil(t, got)
}

func TestWaitAndRetrieveCondition_successWithReturnValueTrue(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("IsDisplayed").Return(true, nil)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWebElement), nil)

	waitAndRetrieveCondition := elements.MakeWaitAndRetrieveCondition()
	seleniumCondition := waitAndRetrieveCondition(selenium.ByName, "value")

	got, err := seleniumCondition(mockWebDriver)

	assert.True(t, got)
	assert.Nil(t, err)
}

func TestWaitAndRetrieveCondition_successWithReturnValueFalseWhenFindElementsThrowsAnError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWebElement), errors.New("error"))

	waitAndRetrieveCondition := elements.MakeWaitAndRetrieveCondition()
	seleniumCondition := waitAndRetrieveCondition(selenium.ByName, "value")

	got, err := seleniumCondition(mockWebDriver)

	assert.False(t, got)
	assert.Nil(t, err)
}

// ---

func TestWaitAndRetrieveAll_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	want := []selenium.WebElement{selenium.WebElement(mockWebElement)}
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(nil)
	mockWebDriver.On("FindElements", mock.Anything, mock.Anything).Return(want, nil)
	mockWaitAndRetrieveAllCondition := elements.MockMakeWaitAndRetrieveAllCondition(true)

	waitAndRetrieve := elements.MakeWaitAndRetrieveAll(mockWebDriver, mockWaitAndRetrieveAllCondition)

	got, err := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestWaitAndRetrieveAll_failsWhenWaitWithTimeoutThrowsError(t *testing.T) {
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(errors.New("error"))
	mockWaitAndRetrieveAllCondition := elements.MockMakeWaitAndRetrieveAllCondition(true)

	waitAndRetrieve := elements.MakeWaitAndRetrieveAll(mockWebDriver, mockWaitAndRetrieveAllCondition)

	_, got := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.NotNil(t, got)
}

func TestWaitAndRetrieveAll_failsFindElementThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("WaitWithTimeout", mock.Anything, mock.Anything).Return(nil)
	mockWebDriver.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockWebElement)}, errors.New("error"))
	mockWaitAndRetrieveAllCondition := elements.MockMakeWaitAndRetrieveAllCondition(true)

	waitAndRetrieve := elements.MakeWaitAndRetrieveAll(mockWebDriver, mockWaitAndRetrieveAllCondition)

	_, got := waitAndRetrieve(selenium.ByName, "value", 10*time.Minute)

	assert.NotNil(t, got)
}

func TestWaitAndRetrieveAllCondition_successWithReturnValueTrue(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("IsDisplayed").Return(true, nil)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockWebElement)}, nil)

	waitAndRetrieveAllCondition := elements.MakeWaitAndRetrieveAllCondition()
	seleniumCondition := waitAndRetrieveAllCondition(selenium.ByName, "value")

	got, err := seleniumCondition(mockWebDriver)

	assert.True(t, got)
	assert.Nil(t, err)
}

func TestWaitAndRetrieveAllCondition_successWithReturnValueFalseWhenFindElementsThrowsAnError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockWebElement)}, errors.New("error"))

	waitAndRetrieveAllCondition := elements.MakeWaitAndRetrieveAllCondition()
	seleniumCondition := waitAndRetrieveAllCondition(selenium.ByName, "value")

	got, err := seleniumCondition(mockWebDriver)

	assert.False(t, got)
	assert.Nil(t, err)
}

func TestWaitAndRetrieveAllCondition_successWithReturnValueFalseWhenFindElementsReturnsALengthZeroSlice(t *testing.T) {
	mockWebDriver := new(chromedriver.MockWebDriver)
	mockWebDriver.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{}, nil)

	waitAndRetrieveAllCondition := elements.MakeWaitAndRetrieveAllCondition()
	seleniumCondition := waitAndRetrieveAllCondition(selenium.ByName, "value")

	got, err := seleniumCondition(mockWebDriver)

	assert.False(t, got)
	assert.Nil(t, err)
}
