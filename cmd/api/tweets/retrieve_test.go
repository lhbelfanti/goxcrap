package tweets_test

import (
	"context"
	"errors"
	"goxcrap/internal/log"
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
	ctx := context.Background()
	ctx = log.With(ctx, log.Param("page_url", "test"))

	want := []tweets.Tweet{mockTweet}
	got, err := retrieveAll(ctx)

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
	ctx := context.Background()
	ctx = log.With(ctx, log.Param("page_url", "test"))

	var want []tweets.Tweet
	got, err := retrieveAll(ctx)

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
	ctx := context.Background()
	ctx = log.With(ctx, log.Param("page_url", "test"))

	var want []tweets.Tweet
	got, err := retrieveAll(ctx)

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
	ctx := context.Background()
	ctx = log.With(ctx, log.Param("page_url", "test"))

	want := []tweets.Tweet{mockTweet}
	got, err := retrieveAll(ctx)

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

func TestOpenAndRetrieveArticleByID_success(t *testing.T) {
	mockOpenNewTab := page.MockOpenNewTab(nil)
	mockArticleWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieveElements := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockArticleWebElement}, nil)
	mockTweetID := "12345"
	mockGetIDFromTweetPage := tweets.MockGetIDFromTweetPage(mockTweetID, nil)

	openAndRetrieveArticleByID := tweets.MakeOpenAndRetrieveArticleByID(mockOpenNewTab, mockWaitAndRetrieveElements, mockGetIDFromTweetPage)

	want := mockArticleWebElement
	got, err := openAndRetrieveArticleByID(context.Background(), "author", mockTweetID)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestOpenAndRetrieveArticleByID_failsWhenOpenNewTabThrowsError(t *testing.T) {
	mockOpenNewTab := page.MockOpenNewTab(errors.New("error while executing OpenAndRetrieveArticleByID"))
	mockArticleWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieveElements := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockArticleWebElement}, nil)
	mockTweetID := "12345"
	mockGetIDFromTweetPage := tweets.MockGetIDFromTweetPage(mockTweetID, nil)

	openAndRetrieveArticleByID := tweets.MakeOpenAndRetrieveArticleByID(mockOpenNewTab, mockWaitAndRetrieveElements, mockGetIDFromTweetPage)

	want := tweets.FailedToLoadTweetPage
	_, got := openAndRetrieveArticleByID(context.Background(), "author", mockTweetID)

	assert.Equal(t, want, got)
}

func TestOpenAndRetrieveArticleByID_failsWhenWaitAndRetrieveElementsThrowsError(t *testing.T) {
	mockOpenNewTab := page.MockOpenNewTab(nil)
	mockWaitAndRetrieveElements := elements.MockWaitAndRetrieveAll([]selenium.WebElement{}, errors.New("error while executing MockWaitAndRetrieveAll"))
	mockTweetID := "12345"
	mockGetIDFromTweetPage := tweets.MockGetIDFromTweetPage(mockTweetID, nil)

	openAndRetrieveArticleByID := tweets.MakeOpenAndRetrieveArticleByID(mockOpenNewTab, mockWaitAndRetrieveElements, mockGetIDFromTweetPage)

	want := tweets.FailedToRetrieveArticles
	_, got := openAndRetrieveArticleByID(context.Background(), "author", mockTweetID)

	assert.Equal(t, want, got)
}

func TestOpenAndRetrieveArticleByID_failsWhenGetTweetIDFromTweetPageThrowsError(t *testing.T) {
	mockOpenNewTab := page.MockOpenNewTab(nil)
	mockArticleWebElement := new(elements.MockWebElement)
	mockWaitAndRetrieveElements := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockArticleWebElement}, nil)
	mockTweetID := "12345"
	mockGetIDFromTweetPage := tweets.MockGetIDFromTweetPage(mockTweetID, errors.New("error while executing GetIDFromTweetPage"))

	openAndRetrieveArticleByID := tweets.MakeOpenAndRetrieveArticleByID(mockOpenNewTab, mockWaitAndRetrieveElements, mockGetIDFromTweetPage)

	want := tweets.FailedToRetrieveArticle
	_, got := openAndRetrieveArticleByID(context.Background(), "author", mockTweetID)

	assert.Equal(t, want, got)
}
