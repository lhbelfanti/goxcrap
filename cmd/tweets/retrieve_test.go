package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
	"goxcrap/cmd/page"
	"goxcrap/cmd/tweets"
)

func TestRetrieveAll_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockTweet := tweets.MockTweet()
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(mockTweet, nil)
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation, mockScroll)

	want := []tweets.Tweet{mockTweet}
	got, err := retrieveAll()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_successEvenWhenGatherTweetInformationThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(tweets.MockTweet(), errors.New("error while executing GatherTweetInformation"))
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation, mockScroll)

	var want []tweets.Tweet
	got, err := retrieveAll()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_successEvenWhenScrollPageThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockTweet := tweets.MockTweet()
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(mockTweet, nil)
	mockScroll := page.MockScroll(errors.New("error while executing Scroll"))

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation, mockScroll)

	want := []tweets.Tweet{mockTweet}
	got, err := retrieveAll()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, errors.New("error while executing WaitAndRetrieveElement"))
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(tweets.MockTweet(), nil)
	mockScroll := page.MockScroll(nil)

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation, mockScroll)

	want := tweets.FailedToRetrieveArticles
	_, got := retrieveAll()

	assert.Equal(t, want, got)
}
