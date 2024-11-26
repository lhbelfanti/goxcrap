package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/tweets"
)

func TestGetTweetBodyInformation_success(t *testing.T) {
	getAvatarErr := errors.New("error while executing GetAvatar")
	getTextErr := errors.New("error while executing GetText")
	getImagesErr := errors.New("error while executing GetImages")
	getQuoteAuthorErr := errors.New("error while executing GetQuoteAuthor")
	getQuoteAvatarErr := errors.New("error while executing GetQuoteAvatar")
	getQuoteTimestampErr := errors.New("error while executing GetQuoteTimestamp")
	getQuoteTextErr := errors.New("error while executing GetQuoteText")
	getQuoteImagesErr := errors.New("error while executing GetQuoteImages")

	for _, test := range []struct {
		isAReply               bool
		hasQuote               bool
		isQuoteAReply          bool
		getAvatarError         error
		getTextError           error
		getImagesError         error
		getQuoteAuthorError    error
		getQuoteAvatarError    error
		getQuoteTimestampError error
		getQuoteTextError      error
		getQuoteImagesError    error
	}{
		// Basic cases
		{isAReply: false, hasQuote: false, isQuoteAReply: false},
		{isAReply: true, hasQuote: false, isQuoteAReply: false},
		{isAReply: false, hasQuote: true, isQuoteAReply: false},
		{isAReply: true, hasQuote: true, isQuoteAReply: false},
		{isAReply: false, hasQuote: true, isQuoteAReply: true},
		{isAReply: true, hasQuote: true, isQuoteAReply: true},
		// With errors
		{isAReply: false, hasQuote: false, isQuoteAReply: false, getAvatarError: getAvatarErr, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: false, hasQuote: true, isQuoteAReply: false, getQuoteAuthorError: getQuoteAuthorErr, getQuoteAvatarError: getQuoteAvatarErr, getQuoteTimestampError: getQuoteTimestampErr},
		// Combination cases
		{isAReply: true, hasQuote: true, isQuoteAReply: true, getTextError: getTextErr, getImagesError: getImagesErr, getQuoteTextError: getQuoteTextErr, getQuoteImagesError: getQuoteImagesErr},
		{isAReply: false, hasQuote: true, isQuoteAReply: true, getAvatarError: getAvatarErr, getQuoteAuthorError: getQuoteAuthorErr, getQuoteAvatarError: getQuoteAvatarErr},
	} {
		mockTweetHash := tweets.MockTweetHash()
		mockIsAReply := tweets.MockIsAReply(test.isAReply)
		mockGetAvatar := tweets.MockGetAvatar("https://tweet_avatar.com", test.getAvatarError)
		mockGetText := tweets.MockGetText("Tweet Text", test.getTextError)
		mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, test.getImagesError)
		mockHasQuote := tweets.MockHasQuote(test.hasQuote)
		mockIsQuoteAReply := tweets.MockIsQuoteAReply(test.isQuoteAReply)
		mockGetQuoteAuthor := tweets.MockGetQuoteAuthor("quoteauthor", test.getQuoteAuthorError)
		mockGetQuoteAvatar := tweets.MockGetQuoteAvatar("https://quote_avatar.com", test.getQuoteAvatarError)
		mockGetQuoteTimestamp := tweets.MockGetQuoteTimestamp("2023-02-26T18:31:49.000Z", test.getQuoteTimestampError)
		mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", test.getQuoteTextError)
		mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, test.getQuoteImagesError)
		mockTweetArticleWebElement := new(elements.MockWebElement)

		getTweetInformation := tweets.MakeGetTweetInformation(mockIsAReply, mockGetAvatar, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteAuthor, mockGetQuoteAvatar, mockGetQuoteTimestamp, mockGetQuoteText, mockGetQuoteImages)

		mockTweet := tweets.MockTweet()
		mockTweet.IsAReply = test.isAReply
		mockTweet.HasQuote = test.hasQuote
		mockTweet.Quote = tweets.MockQuote(test.isQuoteAReply, test.hasQuote, test.hasQuote, "", nil)

		if test.hasQuote {
			mockTweet.Quote.Text = "Quote Text"
			mockTweet.Quote.Images = []string{"https://url1.com", "https://url2.com"}
			mockTweet.Quote.Author = "quoteauthor"
			mockTweet.Quote.Avatar = "https://quote_avatar.com"
			mockTweet.Quote.Timestamp = "2023-02-26T18:31:49.000Z"
		}

		want := mockTweet
		got, err := getTweetInformation(context.Background(), mockTweetArticleWebElement, mockTweetHash)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestMakeGetTweetHash_success(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("tweetauthor", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetHash := tweets.MakeGetTweetHash(mockGetAuthor, mockGetTimestamp)

	want := tweets.MockTweetHash()
	got, _ := getTweetHash(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}

func TestMakeGetTweetHash_failsWhenGetAuthorThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("", errors.New("error while executing GetAuthor"))
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetHash := tweets.MakeGetTweetHash(mockGetAuthor, mockGetTimestamp)

	want := tweets.FailedToObtainTweetAuthorInformation
	_, got := getTweetHash(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}

func TestMakeGetTweetHash_failsWhenGetTimestampThrowsError(t *testing.T) {
	mockGetAuthor := tweets.MockGetAuthor("author", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("", errors.New("error while executing GetTimestamp"))
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetHash := tweets.MakeGetTweetHash(mockGetAuthor, mockGetTimestamp)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetHash(context.Background(), mockTweetArticleWebElement)

	assert.Equal(t, want, got)
}
