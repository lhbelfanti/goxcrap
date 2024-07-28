package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/tweets"
)

func TestGetAuthor_success(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetAuthorWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetAuthorWebElement), nil)
	mockTweetAuthorWebElement.On("Text", mock.Anything).Return("test", nil)

	getAuthor := tweets.MakeGetAuthor()

	want := "test"
	got, err := getAuthor(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetAuthorWebElement.AssertExpectations(t)
}

func TestGetAuthor_failsWhenFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getAuthor := tweets.MakeGetAuthor()

	want := tweets.FailedToObtainTweetAuthorElement
	_, got := getAuthor(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
}

func TestGetAuthor_failsWhenTextThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetAuthorWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetAuthorWebElement), nil)
	mockTweetAuthorWebElement.On("Text", mock.Anything).Return("", errors.New("error while executing Text"))

	getAuthor := tweets.MakeGetAuthor()

	want := tweets.FailedToObtainTweetAuthor
	_, got := getAuthor(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetAuthorWebElement.AssertExpectations(t)
}
