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
	tweetOnlyTextXPath      string = "div[position()=2]/div[position()=3]/div[position()=1]/div[position()=1]/span"
	replyTweetOnlyTextXPath string = "div[position()=2]/div[position()=4]/div[position()=1]/div[position()=1]/span"
	tweetImagesXPath        string = "div[position()=2]/div[position()=3]/div[position()=1]/div/div/div/div"
	replyTweetImagesXPath   string = "div[position()=2]/div[position()=4]/div[position()=1]/div/div/div/div"
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
		got, err := getImages(context.Background(), mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
		mockImg.AssertExpectations(t)
	}
}

func TestGetImages_failsWhenSpanElementIsFound(t *testing.T) {
	for _, test := range []struct {
		isAReply  bool
		spanXPath string
	}{
		{isAReply: false, spanXPath: tweetOnlyTextXPath},
		{isAReply: true, spanXPath: replyTweetOnlyTextXPath},
	} {
		mockTweetArticleWebElement := new(elements.MockWebElement)
		mockTweetImagesWebElement := new(elements.MockWebElement)
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.spanXPath).Return(selenium.WebElement(mockTweetImagesWebElement), nil)

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetImagesElement
		_, got := getImages(context.Background(), mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
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
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.spanXPath).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))
		mockTweetArticleWebElement.On("FindElement", selenium.ByXPATH, test.imagesXPath).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))

		getImages := tweets.MakeGetImages()

		want := tweets.FailedToObtainTweetImagesElement
		_, got := getImages(context.Background(), mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
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
		_, got := getImages(context.Background(), mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
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
		_, got := getImages(context.Background(), mockTweetArticleWebElement, test.isAReply)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
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
		got, err := getQuoteImages(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
		mockImg.AssertExpectations(t)
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
		mockTweetArticleWebElement.On("FindElement", mock.Anything, mock.Anything).Return(selenium.WebElement(mockTweetImagesWebElement), errors.New("error while executing FindElement"))

		getQuoteImages := tweets.MakeGetQuoteImages()

		want := tweets.FailedToObtainQuotedTweetImagesElement
		_, got := getQuoteImages(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
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
		_, got := getQuoteImages(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
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
		_, got := getQuoteImages(context.Background(), mockTweetArticleWebElement, test.isAReply, test.hasTweetOnlyText)

		assert.Equal(t, want, got)
		mockTweetArticleWebElement.AssertExpectations(t)
		mockTweetImagesWebElement.AssertExpectations(t)
	}
}
