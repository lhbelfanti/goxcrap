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

func TestGetTimestamp_success(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetTimestampWebElement := new(elements.MockWebElement)
	mockTweetTimestampTimeTagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTimestampWebElement), nil)
	mockTweetTimestampWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTimestampTimeTagWebElement), nil)
	mockTweetTimestampTimeTagWebElement.On("GetAttribute", mock.Anything).Return("test", nil)

	getTimestamp := tweets.MakeGetTimestamp()

	want := "test"
	got, err := getTimestamp(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetTimestampWebElement.AssertExpectations(t)
	mockTweetTimestampTimeTagWebElement.AssertExpectations(t)
}

func TestGetTimestamp_failsWhenFirstFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getTimestamp := tweets.MakeGetTimestamp()

	want := tweets.FailedToObtainTweetTimestampElement
	_, got := getTimestamp(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
}

func TestGetTimestamp_failsWhenSecondFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetTimestampWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTimestampWebElement), nil)
	mockTweetTimestampWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getTimestamp := tweets.MakeGetTimestamp()

	want := tweets.FailedToObtainTweetTimestampTimeTag
	_, got := getTimestamp(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetTimestampWebElement.AssertExpectations(t)
}

func TestGetTimestamp_failsWhenGetAttributeThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetTimestampWebElement := new(elements.MockWebElement)
	mockTweetTimestampTimeTagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTimestampWebElement), nil)
	mockTweetTimestampWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTimestampTimeTagWebElement), nil)
	mockTweetTimestampTimeTagWebElement.On("GetAttribute", mock.Anything).Return("", errors.New("error while executing GetAttribute"))

	getTimestamp := tweets.MakeGetTimestamp()

	want := tweets.FailedToObtainTweetTimestamp
	_, got := getTimestamp(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetTimestampWebElement.AssertExpectations(t)
	mockTweetTimestampTimeTagWebElement.AssertExpectations(t)
}

func TestGetQuoteTimestamp_success(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuoteTimestampWebElement := new(elements.MockWebElement)
		mockQuoteTimestampTimeTagWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuoteTimestampWebElement), nil)
		mockQuoteTimestampWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuoteTimestampTimeTagWebElement), nil)
		mockQuoteTimestampTimeTagWebElement.On("GetAttribute", mock.Anything).Return("test", nil)

		getQuoteTimestamp := tweets.MakeGetQuoteTimestamp()

		want := "test"
		got, err := getQuoteTimestamp(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuoteTimestampWebElement.AssertExpectations(t)
		mockQuoteTimestampTimeTagWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteTimestamp_failsWhenFirstFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

		getQuoteTimestamp := tweets.MakeGetQuoteTimestamp()

		want := tweets.FailedToObtainQuotedTweetTimestampElement
		_, got := getQuoteTimestamp(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteTimestamp_failsWhenSecondFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuoteTimestampWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuoteTimestampWebElement), nil)
		mockQuoteTimestampWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

		getQuoteTimestamp := tweets.MakeGetQuoteTimestamp()

		want := tweets.FailedToObtainQuotedTweetTimestampTimeTag
		_, got := getQuoteTimestamp(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuoteTimestampWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteTimestamp_failsWhenGetAttributeThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuoteTimestampWebElement := new(elements.MockWebElement)
		mockQuoteTimestampTimeTagWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuoteTimestampWebElement), nil)
		mockQuoteTimestampWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuoteTimestampTimeTagWebElement), nil)
		mockQuoteTimestampTimeTagWebElement.On("GetAttribute", mock.Anything).Return("", errors.New("error while executing GetAttribute"))

		getQuoteTimestamp := tweets.MakeGetQuoteTimestamp()

		want := tweets.FailedToObtainQuotedTweetTimestamp
		_, got := getQuoteTimestamp(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuoteTimestampWebElement.AssertExpectations(t)
		mockQuoteTimestampTimeTagWebElement.AssertExpectations(t)
	}
}
