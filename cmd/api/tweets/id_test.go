package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/tweets"
)

func TestGetID_success(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDWebElement := new(elements.MockWebElement)
	mockTweetIDATagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDWebElement), nil)
	mockTweetIDWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDATagWebElement), nil)
	mockTweetIDATagWebElement.On("GetAttribute", mock.Anything).Return("id/123456789", nil)

	getID := tweets.MakeGetID()

	want := "123456789"
	got, err := getID(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetIDWebElement.AssertExpectations(t)
	mockTweetIDATagWebElement.AssertExpectations(t)
}

func TestGetID_failsWhenFirstFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getID := tweets.MakeGetID()

	want := tweets.FailedToObtainTweetIDElement
	_, got := getID(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
}

func TestGetID_failsWhenSecondFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDWebElement), nil)
	mockTweetIDWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getID := tweets.MakeGetID()

	want := tweets.FailedToObtainTweetIDATag
	_, got := getID(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetIDWebElement.AssertExpectations(t)
}

func TestGetID_failsWhenGetAttributeThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDWebElement := new(elements.MockWebElement)
	mockTweetIDATagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDWebElement), nil)
	mockTweetIDWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDATagWebElement), nil)
	mockTweetIDATagWebElement.On("GetAttribute", mock.Anything).Return("", errors.New("error while executing GetAttribute"))

	getID := tweets.MakeGetID()

	want := tweets.FailedToObtainTweetIDATagHref
	_, got := getID(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetIDWebElement.AssertExpectations(t)
	mockTweetIDATagWebElement.AssertExpectations(t)
}
