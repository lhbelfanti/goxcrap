package elements_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
)

func TestRetrieveAndFillInput_success(t *testing.T) {
	mockInputWebElement := new(elements.MockWebElement)
	mockInputWebElement.On("Click").Return(nil)
	mockInputWebElement.On("SendKeys", "input").Return(nil)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockInputWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	got := retrieveAndClickButton(context.Background(), selenium.ByName, "name", "element", "input", 10*time.Minute)

	assert.Nil(t, got)
}

func TestRetrieveAndFillInput_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, err)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.FailedToRetrieveInput
	got := retrieveAndClickButton(context.Background(), selenium.ByName, "name", "test", "input", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputClickThrowsError(t *testing.T) {
	err := errors.New("error while executing input.Click")
	mockInputWebElement := new(elements.MockWebElement)
	mockInputWebElement.On("Click").Return(err)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockInputWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.FailedToClickInput
	got := retrieveAndClickButton(context.Background(), selenium.ByName, "name", "test", "input", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestRetrieveAndFillInput_failsWhenInputSendKeysThrowsError(t *testing.T) {
	err := errors.New("error while executing input.SendKeys")
	mockInputWebElement := new(elements.MockWebElement)
	mockInputWebElement.On("Click").Return(nil)
	mockInputWebElement.On("SendKeys", "input").Return(err)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockInputWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := elements.FailedToFillInput
	got := retrieveAndClickButton(context.Background(), selenium.ByName, "name", "test", "input", 10*time.Minute)

	assert.Equal(t, want, got)
}
