package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
	"goxcrap/cmd/tweets"
)

func TestGetTimestamp_success(t *testing.T) {
	want := "test"
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("GetAttribute", mock.Anything).Return("test", nil)

	getTimestamp := tweets.MakeGetTimestamp()

	got, err := getTimestamp(mockWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestGetTimestamp_failsWhenFindElementThrowsError(t *testing.T) {
	err := errors.New("error")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), err)

	getTimestamp := tweets.MakeGetTimestamp()

	want := tweets.FailedToObtainTweetTimestampElement
	_, got := getTimestamp(mockWebElement)

	assert.Equal(t, want, got)
}

func TestGetTimestamp_failsWhenGetAttributeThrowsError(t *testing.T) {
	err := errors.New("error")
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("GetAttribute", mock.Anything).Return("", err)

	getTimestamp := tweets.MakeGetTimestamp()

	want := tweets.FailedToObtainTweetTimestamp
	_, got := getTimestamp(mockWebElement)

	assert.Equal(t, want, got)
}
