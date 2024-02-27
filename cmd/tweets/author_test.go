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

func TestGetAuthor_success(t *testing.T) {
	want := "test"
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("Text", mock.Anything).Return("test", nil)

	getAuthor := tweets.MakeGetAuthor()

	got, err := getAuthor(mockWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestGetAuthor_failsWhenFindElementThrowsError(t *testing.T) {
	err := errors.New("error")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), err)

	getAuthor := tweets.MakeGetAuthor()

	want := tweets.NewTweetsError(tweets.FailedToObtainTweetAuthorElement, err)
	_, got := getAuthor(mockWebElement)

	assert.Equal(t, want, got)
}

func TestGetAuthor_failsWhenTextThrowsError(t *testing.T) {
	err := errors.New("error")
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)
	mockWantedWebElement.On("Text", mock.Anything).Return("", err)

	getAuthor := tweets.MakeGetAuthor()

	want := tweets.NewTweetsError(tweets.FailedToObtainTweetAuthor, err)
	_, got := getAuthor(mockWebElement)

	assert.Equal(t, want, got)
}
