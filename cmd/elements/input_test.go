package elements_test

import (
	"errors"
	"fmt"
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
	mockWaitAndRetrieve := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	got := retrieveAndClickButton(selenium.ByName, "name", "element", "input", 10*time.Minute, elements.NewElementError)

	assert.Nil(t, got)
}

func TestRetrieveAndFillInput_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWaitAndRetrieve := elements.MockMakeWaitAndRetrieve(nil, err)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.NewElementError(fmt.Sprintf(elements.FailedToRetrieveInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, elements.NewElementError)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputClickThrowsError(t *testing.T) {
	err := errors.New("error while executing input.Click")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("Click").Return(err)
	mockWaitAndRetrieve := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.NewElementError(fmt.Sprintf(elements.FailedToClickInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, elements.NewElementError)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputSendKeysThrowsError(t *testing.T) {
	err := errors.New("error while executing input.SendKeys")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWebElement.On("SendKeys", "input").Return(err)
	mockWaitAndRetrieve := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.NewElementError(fmt.Sprintf(elements.FailedToFillInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, elements.NewElementError)

	assert.Equal(t, want, got)
}
