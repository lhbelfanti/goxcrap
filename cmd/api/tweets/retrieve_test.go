package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/tweets"
)

func TestRetrieveAll_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, errors.New("can't find empty state"))
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockGetTweetID := tweets.MockGetID("123456789012345", nil)
	mockTweet := tweets.MockTweet()
	mockGetTweetInformation := tweets.MockGetTweetInformation(mockTweet, nil)
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockWaitAndRetrieve, mockRetrieveAll, mockGetTweetID, mockGetTweetInformation, mockScroll)

	want := []tweets.Tweet{mockTweet}
	got, err := retrieveAll(context.Background())

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_successEvenWhenGetTweetIDThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, errors.New("can't find empty state"))
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockGetTweetID := tweets.MockGetID("123456789012345", errors.New("error while executing GetID"))
	mockGetTweetInformation := tweets.MockGetTweetInformation(tweets.MockTweet(), nil)
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockWaitAndRetrieve, mockRetrieveAll, mockGetTweetID, mockGetTweetInformation, mockScroll)

	var want []tweets.Tweet
	got, err := retrieveAll(context.Background())

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_successEvenWhenGetTweetInformationThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, errors.New("can't find empty state"))
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockGetTweetID := tweets.MockGetID("123456789012345", nil)
	mockGetTweetInformation := tweets.MockGetTweetInformation(tweets.MockTweet(), errors.New("error while executing GetTweetInformation"))
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockWaitAndRetrieve, mockRetrieveAll, mockGetTweetID, mockGetTweetInformation, mockScroll)

	var want []tweets.Tweet
	got, err := retrieveAll(context.Background())

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_successEvenWhenScrollPageThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, errors.New("can't find empty state"))
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockGetTweetID := tweets.MockGetID("123456789012345", nil)
	mockTweet := tweets.MockTweet()
	mockGetTweetInformation := tweets.MockGetTweetInformation(mockTweet, nil)
	mockScroll := page.MockScroll(errors.New("error while executing Scroll"))

	retrieveAll := tweets.MakeRetrieveAll(mockWaitAndRetrieve, mockRetrieveAll, mockGetTweetID, mockGetTweetInformation, mockScroll)

	want := []tweets.Tweet{mockTweet}
	got, err := retrieveAll(context.Background())

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_failsWhenWaitAndRetrieveElementDoesntThrowError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll(nil, nil)
	mockGetTweetID := tweets.MockGetID("123456789012345", nil)
	mockGetTweetInformation := tweets.MockGetTweetInformation(tweets.MockTweet(), nil)
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockWaitAndRetrieve, mockRetrieveAll, mockGetTweetID, mockGetTweetInformation, mockScroll)

	want := tweets.EmptyStateNoArticlesToRetrieve
	_, got := retrieveAll(context.Background())

	assert.Equal(t, want, got)
}

func TestRetrieveAll_failsWhenWaitAndRetrieveAllThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, errors.New("can't find empty state"))
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, errors.New("error while executing WaitAndRetrieveElement"))
	mockGetTweetID := tweets.MockGetID("123456789012345", nil)
	mockGetTweetInformation := tweets.MockGetTweetInformation(tweets.MockTweet(), nil)
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockWaitAndRetrieve, mockRetrieveAll, mockGetTweetID, mockGetTweetInformation, mockScroll)

	want := tweets.FailedToRetrieveArticles
	_, got := retrieveAll(context.Background())

	assert.Equal(t, want, got)
}
