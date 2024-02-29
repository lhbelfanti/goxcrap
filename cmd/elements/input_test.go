package elements_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
)

func TestRetrieveAndFillInput_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWebElement.On("SendKeys", "input").Return(nil)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	got := retrieveAndClickButton(selenium.ByName, "name", "element", "input", 10*time.Minute)

	assert.Nil(t, got)
}

func TestRetrieveAndFillInput_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, err)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.FailedToRetrieveInput
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputClickThrowsError(t *testing.T) {
	err := errors.New("error while executing input.Click")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("Click").Return(err)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.FailedToClickInput
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputSendKeysThrowsError(t *testing.T) {
	err := errors.New("error while executing input.SendKeys")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWebElement.On("SendKeys", "input").Return(err)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.FailedToFillInput
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute)

	assert.Equal(t, want, got)
}
