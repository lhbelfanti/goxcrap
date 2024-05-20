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
		mockGetAuthor := tweets.MockGetAuthor("author", nil)
		mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
		mockIsAReply := tweets.MockIsAReply(test.isAReply)
		mockGetText := tweets.MockGetText("Tweet Text", test.getTextError)
		mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, test.getImagesError)
		mockHasQuote := tweets.MockHasQuote(test.hasQuote)
		mockIsQuoteAReply := tweets.MockIsQuoteAReply(test.isQuoteAReply)
		mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", test.getQuoteTextError)
		mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, test.getQuoteImagesError)
		mockTweetArticleWebElement := new(elements.MockWebElement)

		getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteText, mockGetQuoteImages)

		mockTweet := tweets.MockTweet()
		mockTweet.IsAReply = test.isAReply
		mockTweet.HasQuote = test.hasQuote
		mockTweet.Quote = tweets.MockQuote(test.isQuoteAReply, test.hasQuote, test.hasQuote, "", nil)
		if test.hasQuote {
			mockTweet.Quote.Text = "Quote Text"
			mockTweet.Quote.Images = []string{"https://url1.com", "https://url2.com"}
		}

		want := mockTweet
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
	mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", nil)
	mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteText, mockGetQuoteImages)

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
	mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", nil)
	mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockGetAuthor, mockGetTimestamp, mockIsAReply, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteText, mockGetQuoteImages)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetInformation(mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}
