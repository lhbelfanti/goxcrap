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

const (
	tweetIDElementFromHeaderXPath string = "div[position()=2]/div[position()=2]/div[position()=1]/div/div[position()=1]/div/div/div[position()=2]/div/div[position()=3]"
	tweetIDElementFromFooterXPath string = "div[position()=3]/div[position()=4]/div/div[position()=1]/div/div[position()=1]"
)

func TestGetID_success(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDWebElement := new(elements.MockWebElement)
	mockTweetIDATagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDWebElement), nil)
	mockTweetIDWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDATagWebElement), nil)
	mockTweetIDATagWebElement.On("GetAttribute", mock.Anything).Return("https://x.com/user/status/123456789", nil)

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

func TestGetIDFromTweetPage_success(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDFromHeaderWebElement := new(elements.MockWebElement)
	mockTweetIDATagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, tweetIDElementFromHeaderXPath).Return(selenium.WebElement(mockTweetIDFromHeaderWebElement), nil)
	mockTweetIDFromHeaderWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDATagWebElement), nil)
	mockTweetIDATagWebElement.On("GetAttribute", mock.Anything).Return("https://x.com/user/status/123456789", nil)

	getIDFromTweetPage := tweets.MakeGetIDFromTweetPage()

	want := "123456789"
	got, err := getIDFromTweetPage(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetIDFromHeaderWebElement.AssertExpectations(t)
	mockTweetIDATagWebElement.AssertExpectations(t)
}

func TestGetIDFromTweetPage_successEvenWhenFirstFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDFromFooterWebElement := new(elements.MockWebElement)
	mockTweetIDATagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, tweetIDElementFromHeaderXPath).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))
	mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, tweetIDElementFromFooterXPath).Return(selenium.WebElement(mockTweetIDFromFooterWebElement), nil)
	mockTweetIDFromFooterWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDATagWebElement), nil)
	mockTweetIDATagWebElement.On("GetAttribute", mock.Anything).Return("https://x.com/user/status/123456789", nil)

	getIDFromTweetPage := tweets.MakeGetIDFromTweetPage()

	want := "123456789"
	got, err := getIDFromTweetPage(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetIDFromFooterWebElement.AssertExpectations(t)
	mockTweetIDATagWebElement.AssertExpectations(t)
}

func TestGetIDFromTweetPage_failsWhenFirstAndSecondFindElementThrowError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, tweetIDElementFromHeaderXPath).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))
	mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, tweetIDElementFromFooterXPath).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getIDFromTweetPage := tweets.MakeGetIDFromTweetPage()

	want := tweets.FailedToObtainTweetIDElementFromTweetPage
	_, got := getIDFromTweetPage(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
}

func TestGetIDFromTweetPage_failsWhenThirdFindElementThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDWebElement), nil)
	mockTweetIDWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(new(elements.MockWebElement)), errors.New("error while executing FindElement"))

	getIDFromTweetPage := tweets.MakeGetIDFromTweetPage()

	want := tweets.FailedToObtainTweetIDATag
	_, got := getIDFromTweetPage(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetIDWebElement.AssertExpectations(t)
}

func TestGetIDFromTweetPage_failsWhenGetAttributeThrowsError(t *testing.T) {
	mockTweetArticleWebElement := new(elements.MockWebElement)
	mockTweetIDWebElement := new(elements.MockWebElement)
	mockTweetIDATagWebElement := new(elements.MockWebElement)
	mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDWebElement), nil)
	mockTweetIDWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetIDATagWebElement), nil)
	mockTweetIDATagWebElement.On("GetAttribute", mock.Anything).Return("", errors.New("error while executing GetAttribute"))

	getIDFromTweetPage := tweets.MakeGetIDFromTweetPage()

	want := tweets.FailedToObtainTweetIDATagHref
	_, got := getIDFromTweetPage(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
	mockTweetArticleWebElement.AssertExpectations(t)
	mockTweetIDWebElement.AssertExpectations(t)
	mockTweetIDATagWebElement.AssertExpectations(t)
}
