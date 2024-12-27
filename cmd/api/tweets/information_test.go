package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/tweets"
)

func TestGetTweetInformation_success(t *testing.T) {
	getAvatarErr := errors.New("error while executing GetAvatar")
	getTextErr := errors.New("error while executing GetText")
	getImagesErr := errors.New("error while executing GetImages")
	getQuoteAuthorErr := errors.New("error while executing GetQuoteAuthor")
	getQuoteAvatarErr := errors.New("error while executing GetQuoteAvatar")
	getQuoteTimestampErr := errors.New("error while executing GetQuoteTimestamp")
	getQuoteTextErr := errors.New("error while executing GetQuoteText")
	getQuoteImagesErr := errors.New("error while executing GetQuoteImages")
	getLongTextErr := errors.New("error while executing GetLongText")

	for _, test := range []struct {
		isAReply               bool
		hasQuote               bool
		isQuoteAReply          bool
		hasLongText            bool
		getAvatarError         error
		getTextError           error
		getImagesError         error
		getQuoteAuthorError    error
		getQuoteAvatarError    error
		getQuoteTimestampError error
		getQuoteTextError      error
		getQuoteImagesError    error
		getLongTextError       error
	}{
		// Basic cases
		{isAReply: false, hasQuote: false, isQuoteAReply: false, hasLongText: true},
		{isAReply: true, hasQuote: false, isQuoteAReply: false, hasLongText: true},
		{isAReply: false, hasQuote: true, isQuoteAReply: false, hasLongText: true},
		{isAReply: true, hasQuote: true, isQuoteAReply: false, hasLongText: true},
		{isAReply: false, hasQuote: true, isQuoteAReply: true, hasLongText: true},
		{isAReply: true, hasQuote: true, isQuoteAReply: true, hasLongText: true},
		{isAReply: false, hasQuote: false, isQuoteAReply: false, hasLongText: false},
		{isAReply: true, hasQuote: false, isQuoteAReply: false, hasLongText: false},
		{isAReply: false, hasQuote: true, isQuoteAReply: false, hasLongText: false},
		{isAReply: true, hasQuote: true, isQuoteAReply: false, hasLongText: false},
		{isAReply: false, hasQuote: true, isQuoteAReply: true, hasLongText: false},
		{isAReply: true, hasQuote: true, isQuoteAReply: true, hasLongText: false},
		// With errors
		{isAReply: false, hasQuote: false, isQuoteAReply: false, hasLongText: true, getAvatarError: getAvatarErr, getTextError: getTextErr, getImagesError: getImagesErr},
		{isAReply: false, hasQuote: true, isQuoteAReply: false, hasLongText: true, getQuoteAuthorError: getQuoteAuthorErr, getQuoteAvatarError: getQuoteAvatarErr, getQuoteTimestampError: getQuoteTimestampErr},
		// Combination cases
		{isAReply: true, hasQuote: true, isQuoteAReply: true, hasLongText: true, getTextError: getTextErr, getLongTextError: getLongTextErr, getImagesError: getImagesErr, getQuoteTextError: getQuoteTextErr, getQuoteImagesError: getQuoteImagesErr},
		{isAReply: false, hasQuote: true, isQuoteAReply: true, hasLongText: true, getAvatarError: getAvatarErr, getQuoteAuthorError: getQuoteAuthorErr, getQuoteAvatarError: getQuoteAvatarErr},
	} {
		mockIsAReply := tweets.MockIsAReply(test.isAReply)
		mockGetAuthor := tweets.MockGetAuthor("tweetauthor", nil)
		mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
		mockGetAvatar := tweets.MockGetAvatar("https://tweet_avatar.com", test.getAvatarError)
		mockGetText := tweets.MockGetText("Tweet Text", test.hasLongText, test.getTextError)
		mockHasQuote := tweets.MockHasQuote(test.hasQuote)
		mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, test.getImagesError)
		mockIsQuoteAReply := tweets.MockIsQuoteAReply(test.isQuoteAReply)
		mockGetQuoteAuthor := tweets.MockGetQuoteAuthor("quoteauthor", test.getQuoteAuthorError)
		mockGetQuoteAvatar := tweets.MockGetQuoteAvatar("https://quote_avatar.com", test.getQuoteAvatarError)
		mockGetQuoteTimestamp := tweets.MockGetQuoteTimestamp("2023-02-26T18:31:49.000Z", test.getQuoteTimestampError)
		mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", test.getQuoteTextError)
		mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, test.getQuoteImagesError)
		mockTweetWebElement := new(elements.MockWebElement)
		mockOpenAndRetrieveArticleByID := tweets.MockOpenAndRetrieveArticleByID(mockTweetWebElement, nil)
		mockGetLongText := tweets.MockGetLongText("Long Tweet Text 🙂", test.getLongTextError)
		mockCloseOpenedTabs := page.MockCloseOpenedTabs(nil)
		mockTweetArticleWebElement := new(elements.MockWebElement)

		getTweetInformation := tweets.MakeGetTweetInformation(mockIsAReply, mockGetAuthor, mockGetTimestamp, mockGetAvatar, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteAuthor, mockGetQuoteAvatar, mockGetQuoteTimestamp, mockGetQuoteText, mockGetQuoteImages, mockOpenAndRetrieveArticleByID, mockGetLongText, mockCloseOpenedTabs)

		mockTweet := tweets.MockTweet()
		mockTweet.IsAReply = test.isAReply
		mockTweet.HasQuote = test.hasQuote
		mockTweet.Quote = tweets.MockQuote(test.isQuoteAReply, test.hasQuote, test.hasQuote, "", nil)

		if test.hasLongText && test.getLongTextError == nil {
			mockTweet.Text = "Long Tweet Text 🙂"
		}

		if test.hasQuote {
			mockTweet.Quote.Text = "Quote Text"
			mockTweet.Quote.Images = []string{"https://url1.com", "https://url2.com"}
			mockTweet.Quote.Author = "quoteauthor"
			mockTweet.Quote.Avatar = "https://quote_avatar.com"
			mockTweet.Quote.Timestamp = "2023-02-26T18:31:49.000Z"
		}

		want := mockTweet
		got, err := getTweetInformation(context.Background(), mockTweetArticleWebElement, "123456789012345")

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetTweetInformation_failsWhenGetAuthorThrowsError(t *testing.T) {
	mockIsAReply := tweets.MockIsAReply(false)
	mockGetAuthor := tweets.MockGetAuthor("tweetauthor", errors.New("error while executing GetAuthor"))
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockIsAReply, mockGetAuthor, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	want := tweets.FailedToObtainTweetAuthorInformation
	_, got := getTweetInformation(context.Background(), mockTweetArticleWebElement, "123456789012345")

	assert.Equal(t, want, got)
}

func TestGetTweetInformation_failsWhenGetTimestampThrowsError(t *testing.T) {
	mockIsAReply := tweets.MockIsAReply(false)
	mockGetAuthor := tweets.MockGetAuthor("tweetauthor", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", errors.New("error while executing GetTimestamp"))
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockIsAReply, mockGetAuthor, mockGetTimestamp, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	want := tweets.FailedToObtainTweetTimestampInformation
	_, got := getTweetInformation(context.Background(), mockTweetArticleWebElement, "123456789012345")

	assert.Equal(t, want, got)
}

func TestGetTweetInformation_failsWhenOpenAndRetrieveTweetArticleByIDThrowsError(t *testing.T) {
	mockIsAReply := tweets.MockIsAReply(false)
	mockGetAuthor := tweets.MockGetAuthor("tweetauthor", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockGetAvatar := tweets.MockGetAvatar("https://tweet_avatar.com", nil)
	mockGetText := tweets.MockGetText("Tweet Text", true, nil)
	mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockHasQuote := tweets.MockHasQuote(false)
	mockIsQuoteAReply := tweets.MockIsQuoteAReply(false)
	mockGetQuoteAuthor := tweets.MockGetQuoteAuthor("quoteauthor", nil)
	mockGetQuoteAvatar := tweets.MockGetQuoteAvatar("https://quote_avatar.com", nil)
	mockGetQuoteTimestamp := tweets.MockGetQuoteTimestamp("2023-02-26T18:31:49.000Z", nil)
	mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", nil)
	mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockOpenAndRetrieveArticleByID := tweets.MockOpenAndRetrieveArticleByID(nil, tweets.FailedToLoadTweetPage)
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockIsAReply, mockGetAuthor, mockGetTimestamp, mockGetAvatar, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteAuthor, mockGetQuoteAvatar, mockGetQuoteTimestamp, mockGetQuoteText, mockGetQuoteImages, mockOpenAndRetrieveArticleByID, nil, nil)

	want := tweets.FailedToLoadTweetPage
	_, got := getTweetInformation(context.Background(), mockTweetArticleWebElement, "123456789012345")

	assert.Equal(t, want, got)
}

func TestGetTweetInformation_failsWhenCloseOpenedTabsThrowsError(t *testing.T) {
	mockIsAReply := tweets.MockIsAReply(false)
	mockGetAuthor := tweets.MockGetAuthor("tweetauthor", nil)
	mockGetTimestamp := tweets.MockGetTimestamp("2024-02-26T18:31:49.000Z", nil)
	mockGetAvatar := tweets.MockGetAvatar("https://tweet_avatar.com", nil)
	mockGetText := tweets.MockGetText("Tweet Text", true, nil)
	mockGetImages := tweets.MockGetImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockHasQuote := tweets.MockHasQuote(false)
	mockIsQuoteAReply := tweets.MockIsQuoteAReply(false)
	mockGetQuoteAuthor := tweets.MockGetQuoteAuthor("quoteauthor", nil)
	mockGetQuoteAvatar := tweets.MockGetQuoteAvatar("https://quote_avatar.com", nil)
	mockGetQuoteTimestamp := tweets.MockGetQuoteTimestamp("2023-02-26T18:31:49.000Z", nil)
	mockGetQuoteText := tweets.MockGetQuoteText("Quote Text", nil)
	mockGetQuoteImages := tweets.MockGetQuoteImages([]string{"https://url1.com", "https://url2.com"}, nil)
	mockTweetWebElement := new(elements.MockWebElement)
	mockOpenAndRetrieveArticleByID := tweets.MockOpenAndRetrieveArticleByID(mockTweetWebElement, nil)
	mockGetLongText := tweets.MockGetLongText("Long Tweet Text 🙂", nil)
	mockCloseOpenedTabs := page.MockCloseOpenedTabs(errors.New("error while executing page.CloseOpenedTabs"))
	mockTweetArticleWebElement := new(elements.MockWebElement)

	getTweetInformation := tweets.MakeGetTweetInformation(mockIsAReply, mockGetAuthor, mockGetTimestamp, mockGetAvatar, mockGetText, mockGetImages, mockHasQuote, mockIsQuoteAReply, mockGetQuoteAuthor, mockGetQuoteAvatar, mockGetQuoteTimestamp, mockGetQuoteText, mockGetQuoteImages, mockOpenAndRetrieveArticleByID, mockGetLongText, mockCloseOpenedTabs)

	want := tweets.FailedToCloseOpenedTabs
	_, got := getTweetInformation(context.Background(), mockTweetArticleWebElement, "123456789012345")

	assert.Equal(t, want, got)
}
