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

func TestGetTweetTimestamp_success(t *testing.T) {
	want := "test"
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("GetAttribute", mock.Anything).Return("test", nil)

	getTweetTimestamp := tweets.MakeGetTweetTimestamp()

	got, err := getTweetTimestamp(mockWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestGetTweetTimestamp_failsWhenFindElementThrowsError(t *testing.T) {
	err := errors.New("error")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), err)

	getTweetTimestamp := tweets.MakeGetTweetTimestamp()

	want := tweets.NewTweetsError(tweets.FailedToObtainTweetTimestampElement, err)
	_, got := getTweetTimestamp(mockWebElement)

	assert.Equal(t, want, got)
}

func TestGetTweetTimestamp_failsWhenGetAttributeThrowsError(t *testing.T) {
	err := errors.New("error")
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("GetAttribute", mock.Anything).Return("", err)

	getTweetTimestamp := tweets.MakeGetTweetTimestamp()

	want := tweets.NewTweetsError(tweets.FailedToObtainTweetTimestamp, err)
	_, got := getTweetTimestamp(mockWebElement)

	assert.Equal(t, want, got)
}
