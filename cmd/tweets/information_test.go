package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"goxcrap/cmd/elements"
	"goxcrap/cmd/tweets"
)

func TestGetTweetInformation_success(t *testing.T) {
	getTextErr := errors.New("error while executing GetText")
	getImagesErr := errors.New("error while executing GetImages")

	for _, test := range []struct {
		isAReply       bool
		getTextError   error
		getImagesError error
	}{
		{isAReply: false},
		{isAReply: false, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: true},
		{isAReply: true, getTextError: getTextErr, getImagesError: getImagesErr},
	} {
		mockGetAuthor := tweets.MockGetAuthor("author", nil)
		mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
		mockIsAReply := tweets.MockIsAReply(test.isAReply)
		mockGetText := tweets.MockGetText("Tweet Text", test.getTextError)
		mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, test.getImagesError)
		mockTweetArticleWebElement := new(elements.MockWebElement)

		getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages)

		want := tweets.MockTweet()
		want.IsAReply = test.isAReply
		got, err := getTweetInformation(mockTweetArticleWebElement)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetTweetInformation_failsWhenGetAuthorThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("", errors.New("error while executing GetAuthor"))
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockIsAReply := tweets.MockIsAReply(false)
	mockGetText := tweets.MockGetText("Tweet Text", nil)
	mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages)

	want := tweets.FailedToObtainTweetAuthorInformation
	_, got := getTweetInformation(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}

func TestGetTweetInformation_failsWhenGetTimestampThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("", errors.New("error while executing GetTimestamp"))
	mockIsAReply := tweets.MockIsAReply(false)
	mockGetText := tweets.MockGetText("Tweet Text", nil)
	mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetInformation(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}
