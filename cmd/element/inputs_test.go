package element_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/element"
)

func TestRetrieveAndFillInput_success(t *testing.T) {
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWebElement.On("SendKeys", "input").Return(nil)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	got := retrieveAndClickButton(selenium.ByName, "name", "element", "input", 10*time.Minute, element.NewElementError)

	assert.Nil(t, got)
}

func TestRetrieveAndFillInput_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(nil, err)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToRetrieveInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, element.NewElementError)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputClickThrowsError(t *testing.T) {
	err := errors.New("error while executing input.Click")
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(err)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToClickInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, element.NewElementError)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputSendKeysThrowsError(t *testing.T) {
	err := errors.New("error while executing input.SendKeys")
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWebElement.On("SendKeys", "input").Return(err)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToFillInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, element.NewElementError)

	assert.Equal(t, want, got)
}
