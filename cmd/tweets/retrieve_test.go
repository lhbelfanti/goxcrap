package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
	"goxcrap/cmd/tweets"
)

func TestRetrieveAll_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockTweet := tweets.MockTweet()
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(mockTweet, nil)

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation)

	want := []tweets.Tweet{mockTweet, mockTweet}
	got, err := retrieveAll()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_successEvenWhenGatherTweetInformationThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, nil)
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(tweets.MockTweet(), errors.New("error"))

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation)

	var want []tweets.Tweet
	got, err := retrieveAll()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestRetrieveAll_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockRetrieveAll := elements.MockWaitAndRetrieveAll([]selenium.WebElement{mockWebElement, mockWebElement}, errors.New("error"))
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(tweets.MockTweet(), nil)

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation)

	want := tweets.FailedToRetrieveArticles
	_, got := retrieveAll()

	assert.Equal(t, want, got)
}
