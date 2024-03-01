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
	for _, test := range []struct {
		isAReply         bool
		findElementError error
	}{
		{isAReply: false, findElementError: errors.New("error")},
		{isAReply: true, findElementError: nil},
	} {
		mockGetAuthor := tweets.MockGetAuthor("author", nil)
		mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
		mockGetText := tweets.MockGetText("Tweet Text", nil)
		mockWebElement := new(elements.MockWebElement)
		mockWantedWebElement := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), test.findElementError)

		getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockGetText)

		want := tweets.MockTweet()
		want.IsAReply = test.isAReply
		got, err := getTweetInformation(mockWebElement)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetTweetInformation_failsWhenGetAuthorThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("", errors.New("error"))
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockGetText := tweets.MockGetText("Tweet Text", nil)
	mockWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockGetText)

	want := tweets.FailedToObtainTweetAuthorInformation
	_, got := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
}

func TestGetTweetInformation_failsWhenGetTimestampThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("", errors.New("error"))
	mockGetText := tweets.MockGetText("Tweet Text", nil)
	mockWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockGetText)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
}
