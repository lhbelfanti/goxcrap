package tweets_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
	"goxcrap/cmd/tweets"
)

func TestRetrieveAll_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebElements := []selenium.WebElement{mockWebElement, mockWebElement}
	mockRetrieveAll := elements.MockWaitAndRetrieveAll(mockWebElements, nil)
	mockTweet := tweets.MockTweet()
	mockGatherTweetInformation := tweets.MockGatherTweetInformation(mockTweet, nil)

	retrieveAll := tweets.MakeRetrieveAll(mockRetrieveAll, mockGatherTweetInformation)

	want := []tweets.Tweet{mockTweet, mockTweet}
	got, err := retrieveAll()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}
