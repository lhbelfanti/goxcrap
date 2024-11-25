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

func TestGetAvatar_success(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetAvatarWebElement := new(elements.MockWebElement)
	mockImg := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetAvatarWebElement), nil)
	mockTweetAvatarWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockImg), nil)
	mockImg.On("GetAttribute", "src").Return("test_url", nil)

	getAvatar := tweets.MakeGetAvatar()

	want := "test_url"
	got, err := getAvatar(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetAvatarWebElement.AssertExpectations(t)
}

func TestGetAvatar_failsWhenFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getAvatar := tweets.MakeGetAvatar()

	want := tweets.FailedToObtainTweetAvatarElement
	_, got := getAvatar(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
}

func TestGetAvatar_failsWhenAvatarImageThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetAvatarWebElement := new(elements.MockWebElement)
	mockImg := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetAvatarWebElement), nil)
	mockTweetAvatarWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockImg), errors.New("error while executing FindElement"))
	mockImg.On("GetAttribute", "src").Return("test_url", nil)

	getAvatar := tweets.MakeGetAvatar()

	want := tweets.FailedToObtainTweetAvatarImage
	_, got := getAvatar(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetAvatarWebElement.AssertExpectations(t)
}

func TestGetAvatar_failsWhenImgGetAttributeThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetAvatarWebElement := new(elements.MockWebElement)
	mockImg := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetAvatarWebElement), nil)
	mockTweetAvatarWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockImg), nil)
	mockImg.On("GetAttribute", "src").Return("", errors.New("error while executing GetAttribute"))

	getAvatar := tweets.MakeGetAvatar()

	want := tweets.FailedToObtainTweetAvatarSrcFromImage
	_, got := getAvatar(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetAvatarWebElement.AssertExpectations(t)
}

func TestGetQuoteAvatar_success(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuotedTweetAvatarWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuotedTweetAvatarWebElement), nil)
		mockQuotedTweetAvatarWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockImg), nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getQuoteAvatar := tweets.MakeGetQuoteAvatar()

		want := "test_url"
		got, err := getQuoteAvatar(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuotedTweetAvatarWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteAvatar_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

		getQuoteAvatar := tweets.MakeGetQuoteAvatar()

		want := tweets.FailedToObtainQuotedTweetAvatarElement
		_, got := getQuoteAvatar(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteAvatar_failsWhenAvatarImageThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuotedTweetAvatarWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuotedTweetAvatarWebElement), nil)
		mockQuotedTweetAvatarWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockImg), errors.New("error while executing FindElement"))
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getQuoteAvatar := tweets.MakeGetQuoteAvatar()

		want := tweets.FailedToObtainQuotedTweetAvatarImage
		_, got := getQuoteAvatar(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuotedTweetAvatarWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteAvatar_failsWhenImgGetAttributeThrowsError(t *testing.T) {
	for _, test := range []struct {
		hasTweetOnlyText bool
	}{
		{hasTweetOnlyText: false},
		{hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockQuotedTweetAvatarWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockQuotedTweetAvatarWebElement), nil)
		mockQuotedTweetAvatarWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockImg), nil)
		mockImg.On("GetAttribute", "src").Return("", errors.New("error while executing GetAttribute"))

		getQuoteAvatar := tweets.MakeGetQuoteAvatar()

		want := tweets.FailedToObtainQuotedTweetAvatarSrcFromImage
		_, got := getQuoteAvatar(context.Background(), mockTweetArticleWebElement, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockQuotedTweetAvatarWebElement.AssertExpectations(t)
	}
}
