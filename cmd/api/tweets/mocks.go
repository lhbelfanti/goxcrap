package tweets

import "github.com/tebeka/selenium"

// MockRetrieveAll mocks RetrieveAll function
func MockRetrieveAll(tweets []Tweet, err error) RetrieveAll {
	return func() ([]Tweet, error) {
		return tweets, err
	}
}

// MockGetTweetHash mocks GetTweetHash function
func MockGetTweetHash(tweetHash TweetHash, err error) GetTweetHash {
	return func(tweetArticleElement selenium.WebElement) (TweetHash, error) {
		return tweetHash, err
	}
}

// MockGetTweetInformation mocks GetTweetInformation function
func MockGetTweetInformation(tweet Tweet, err error) GetTweetInformation {
	return func(tweetArticleElement selenium.WebElement, tweetID, tweetTimestamp string) (Tweet, error) {
		return tweet, err
	}
}

// MockGetAuthor mocks GetAuthor function
func MockGetAuthor(author string, err error) GetAuthor {
	return func(element selenium.WebElement) (string, error) {
		return author, err
	}
}

// MockGetTimestamp mocks GetTimestamp function
func MockGetTimestamp(timestamp string, err error) GetTimestamp {
	return func(element selenium.WebElement) (string, error) {
		return timestamp, err
	}
}

// MockIsAReply mocks IsAReply function
func MockIsAReply(isAReply bool) IsAReply {
	return func(element selenium.WebElement) bool {
		return isAReply
	}
}

// MockGetText mocks GetText function
func MockGetText(text string, err error) GetText {
	return func(element selenium.WebElement, isAReply bool) (string, error) {
		return text, err
	}
}

// MockGetImages mocks GetImages function
func MockGetImages(urls []string, err error) GetImages {
	return func(element selenium.WebElement, isAReply bool) ([]string, error) {
		return urls, err
	}
}

// MockHasQuote mocks HasQuote function
func MockHasQuote(hasQuote bool) HasQuote {
	return func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) bool {
		return hasQuote
	}
}

// MockIsQuoteAReply mocks IsQuoteAReply function
func MockIsQuoteAReply(isQuoteAReply bool) IsQuoteAReply {
	return func(tweetArticleElement selenium.WebElement, isAReply, hasTweetOnlyText bool) bool {
		return isQuoteAReply
	}
}

// MockGetQuoteText mocks GetQuoteText function
func MockGetQuoteText(text string, err error) GetQuoteText {
	return func(element selenium.WebElement, isAReply, hasTweetOnlyText, hasTweetOnlyImages, isQuoteAReply bool) (string, error) {
		return text, err
	}
}

// MockGetQuoteImages mocks GetQuoteImages function
func MockGetQuoteImages(urls []string, err error) GetQuoteImages {
	return func(element selenium.WebElement, isAReply, hasTweetOnlyText bool) ([]string, error) {
		return urls, err
	}
}

// MockTweet mocks a Tweet
func MockTweet() Tweet {
	return Tweet{
		ID:        "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b",
		Timestamp: "2024-02-26T18:31:49.000Z",
		IsAReply:  true,
		HasQuote:  true,
		Data: Data{
			HasText:   true,
			HasImages: true,
			Text:      "Tweet Text",
			Images:    []string{"https://url1.com", "https://url2.com"},
		},
	}
}

// MockQuote mocks a Quote
func MockQuote(IsAReply, hasText, hasImages bool, text string, images []string) Quote {
	return Quote{
		IsAReply: IsAReply,
		Data: Data{
			HasText:   hasText,
			HasImages: hasImages,
			Text:      text,
			Images:    images,
		},
	}
}

// MockTweetHash mocks a TweetHash
func MockTweetHash() TweetHash {
	return TweetHash{
		ID:        "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b",
		Timestamp: "2024-02-26T18:31:49.000Z",
	}
}
