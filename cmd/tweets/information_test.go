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
		hasQuote       bool
		isQuoteAReply  bool
		getTextError   error
		getImagesError error
	}{
		{isAReply: false},
		{isAReply: false, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: true},
		{isAReply: true, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: false, hasQuote: true},
		{isAReply: false, hasQuote: true, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: true, hasQuote: true},
		{isAReply: true, hasQuote: true, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: false, hasQuote: true, isQuoteAReply: true},
		{isAReply: false, hasQuote: true, isQuoteAReply: true, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: true, hasQuote: true, isQuoteAReply: true},
		{isAReply: true, hasQuote: true, isQuoteAReply: true, getTextError: getTextErr, getImagesError: getImagesErr},
	} {
		mockGetAuthor := tweets.MockGetAuthor("author", nil)
		mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
		mockIsAReply := tweets.MockIsAReply(test.isAReply)
		mockGetText := tweets.MockGetText("Tweet Text", test.getTextError)
		mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, test.getImagesError)
		mockHasQuote := tweets.MockHasQuote(test.hasQuote)
		mockIsQuoteAReply := tweets.MockIsQuoteAReply(test.isQuoteAReply)
		mockTweetArticleWebElement := new(elements.MockWebElement)

		getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply)

		want := tweets.MockTweet()
		want.IsAReply = test.isAReply
		want.HasQuote = test.hasQuote
		want.Quote = tweets.Quote{IsAReply: test.isQuoteAReply}

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
	mockHasQuote := tweets.MockHasQuote(false)
	mockIsQuoteAReply := tweets.MockIsQuoteAReply(false)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply)

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
	mockHasQuote := tweets.MockHasQuote(false)
	mockIsQuoteAReply := tweets.MockIsQuoteAReply(false)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetInformation(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}
