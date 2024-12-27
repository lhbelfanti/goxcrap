package tweets

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
)

// MockRetrieveAll mocks RetrieveAll function
func MockRetrieveAll(tweets []Tweet, err error) RetrieveAll {
	return func(ctx context.Context) ([]Tweet, error) {
		return tweets, err
	}
}

// MockGetTweetInformation mocks GetTweetInformation function
func MockGetTweetInformation(tweet Tweet, err error) GetTweetInformation {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement, tweetID string) (Tweet, error) {
		return tweet, err
	}
}

// MockGetID mocks GetID function
func MockGetID(id string, err error) GetID {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		return id, err
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

// MockGetAvatar mocks GetAvatar function
func MockGetAvatar(avatar string, err error) GetAvatar {
	return func(ctx context.Context, tweetArticleElement selenium.WebElement) (string, error) {
		return avatar, err
	}
}

// MockIsAReply mocks IsAReply function
func MockIsAReply(isAReply bool) IsAReply {
	return func(element selenium.WebElement) bool {
		return isAReply
	}
}

// MockGetText mocks GetText function
func MockGetText(text string, hasLongText bool, err error) GetText {
	return func(ctx context.Context, element selenium.WebElement, isAReply bool) (string, bool, error) {
		return text, hasLongText, err
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

// MockOpenAndRetrieveArticleByID mocks OpenAndRetrieveArticleByID function
func MockOpenAndRetrieveArticleByID(element selenium.WebElement, err error) OpenAndRetrieveArticleByID {
	return func(ctx context.Context, author, id string) (selenium.WebElement, error) {
		return element, err
	}
}

// MockGetIDFromTweetPage mocks GetIDFromTweetPage function
func MockGetIDFromTweetPage(id string, err error) GetIDFromTweetPage {
	return func(ctx context.Context, element selenium.WebElement) (string, error) {
		return id, err
	}
}

// MockGetLongText mocks GetLongText function
func MockGetLongText(text string, err error) GetLongText {
	return func(ctx context.Context, element selenium.WebElement, isAReply bool) (string, error) {
		return text, err
	}
}

// MockTweet mocks a Tweet
func MockTweet() Tweet {
	return Tweet{
		ID:       "123456789012345",
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

// MockLongTextElement mocks the tweet Long Text Elements, its dependencies and all the functions that will be executed on them
func MockLongTextElement() (*elements.MockWebElement, *elements.MockWebElement, *elements.MockWebElement) {
	mockTweetLongTextWebElement := new(elements.MockWebElement)
	mockTextPartSpanWebElement := new(elements.MockWebElement)
	mockTextPartImg := new(elements.MockWebElement)
	mockTweetLongTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement), selenium.WebElement(mockTextPartImg)}, nil)
	mockTextPartSpanWebElement.On("TagName").Return("span", nil)
	mockTextPartSpanWebElement.On("Text").Return("Tweet Text ", nil)
	mockTextPartImg.On("TagName").Return("img", nil)
	mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", nil)

	return mockTweetLongTextWebElement, mockTextPartSpanWebElement, mockTextPartImg
}
