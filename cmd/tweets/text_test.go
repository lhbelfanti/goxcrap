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

func TestGetText_success(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockWebElement := new(elements.MockWebElement)
		mockTweetTextElement := new(elements.MockWebElement)
		mockTextPartSpan := new(elements.MockWebElement)
		mockTextPartImg := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextElement), nil)
		mockTweetTextElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpan), selenium.WebElement(mockTextPartImg)}, nil)
		mockTextPartSpan.On("TagName").Return("span", nil)
		mockTextPartSpan.On("Text").Return("text", nil)
		mockTextPartImg.On("TagName").Return("img", nil)
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", nil)

		getText := tweets.MakeGetText()

		want := "text🙂"
		got, err := getText(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetText_successEvenIfEmojisCantBeObtained(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		imgErr := errors.New("error")
		mockWebElement := new(elements.MockWebElement)
		mockTweetTextElement := new(elements.MockWebElement)
		mockTextPartSpan := new(elements.MockWebElement)
		mockTextPartImg := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextElement), nil)
		mockTweetTextElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpan), selenium.WebElement(mockTextPartImg)}, nil)
		mockTextPartSpan.On("TagName").Return("span", nil)
		mockTextPartSpan.On("Text").Return("text", nil)
		mockTextPartImg.On("TagName").Return("img", nil)
		mockTextPartImg.On("GetAttribute", mock.Anything).Return("🙂", imgErr)

		getText := tweets.MakeGetText()

		want := "text"
		got, err := getText(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetText_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		err := errors.New("error")
		mockWebElement := new(elements.MockWebElement)
		mockTweetTextElement := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextElement), err)

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextElement
		_, got := getText(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetText_failsWhenFindElementsThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		err := errors.New("error")
		mockWebElement := new(elements.MockWebElement)
		mockTweetTextElement := new(elements.MockWebElement)
		mockTextPartSpan := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextElement), nil)
		mockTweetTextElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpan)}, err)

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextParts
		_, got := getText(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetText_failsWhenTagNameThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		err := errors.New("error")
		mockWebElement := new(elements.MockWebElement)
		mockTweetTextElement := new(elements.MockWebElement)
		mockTextPartSpan := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextElement), nil)
		mockTweetTextElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpan)}, nil)
		mockTextPartSpan.On("TagName").Return("span", err)

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextPartTagName
		_, got := getText(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetText_failsWhenTextThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		err := errors.New("error")
		mockWebElement := new(elements.MockWebElement)
		mockTweetTextElement := new(elements.MockWebElement)
		mockTextPartSpan := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetTextElement), nil)
		mockTweetTextElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTextPartSpan)}, nil)
		mockTextPartSpan.On("TagName").Return("span", nil)
		mockTextPartSpan.On("Text").Return("text", err)

		getText := tweets.MakeGetText()

		want := tweets.FailedToObtainTweetTextFromSpan
		_, got := getText(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}