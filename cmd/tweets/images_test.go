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

const (
	tweetOnlyTextXPath      string = "div[position()=3]/div[position()=1]/div[position()=1]/span"
	replyTweetOnlyTextXPath string = "div[position()=4]/div[position()=1]/div[position()=1]/span"
	tweetImagesXPath        string = "div[position()=3]/div[position()=1]/div/div/div/div"
	replyTweetImagesXPath   string = "div[position()=4]/div[position()=1]/div/div/div/div"
)

func TestGetImages_success(t *testing.T) {
	for _, test := range []struct {
		isAReply    bool
		spanXPath   string
		imagesXPath string
	}{
		{isAReply: false, spanXPath: tweetOnlyTextXPath, imagesXPath: tweetImagesXPath},
		{isAReply: true, spanXPath: replyTweetOnlyTextXPath, imagesXPath: replyTweetImagesXPath},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.spanXPath).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.imagesXPath).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getImages := tweets.MakeGetImages()

		want := []string{"test_url", "test_url"}
		got, err := getImages(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetImages_failsWhenSpanElementIsFound(t *testing.T) {
	for _, test := range []struct {
		isAReply             bool
		spanXPath            string
		imagesXPath          string
		findSpanElementError error
	}{
		{isAReply: false, spanXPath: tweetOnlyTextXPath, imagesXPath: tweetImagesXPath, findSpanElementError: nil},
		{isAReply: true, spanXPath: replyTweetOnlyTextXPath, imagesXPath: replyTweetImagesXPath, findSpanElementError: nil},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.spanXPath).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.imagesXPath).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetImagesElement
		_, got := getImages(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetImages_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply    bool
		spanXPath   string
		imagesXPath string
	}{
		{isAReply: false, spanXPath: tweetOnlyTextXPath, imagesXPath: tweetImagesXPath},
		{isAReply: true, spanXPath: replyTweetOnlyTextXPath, imagesXPath: replyTweetImagesXPath},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.spanXPath).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.imagesXPath).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetImagesElement
		_, got := getImages(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetImages_failsWhenFindElementsThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply    bool
		spanXPath   string
		imagesXPath string
	}{
		{isAReply: false, spanXPath: tweetOnlyTextXPath, imagesXPath: tweetImagesXPath},
		{isAReply: true, spanXPath: replyTweetOnlyTextXPath, imagesXPath: replyTweetImagesXPath},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockTweetImageWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.spanXPath).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.imagesXPath).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTweetImageWebElement), selenium.WebElement(mockTweetImageWebElement)}, errors.New("error while executing FindElements"))
		mockTweetImageWebElement.On("GetAttribute", "src").Return("test_url", nil)

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetImages
		_, got := getImages(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetImages_failsWhenAllImgGetAttributeThrowError(t *testing.T) {
	for _, test := range []struct {
		isAReply    bool
		spanXPath   string
		imagesXPath string
	}{
		{isAReply: false, spanXPath: tweetOnlyTextXPath, imagesXPath: tweetImagesXPath},
		{isAReply: true, spanXPath: replyTweetOnlyTextXPath, imagesXPath: replyTweetImagesXPath},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockTweetImageWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.spanXPath).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.imagesXPath).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTweetImageWebElement), selenium.WebElement(mockTweetImageWebElement)}, nil)
		mockTweetImageWebElement.On("GetAttribute", "src").Return("test_url", errors.New("error while executing GetAttribute"))

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetSrcFromImage
		_, got := getImages(mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
	}
}

func TestGetQuoteImages_success(t *testing.T) {
	for _, test := range []struct {
		isAReply         bool
		hasTweetOnlyText bool
	}{
		{isAReply: false, hasTweetOnlyText: false},
		{isAReply: false, hasTweetOnlyText: true},
		{isAReply: true, hasTweetOnlyText: false},
		{isAReply: true, hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getQuoteImages := tweets.MakeGetQuoteImages()

		want := []string{"test_url", "test_url"}
		got, err := getQuoteImages(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestGetQuoteImages_failsWhenFindElementThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply         bool
		hasTweetOnlyText bool
	}{
		{isAReply: false, hasTweetOnlyText: false},
		{isAReply: false, hasTweetOnlyText: true},
		{isAReply: true, hasTweetOnlyText: false},
		{isAReply: true, hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockImg := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockImg), selenium.WebElement(mockImg)}, nil)
		mockImg.On("GetAttribute", "src").Return("test_url", nil)

		getQuoteImages := tweets.MakeGetQuoteImages()

		want := tweets.FailedToObtainQuotedTweetImagesElement
		_, got := getQuoteImages(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
	}
}

func TestGetQuoteImages_failsWhenFindElementsThrowsError(t *testing.T) {
	for _, test := range []struct {
		isAReply         bool
		hasTweetOnlyText bool
	}{
		{isAReply: false, hasTweetOnlyText: false},
		{isAReply: false, hasTweetOnlyText: true},
		{isAReply: true, hasTweetOnlyText: false},
		{isAReply: true, hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockTweetImageWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTweetImageWebElement), selenium.WebElement(mockTweetImageWebElement)}, errors.New("error while executing FindElements"))
		mockTweetImageWebElement.On("GetAttribute", "src").Return("test_url", nil)

		getQuoteImages := tweets.MakeGetQuoteImages()

		want := tweets.FailedToObtainQuotedTweetImages
		_, got := getQuoteImages(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
	}
}

func TestGetQuoteImages_failsWhenAllImgGetAttributeThrowError(t *testing.T) {
	for _, test := range []struct {
		isAReply         bool
		hasTweetOnlyText bool
	}{
		{isAReply: false, hasTweetOnlyText: false},
		{isAReply: false, hasTweetOnlyText: true},
		{isAReply: true, hasTweetOnlyText: false},
		{isAReply: true, hasTweetOnlyText: true},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockTweetImageWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesWebElement), nil)
		mockTweetImagesWebElement.On("FindElements", mock.Anything, mock.Anything).Return([]selenium.WebElement{selenium.WebElement(mockTweetImageWebElement), selenium.WebElement(mockTweetImageWebElement)}, nil)
		mockTweetImageWebElement.On("GetAttribute", "src").Return("test_url", errors.New("error while executing GetAttribute"))

		getQuoteImages := tweets.MakeGetQuoteImages()

		want := tweets.FailedToObtainQuotedTweetSrcFromImage
		_, got := getQuoteImages(mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
	}
}
