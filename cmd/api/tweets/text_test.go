package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/tweets"
)

func TestGetText_success(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTextPartImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement), selenium.WebElement(mockTextPartImg)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", nil)
		mockTextPartImg.On("TagName").Return("img", nil)
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", nil)

		getText := tweets.MakeGetText()

		want := "text🙂"
		got, err := getText(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
		mockTextPartImg.AssertExpectations(t)
	}
}

func TestGetText_successEvenIfEmojisCantBeObtained(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTextPartImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement), selenium.WebElement(mockTextPartImg)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", nil)
		mockTextPartImg.On("TagName").Return("img", nil)
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", errors.New("error while executing GetAttribute"))

		getText := tweets.MakeGetText()

		want := "text"
		got, err := getText(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
		mockTextPartImg.AssertExpectations(t)
	}
}

func TestGetText_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), errors.New("error while executing FindElement"))

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextElement
		_, got := getText(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
	}
}

func TestGetText_failsWhenFindElementsThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, errors.New("error while executing FindElements"))

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextParts
		_, got := getText(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetText_failsWhenTagNameThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", errors.New("error while executing TagName"))

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextPartTagName
		_, got := getText(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetText_failsWhenTextThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", errors.New("error while executing Text"))

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextFromSpan
		_, got := getText(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteText_success(t *testing.T) {
	for _, test := range []struct {
		isAReply           bool
		hasTweetOnlyText   bool
		hasTweetOnlyImages bool
		isQuoteAReply      bool
	}{
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTextPartImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement), selenium.WebElement(mockTextPartImg)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", nil)
		mockTextPartImg.On("TagName").Return("img", nil)
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", nil)

		getQuoteText := tweets.MakeGetQuoteText()

		want := "text🙂"
		got, err := getQuoteText(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
		mockTextPartImg.AssertExpectations(t)
	}
}

func TestGetQuoteText_successEvenIfEmojisCantBeObtained(t *testing.T) {
	for _, test := range []struct {
		isAReply           bool
		hasTweetOnlyText   bool
		hasTweetOnlyImages bool
		isQuoteAReply      bool
	}{
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTextPartImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement), selenium.WebElement(mockTextPartImg)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", nil)
		mockTextPartImg.On("TagName").Return("img", nil)
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", errors.New("error while executing GetAttribute"))

		getQuoteText := tweets.MakeGetQuoteText()

		want := "text"
		got, err := getQuoteText(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
		mockTextPartImg.AssertExpectations(t)
	}
}

func TestGetQuoteText_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply           bool
		hasTweetOnlyText   bool
		hasTweetOnlyImages bool
		isQuoteAReply      bool
	}{
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), errors.New("error while executing FindElement"))

		getQuoteText := tweets.MakeGetQuoteText()

		want := tweets.FailedToObtainQuotedTweetTextElement
		_, got := getQuoteText(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteText_failsWhenFindElementsThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply           bool
		hasTweetOnlyText   bool
		hasTweetOnlyImages bool
		isQuoteAReply      bool
	}{
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, errors.New("error while executing FindElements"))

		getQuoteText := tweets.MakeGetQuoteText()

		want := tweets.FailedToObtainQuotedTweetTextParts
		_, got := getQuoteText(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteText_failsWhenTagNameThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply           bool
		hasTweetOnlyText   bool
		hasTweetOnlyImages bool
		isQuoteAReply      bool
	}{
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", errors.New("error while executing TagName"))

		getQuoteText := tweets.MakeGetQuoteText()

		want := tweets.FailedToObtainQuotedTweetTextPartTagName
		_, got := getQuoteText(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetQuoteText_failsWhenTextThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply           bool
		hasTweetOnlyText   bool
		hasTweetOnlyImages bool
		isQuoteAReply      bool
	}{
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: true, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: true, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: true},
		{isAReply: false, hasTweetOnlyText: true, hasTweetOnlyImages: false, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: true, isQuoteAReply: false},
		{isAReply: false, hasTweetOnlyText: false, hasTweetOnlyImages: false, isQuoteAReply: false},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", errors.New("error while executing Text"))

		getQuoteText := tweets.MakeGetQuoteText()

		want := tweets.FailedToObtainQuotedTweetTextFromSpan
		_, got := getQuoteText(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}
