package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
	"goxcrap/cmd/tweets"
)

func TestIsAReply_success(t *testing.T) {
	for _, test := range []struct {
		findElementError error
		text             string
		want             bool
	}{
		{findElementError: nil, text: "Replying to", want: true},
		{findElementError: errors.New("error while executing FindElement"), text: "", want: false},
		{findElementError: nil, text: "Text does not contain 'Replying To'", want: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetReplyTextWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetReplyTextWebElement), test.findElementError)
		mockTweetReplyTextWebElement.On("Text").Return(test.text, nil)

		isAReply := tweets.MakeIsAReply()

		got := isAReply(mockTweetArticleWebElement)

		assert.Equal(t, test.want, got)
	}
}

func TestHasQuote_success(t *testing.T) {
	for _, test := range []struct {
		isAReply         bool
		findElementError error
		want             bool
	}{
		{isAReply: false, findElementError: nil, want: true},
		{isAReply: true, findElementError: nil, want: true},
		{isAReply: false, findElementError: errors.New("error while executing FindElement"), want: false},
		{isAReply: true, findElementError: errors.New("error while executing FindElement"), want: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(new(elements.MockWebElement), test.findElementError)

		hasQuote := tweets.MakeHasQuote()

		got := hasQuote(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, test.want, got)
	}
}

func TestIsQuoteAReply_success(t *testing.T) {
	for _, test := range []struct {
		isAReply         bool
		hasTweetOnlyText bool
		findElementError error
		want             bool
	}{
		{isAReply: false, hasTweetOnlyText: false, findElementError: nil, want: true},
		{isAReply: true, hasTweetOnlyText: false, findElementError: nil, want: true},
		{isAReply: false, hasTweetOnlyText: false, findElementError: errors.New("error while executing FindElement"), want: false},
		{isAReply: true, hasTweetOnlyText: false, findElementError: errors.New("error while executing FindElement"), want: false},
		{isAReply: false, hasTweetOnlyText: true, findElementError: nil, want: true},
		{isAReply: true, hasTweetOnlyText: true, findElementError: nil, want: true},
		{isAReply: false, hasTweetOnlyText: true, findElementError: errors.New("error while executing FindElement"), want: false},
		{isAReply: true, hasTweetOnlyText: true, findElementError: errors.New("error while executing FindElement"), want: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(new(elements.MockWebElement), test.findElementError)

		isQuoteAReply := tweets.MakeIsQuoteAReply()

		got := isQuoteAReply(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, test.want, got)
	}
}
