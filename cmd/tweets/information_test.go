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

func TestGetTweetInformation_success(t *testing.T) {
	findElementErr := errors.New("error while executing FindElement")
	getTextErr := errors.New("error while executing GetText")
	getImagesErr := errors.New("error while executing GetImages")

	for _, test := range []struct {
		isAReply         bool
		findElementError error
		getTextError     error
		getImagesError   error
	}{
		{isAReply: false, findElementError: findElementErr},
		{isAReply: false, findElementError: findElementErr, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: true, findElementError: nil},
		{isAReply: true, findElementError: nil, getTextError: getTextErr, getImagesError: getImagesErr},
	} {
		mockGetAuthor := tweets.MockGetAuthor("author", nil)
		mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
		mockGetText := tweets.MockGetText("Tweet Text", test.getTextError)
		mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, test.getImagesError)
		mockWebElement := new(elements.MockWebElement)
		mockWantedWebElement := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), test.findElementError)

		getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockGetText, mockGetImages)

		want := tweets.MockTweet()
		want.IsAReply = test.isAReply
		got, err := getTweetInformation(mockWebElement)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetTweetInformation_failsWhenGetAuthorThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("", errors.New("error while executing GetAuthor"))
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockGetText := tweets.MockGetText("Tweet Text", nil)
	mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockGetText, mockGetImages)

	want := tweets.FailedToObtainTweetAuthorInformation
	_, got := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
}

func TestGetTweetInformation_failsWhenGetTimestampThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("", errors.New("error while executing GetTimestamp"))
	mockGetText := tweets.MockGetText("Tweet Text", nil)
	mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockGetText, mockGetImages)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
}
