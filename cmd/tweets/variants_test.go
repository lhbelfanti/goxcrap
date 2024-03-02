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
