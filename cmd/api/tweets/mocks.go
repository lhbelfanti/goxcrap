package tweets

import (
	"context"

	"github.com/tebeka/selenium"
)

// MockRetrieveAll mocks RetrieveAll function
func MockRetrieveAll(tweets []Tweet, err error) RetrieveAll {
	return func(ctx context.Context) ([]Tweet, error) {
		return tweets, err
	}
}

// MockGetTweetHash mocks GetTweetHash function
func MockGetTweetHash(tweetHash TweetHash, err error) GetTweetHash {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (TweetHash, error) {
		return tweetHash, err
	}
}

// MockGetTweetInformation mocks GetTweetInformation function
func MockGetTweetInformation(tweet Tweet, err error) GetTweetInformation {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, tweetHash TweetHash) (Tweet, error) {
		return tweet, err
	}
}

// MockGetAvatar mocks GetAvatar function
func MockGetAvatar(avatar string, err error) GetAvatar {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		return avatar, err
	}
}

// MockGetAuthor mocks GetAuthor function
func MockGetAuthor(author string, err error) GetAuthor {
	return func(ctx context.Context, element selenium.WebElement) (string, error) {
		return author, err
	}
}

// MockGetTimestamp mocks GetTimestamp function
func MockGetTimestamp(timestamp string, err error) GetTimestamp {
	return func(ctx context.Context, element selenium.WebElement) (string, error) {
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
	return func(ctx context.Context, element selenium.WebElement, isAReply bool) (string, error) {
		return text, err
	}
}

// MockGetImages mocks GetImages function
func MockGetImages(urls []string, err error) GetImages {
	return func(ctx context.Context, element selenium.WebElement, isAReply bool) ([]string, error) {
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

// MockGetQuoteAuthor mocks GetQuoteAuthor function
func MockGetQuoteAuthor(author string, err error) GetQuoteAuthor {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error) {
		return author, err
	}
}

// MockGetQuoteAvatar mocks GetQuoteAvatar function
func MockGetQuoteAvatar(avatar string, err error) GetQuoteAvatar {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error) {
		return avatar, err
	}
}

// MockGetQuoteTimestamp mocks GetQuoteTimestamp function
func MockGetQuoteTimestamp(timestamp string, err error) GetQuoteTimestamp {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, hasTweetOnlyText bool) (string, error) {
		return timestamp, err
	}
}

// MockGetQuoteText mocks GetQuoteText function
func MockGetQuoteText(text string, err error) GetQuoteText {
	return func(ctx context.Context, element selenium.WebElement, isAReply, hasTweetOnlyText, hasTweetOnlyImages, isQuoteAReply bool) (string, error) {
		return text, err
	}
}

// MockGetQuoteImages mocks GetQuoteImages function
func MockGetQuoteImages(urls []string, err error) GetQuoteImages {
	return func(ctx context.Context, element selenium.WebElement, isAReply, hasTweetOnlyText bool) ([]string, error) {
		return urls, err
	}
}

// MockTweet mocks a Tweet
func MockTweet() Tweet {
	return Tweet{
		ID:       "6b19232cdaa5ab34588aa59614fb2e868d6ad3a9f75f3ac4166fef23da9f209b",
		HasQuote: true,
		Data: Data{
			Author:    "tweetauthor",
			Avatar:    "https://tweet_avatar.com",
			Timestamp: "2024-02-26T18:31:49.000Z",
			IsAReply:  true,
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
		Data: Data{
			IsAReply:  IsAReply,
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
		ID:        "6b19232cdaa5ab34588aa59614fb2e868d6ad3a9f75f3ac4166fef23da9f209b",
		Author:    "tweetauthor",
		Timestamp: "2024-02-26T18:31:49.000Z",
	}
}
