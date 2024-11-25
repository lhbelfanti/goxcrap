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

func TestGetAuthor_success(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetAuthorWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetAuthorWebElement), nil)
	mockTweetAuthorWebElement.On("Text", mock.Anything).Return("test", nil)

	getAuthor := tweets.MakeGetAuthor()

	want := "test"
	got, err := getAuthor(context.Background(), mockTweetArticleWebElement)

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
	_, got := getAuthor(context.Background(), mockTweetArticleWebElement)

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
	_, got := getAuthor(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetAuthorWebElement.AssertExpectations(t)
}

func TestGetQuoteAuthor_success(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuoteAuthorWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuoteAuthorWebElement), nil)
		mockQuoteAuthorWebElement.On("Text", mock.Anything).Return("test", nil)

		getQuoteAuthor := tweets.MakeGetQuoteAuthor()

		want := "test"
		got, err := getQuoteAuthor(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuoteAuthorWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteAuthor_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

		getQuoteAuthor := tweets.MakeGetQuoteAuthor()

		want := tweets.FailedToObtainQuotedTweetAuthorElement
		_, got := getQuoteAuthor(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteAuthor_failsWhenTextThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuoteAuthorWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuoteAuthorWebElement), nil)
		mockQuoteAuthorWebElement.On("Text", mock.Anything).Return("", errors.New("error while executing Text"))

		getQuoteAuthor := tweets.MakeGetQuoteAuthor()

		want := tweets.FailedToObtainQuotedTweetAuthor
		_, got := getQuoteAuthor(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuoteAuthorWebElement.AssertExpectations(t)
	}
}
