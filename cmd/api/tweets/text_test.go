package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/tweets"
)

const (
	tweetTextXPath              string = "div[position()=2]/div[position()=2]/div[position()=1]"
	replyTweetTextXPath         string = "div[position()=2]/div[position()=3]/div[position()=1]"
	tweetShowMoreTextXpath      string = "div[position()=2]/div[position()=2]/div[position()=2]"
	replyTweetShowMoreTextXPath string = "div[position()=2]/div[position()=3]/div[position()=1]"
)

func TestGetText_success(t *testing.T) {
	for _, test := range []struct {
		isAReply      bool
		textXPath     string
		longTextXPath string
	}{
		{isAReply: false, textXPath: tweetTextXPath, longTextXPath: tweetShowMoreTextXpath},
		{isAReply: true, textXPath: replyTweetTextXPath, longTextXPath: replyTweetShowMoreTextXPath},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetTextWebElement := new(elements.MockWebElement)
		mockTweetLongTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTextPartImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.textXPath).Return(selenium.WebElement(mockTweetTextWebElement), nil)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.longTextXPath).Return(selenium.WebElement(mockTweetLongTextWebElement), nil)
		mockTweetTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement), selenium.WebElement(mockTextPartImg)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", nil)
		mockTextPartImg.On("TagName").Return("img", nil)
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", nil)

		getText := tweets.MakeGetText()

		want := "text🙂"
		got, hasLongText, err := getText(context.Background(), mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.True(t, hasLongText)
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
		got, _, err := getText(context.Background(), mockTweetArticleWebElement, test.isAReply)

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
		_, _, got := getText(context.Background(), mockTweetArticleWebElement, test.isAReply)

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
		_, _, got := getText(context.Background(), mockTweetArticleWebElement, test.isAReply)

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
		_, _, got := getText(context.Background(), mockTweetArticleWebElement, test.isAReply)

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
		_, _, got := getText(context.Background(), mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetLongText_success(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetLongTextWebElement, mockTextPartSpanWebElement, mockTextPartImg := tweets.MockLongTextElement()
		mockTweetWebElement := new(elements.MockWebElement)
		mockTweetWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetLongTextWebElement), nil)

		getLongText := tweets.MakeGetLongText()

		want := "Tweet Text 🙂"
		got, err := getLongText(context.Background(), mockTweetWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetLongTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
		mockTextPartImg.AssertExpectations(t)
	}
}

func TestGetLongText_successEvenIfEmojisCantBeObtained(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetLongTextWebElement, mockTextPartSpanWebElement, mockTextPartImg := tweets.MockLongTextElement()
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", errors.New("error while executing GetAttribute"))
		mockTweetWebElement := new(elements.MockWebElement)
		mockTweetWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetLongTextWebElement), nil)

		getLongText := tweets.MakeGetLongText()

		want := "Tweet Text 🙂"
		got, err := getLongText(context.Background(), mockTweetWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetLongTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
		mockTextPartImg.AssertExpectations(t)
	}
}

func TestGetLongText_failsWhenFindElementsThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetLongTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetLongTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, errors.New("error while executing FindElements"))
		mockTweetWebElement := new(elements.MockWebElement)
		mockTweetWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetLongTextWebElement), nil)

		getLongText := tweets.MakeGetLongText()

		want := tweets.FailedToObtainTweetLongTextParts
		_, got := getLongText(context.Background(), mockTweetWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetLongTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetLongText_failsWhenTagNameThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetLongTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetLongTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", errors.New("error while executing TagName"))
		mockTweetWebElement := new(elements.MockWebElement)
		mockTweetWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetLongTextWebElement), nil)

		getLongText := tweets.MakeGetLongText()

		want := tweets.FailedToObtainTweetLongTextPartTagName
		_, got := getLongText(context.Background(), mockTweetWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetLongTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}

func TestGetLongText_failsWhenTextThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockTweetLongTextWebElement := new(elements.MockWebElement)
		mockTextPartSpanWebElement := new(elements.MockWebElement)
		mockTweetLongTextWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpanWebElement)}, nil)
		mockTextPartSpanWebElement.On("TagName").Return("span", nil)
		mockTextPartSpanWebElement.On("Text").Return("text", errors.New("error while executing Text"))
		mockTweetWebElement := new(elements.MockWebElement)
		mockTweetWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetLongTextWebElement), nil)

		getLongText := tweets.MakeGetLongText()

		want := tweets.FailedToObtainTweetLongTextFromSpan
		_, got := getLongText(context.Background(), mockTweetWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetLongTextWebElement.AssertExpectations(t)
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
		got, err := getQuoteText(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

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
		got, err := getQuoteText(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

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
		_, got := getQuoteText(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

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
		_, got := getQuoteText(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

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
		_, got := getQuoteText(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

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
		_, got := getQuoteText(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText, test.hasTweetOnlyImages, test.isQuoteAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetTextWebElement.AssertExpectations(t)
		mockTextPartSpanWebElement.AssertExpectations(t)
	}
}
