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

func TestGetTweetInformation_successWhenTweetIsAReply(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), nil)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp)

	want := tweets.MockTweet()
	got, err := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestGetTweetInformation_successWhenTweetIsNotAReply(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockWebElement := new(elements.MockWebElement)
	mockWantedWebElement := new(elements.MockWebElement)
	mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockWantedWebElement), errors.New("error"))

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp)

	want := tweets.MockTweet()
	want.IsAReply = false
	got, err := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestGetTweetInformation_failsWhenGetAuthorThrowsError(t *testing.T) {
	want := errors.New("error")
	mockGetAuthor := tweets.MockGetAuthor("", want)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp)

	_, got := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
}

func TestGetTweetInformation_failsWhenGetTimestampThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	want := errors.New("error")
	mockGetTimestamp := tweets.MockGetTimestamp("", want)
	mockWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp)

	_, got := getTweetInformation(mockWebElement)

	assert.Equal(t, want, got)
}
