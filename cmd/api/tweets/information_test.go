package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/tweets"
)

func TestGetTweetBodyInformation_success(t *testing.T) {
	getTextErr := errors.New("error while executing GetText")
	getImagesErr := errors.New("error while executing GetImages")
	getQuoteTextErr := errors.New("error while executing GetQuoteText")
	getQuoteImagesErr := errors.New("error while executing GetQuoteImages")

	for _, test := range []struct {
		isAReply            bool
		hasQuote            bool
		isQuoteAReply       bool
		getTextError        error
		getImagesError      error
		getQuoteTextError   error
		getQuoteImagesError error
	}{
		{isAReply: false, hasQuote: false, isQuoteAReply: false},
		{isAReply: false, hasQuote: false, isQuoteAReply: false, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: true, hasQuote: false, isQuoteAReply: false},
		{isAReply: true, hasQuote: false, isQuoteAReply: false, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: false, hasQuote: true, isQuoteAReply: false},
		{isAReply: false, hasQuote: true, isQuoteAReply: false, getTextError: getTextErr, getImagesError: getImagesErr, getQuoteTextError: getQuoteTextErr, getQuoteImagesError: getQuoteImagesErr},
		{isAReply: true, hasQuote: true, isQuoteAReply: false},
		{isAReply: true, hasQuote: true, isQuoteAReply: false, getTextError: getTextErr, getImagesError: getImagesErr, getQuoteTextError: getQuoteTextErr, getQuoteImagesError: getQuoteImagesErr},
		{isAReply: false, hasQuote: true, isQuoteAReply: true},
		{isAReply: false, hasQuote: true, isQuoteAReply: true, getTextError: getTextErr, getImagesError: getImagesErr, getQuoteTextError: getQuoteTextErr, getQuoteImagesError: getQuoteImagesErr},
		{isAReply: true, hasQuote: true, isQuoteAReply: true},
		{isAReply: true, hasQuote: true, isQuoteAReply: true, getTextError: getTextErr, getImagesError: getImagesErr, getQuoteTextError: getQuoteTextErr, getQuoteImagesError: getQuoteImagesErr},
	} {
		mockTweetHash := tweets.MockTweetHash()
		mockIsAReply := tweets.MockIsAReply(test.isAReply)
		mockGetText := tweets.MockGetText("Tweet Text", test.getTextError)
		mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, test.getImagesError)
		mockHasQuote := tweets.MockHasQuote(test.hasQuote)
		mockIsQuoteAReply := tweets.MockIsQuoteAReply(test.isQuoteAReply)
		mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", test.getQuoteTextError)
		mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, test.getQuoteImagesError)
		mockTweetArticleWebElement := new(elements.MockWebElement)

		getTweetInformation := tweets.MakeGetTweetInformation(mockIsAReply, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteText, mockGetQuoteImages)

		mockTweet := tweets.MockTweet()
		mockTweet.IsAReply = test.isAReply
		mockTweet.HasQuote = test.hasQuote
		mockTweet.Quote = tweets.MockQuote(test.isQuoteAReply, test.hasQuote, test.hasQuote, "", nil)
		if test.hasQuote {
			mockTweet.Quote.Text = "Quote Text"
			mockTweet.Quote.Images = []string{"https://url1.com", "https://url2.com"}
		}

		want := mockTweet
		got, err := getTweetInformation(mockTweetArticleWebElement, mockTweetHash.ID, mockTweetHash.Timestamp)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestMakeGetTweetHash_success(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetHash := tweets.MakeGetTweetHash(mockGetAuthor, mockGetTimestamp)

	want := tweets.MockTweetHash()
	got, _ := getTweetHash(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}

func TestMakeGetTweetHash_failsWhenGetAuthorThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("", errors.New("error while executing GetAuthor"))
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetHash := tweets.MakeGetTweetHash(mockGetAuthor, mockGetTimestamp)

	want := tweets.FailedToObtainTweetAuthorInformation
	_, got := getTweetHash(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}

func TestMakeGetTweetHash_failsWhenGetTimestampThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("", errors.New("error while executing GetTimestamp"))
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetHash := tweets.MakeGetTweetHash(mockGetAuthor, mockGetTimestamp)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetHash(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}
