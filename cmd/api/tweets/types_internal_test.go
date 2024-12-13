package tweets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTweet_String_success(t *testing.T) {
	want := `
------------------------
--- Tweet ---
ID: 123456789012345 
HasQuote: true
 --- Data ---
 Author: tweetauthor
 Avatar: https://tweet_avatar.com
 Timestamp: 2024-02-26T18:31:49.000Z
 IsAReply: true
 HasText: true
 HasImages: true
 Text: Tweet Text
 Images: [https://url1.com https://url2.com]
--- Quote ---
 --- Data ---
 Author: quoteauthor
 Avatar: https://quote_avatar.com
 Timestamp: 2023-02-26T18:31:49.000Z
 IsAReply: false
 HasText: true
 HasImages: true
 Text: test
 Images: [https://quote_url.com]
------------------------

`
	mockTweet := MockTweet()
	mockQuote := MockQuote(false, true, true, "test", []string{"https://quote_url.com"})
	mockQuote.Author = "quoteauthor"
	mockQuote.Avatar = "https://quote_avatar.com"
	mockQuote.Timestamp = "2023-02-26T18:31:49.000Z"
	mockTweet.Quote = mockQuote

	got := mockTweet.String()

	assert.Equal(t, want, got)
}
