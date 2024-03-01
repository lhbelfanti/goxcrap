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
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("GetAttribute", mock.Anything).Return("test", nil)

	getTimestamp := tweets.MakeGetTimestamp()

	want := "test"
	got, err := getTimestamp(mockWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestGetTimestamp_failsWhenFindElementThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getTimestamp := tweets.MakeGetTimestamp()

	want := tweets.FailedToObtainTweetTimestampElement
	_, got := getTimestamp(mockWebElement)

	assert.Equal(t, want, got)
}

func TestGetTimestamp_failsWhenGetAttributeThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("GetAttribute", mock.Anything).Return("", errors.New("error while executing GetAttribute"))

	getTimestamp := tweets.MakeGetTimestamp()

	want := tweets.FailedToObtainTweetTimestamp
	_, got := getTimestamp(mockWebElement)

	assert.Equal(t, want, got)
}
