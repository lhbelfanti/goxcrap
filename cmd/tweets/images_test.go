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

func TestGetImages_success(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockWebElement := new(elements.MockWebElement)
		mockTweetImagesElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesElement), nil)
		mockTweetImagesElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getImages := tweets.MakeGetImages()

		want := []string{"test_url", "test_url"}
		got, err := getImages(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetImages_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockWebElement := new(elements.MockWebElement)
		mockTweetImagesElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesElement), errors.New("error while executing FindElement"))
		mockTweetImagesElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetImagesElement
		_, got := getImages(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetImages_failsWhenFindElementsThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockWebElement := new(elements.MockWebElement)
		mockTweetImagesElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesElement), nil)
		mockTweetImagesElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, errors.New("error while executing FindElements"))
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetImages
		_, got := getImages(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetImages_failsWhenAllImgGetAttributeThrowError(t *testing.T) {
	for _, test := range []struct {
		isAReply bool
	}{
		{isAReply: false},
		{isAReply: true},
	} {
		mockWebElement := new(elements.MockWebElement)
		mockTweetImagesElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesElement), nil)
		mockTweetImagesElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", errors.New("error while executing GetAttribute"))

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetSrcFromImage
		_, got := getImages(mockWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}
