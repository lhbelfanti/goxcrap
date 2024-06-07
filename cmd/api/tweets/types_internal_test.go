package tweets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTweet_ToString_success(t *testing.T) {
	want := `
------------------------
--- Tweet ---
 ID: 02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b 
 Timestamp: 2024-02-26T18:31:49.000Z 
 IsAReply: true 
 HasQuote: true 
   --- Data ---
   HasText: true 
   HasImages: true 
   Text: Tweet Text 
   Images: [https://url1.com https://url2.com] 
 --- Quote ---
 IsAReply: false 
   --- Data ---
   HasText: false 
   HasImages: false 
   Text:  
   Images: []
------------------------

`
	got := MockTweet().toString()

	assert.Equal(t, want, got)
}
